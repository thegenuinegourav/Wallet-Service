package repository

import (
	"github.com/jinzhu/gorm"
	. "github.com/WalletService/model"
)

type ITransactionRepository interface{
	GetTransactionById(id int) (*Transaction, error)
	GetTransactionsByWalletId(id int) (*[]Transaction, error)
	GetAllTransactions() (*[]Transaction, error)
	GetAllActiveTransactions() (*[]Transaction, error)
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	UpdateTransaction(transaction *Transaction) (*Transaction, error)
	UpdateAllActiveTransactions() error
	//DeleteTransaction(transaction *Transaction) error
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &transactionRepository{db}
}

func (transactionRepository *transactionRepository) GetTransactionById(id int) (*Transaction, error) {
	var transaction Transaction
	result := transactionRepository.DB.Preload("Wallet").Preload("Wallet.User").First(&transaction,id)
	return &transaction, result.Error
}

func (transactionRepository *transactionRepository) GetTransactionsByWalletId(id int) (*[]Transaction, error) {
	var transaction []Transaction
	result := transactionRepository.DB.Where("wallet_id = ?", id).Preload("Wallet").Preload("Wallet.User").Find(&transaction)
	return &transaction, result.Error
}

func (transactionRepository *transactionRepository) GetAllTransactions() (*[]Transaction, error) {
	var transaction []Transaction
	result := transactionRepository.DB.Preload("Wallet").Preload("Wallet.User").Find(&transaction)
	return &transaction, result.Error
}

func (transactionRepository *transactionRepository) GetAllActiveTransactions() (*[]Transaction, error) {
	var transaction []Transaction
	result := transactionRepository.DB.Where("active = ?",true).Preload("Wallet").Preload("Wallet.User").Find(&transaction)
	return &transaction, result.Error
}

func (transactionRepository *transactionRepository) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	result := transactionRepository.DB.Create(transaction)
	return transaction, result.Error
}

func (transactionRepository *transactionRepository) UpdateTransaction(transaction *Transaction) (*Transaction, error) {
	result := transactionRepository.DB.Save(transaction)
	return transaction, result.Error
}

func (transactionRepository *transactionRepository) UpdateAllActiveTransactions() error {
	result := transactionRepository.DB.Model(Transaction{}).Where("active = ?", true).Updates(map[string]interface{}{"active": false})
	return result.Error
}

//func (transactionRepository *TransactionRepository) DeleteTransaction(transaction *Transaction) error {
//	result := transactionRepository.DB.Delete(transaction)
//	return result.Error
//}












