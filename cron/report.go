package cron

import (
	"fmt"
	"github.com/WalletService/service"
	. "github.com/robfig/cron"
	"github.com/robfig/cron/v3"
	"os"
	"sync"
	"log"
	"encoding/csv"
)

type IReportCron interface {
	StartReportCron() *Cron
}

type reportcron struct {}

var(
	c 		*reportcron
	once 	sync.Once
	transactionService service.ITransactionService
)

// Making reportcron instance as singleton
func NewReportCron(service service.ITransactionService) IReportCron {
	if c==nil {
		once.Do(func() {
			transactionService = service
			c = &reportcron{}
		})
	}
	return c
}

func (c *reportcron) StartReportCron() *Cron {
	mCron := cron.New()
	//_, err := mCron.AddFunc("0 9 * * *", fetchReport)
	_, err := mCron.AddFunc("@every 30s", fetchReport)
	if err != nil {
		log.Println("Something went wrong with setting up the cron, error : ", err)
	}
	mCron.Start()
	return mCron
}

func fetchReport() {
	transactions, err := transactionService.GetActiveTransactionsService()
	if err != nil {
		log.Println("Error occured while generating report : ", err)
	}
	records := [][]string{}
	// if no transactions are active, simply return
	if len(*transactions) == 0 {
		return
	}
	records = append(records, []string{"userid", "username", "txnid", "txntype", "amount"})
	for _, t := range *transactions {
		records = append(records, []string{fmt.Sprint(t.Wallet.UserID), t.Wallet.User.Name,
			fmt.Sprint(t.ID), t.TxnType, fmt.Sprint(t.Amount)})
	}
	log.Println("Records : ", records)

	f, err := os.Create("report.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}

	err = transactionService.UpdateActiveTransactionsService()
	if err != nil {
		log.Fatal("Error occurred while marking transactions inactive : ", err)
	}
}
