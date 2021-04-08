package scheduler

import (
	"fmt"
	"github.com/WalletService/cache"
	"github.com/WalletService/service"
	"github.com/robfig/cron"
	"os"
	"strconv"
	"sync"
	"log"
	"encoding/csv"
)

type IReportCron interface {
	StartReportCron() *cron.Cron
}

type reportcron struct {}

var(
	c 		*reportcron
	once 	sync.Once
	transactionService service.ITransactionService
	cronCache cache.ICronCache
)

// Making reportcron instance as singleton
func NewReportCron(service service.ITransactionService, cache cache.ICronCache) IReportCron {
	if c==nil {
		once.Do(func() {
			transactionService = service
			cronCache = cache
			c = &reportcron{}
		})
	}
	return c
}

func (c *reportcron) StartReportCron() *cron.Cron {
	mCron := cron.New()
	//_, err := mCron.AddFunc("0 9 * * *", fetchReport)
	err := mCron.AddFunc("@every 30s", lockAndFetchReport)
	if err != nil {
		log.Println("Something went wrong with setting up the scheduler, error : ", err)
	}
	mCron.Start()
	return mCron
}

func lockAndFetchReport() {
	key := "wallet-db-mysql" // shared resource
	value := "wallet-service-instance_process-" + strconv.Itoa(os.Getpid())   // client name
	// acquire lock to execute cron
	if cronCache.AcquireLock(key, value) {
		fetchReport()
		cronCache.ReleaseLock(key, value)
	}
}

func fetchReport() {
	transactions, err := transactionService.GetActiveTransactionsService()
	if err != nil {
		log.Println("Error occurred while generating report : ", err)
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
