package service

import (
	"errors"
	. "github.com/WalletService/model"
	"github.com/WalletService/repository"
	"strings"
)

type ITransactionService interface{
	GetTransactionService(id int) (*Transaction, error)
	GetTransactionsByWalletIdService(id int) (*[]Transaction, error)
	GetTransactionsService() (*[]Transaction, error)
	GetActiveTransactionsService() (*[]Transaction, error)
	PostTransactionService(transaction *Transaction, walletID int) (*Transaction, error)
	UpdateTransactionService(id int, transaction *Transaction) (*Transaction, error)
	UpdateActiveTransactionsService() error
	//DeleteTransactionService(id int) error
}

type transactionService struct {}

var (
	transactionRepository repository.ITransactionRepository
	iWalletService IWalletService
)

func NewTransactionService(repository repository.ITransactionRepository, iService IWalletService) ITransactionService {
	transactionRepository = repository
	iWalletService = iService
	return &transactionService{}
}

func (transactionService *transactionService) GetTransactionService(id int) (*Transaction, error) {
	return transactionRepository.GetTransactionById(id)
}

func (transactionService *transactionService) GetTransactionsByWalletIdService(id int) (*[]Transaction, error) {
	return transactionRepository.GetTransactionsByWalletId(id)
}

func (transactionService *transactionService) GetTransactionsService() (*[]Transaction, error) {
	return transactionRepository.GetAllTransactions()
}

func (transactionService *transactionService) GetActiveTransactionsService() (*[]Transaction, error) {
	return transactionRepository.GetAllActiveTransactions()
}

func (transactionService *transactionService) PostTransactionService(transaction *Transaction, walletID int) (*Transaction, error) {
	txnType := strings.ToLower(transaction.TxnType)
	if txnType != "credit" && txnType != "debit" {
		return nil, errors.New("Txn Type can only be credit or debit!")
	}
	wallet, err := iWalletService.GetWalletService(walletID, false)
	if err != nil {
		return nil, err
	}
	if wallet.IsBlock {
		return nil, errors.New("This wallet is blocked. Can't perform any transactions.")
	}
	if txnType == "credit" {
		wallet.Balance += transaction.Amount
	}else {
		if wallet.Balance < transaction.Amount {
			return nil, errors.New("Wallet Balance is insufficient to deduct given amount")
		}
		wallet.Balance -= transaction.Amount
	}
	if _, err = iWalletService.UpdateWalletService(wallet); err != nil {
		return nil, err
	}
	transaction.WalletID = uint(walletID)
	transaction.Wallet = *wallet
	return transactionRepository.CreateTransaction(transaction)
}

func (transactionService *transactionService) UpdateActiveTransactionsService() error {
	return transactionRepository.UpdateAllActiveTransactions()
}

func (transactionService *transactionService) UpdateTransactionService(id int, transaction *Transaction) (*Transaction, error) {
	res, err := transactionRepository.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	updateTransactionsGivenFields(res, transaction)
	return transactionRepository.UpdateTransaction(res)
}

//func (transactionService *TransactionService) DeleteTransactionService(id int) error {
//	res, err := transactionService.GetTransactionById(id)
//	if err != nil {
//		return err
//	}
//	return transactionService.DeleteTransaction(res)
//}

func updateTransactionsGivenFields(u1 *Transaction, u2 *Transaction) {
	if len(u2.TxnType)!=0 {
		u1.TxnType = u2.TxnType
	}
	if u2.Amount!=0.0 {
		u1.Amount = u2.Amount
	}
	// if provided json value is true, then only update it
	if u2.Active {
		u1.Active = u2.Active
	}
}

