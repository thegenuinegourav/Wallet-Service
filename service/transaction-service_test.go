package service

import (
	"github.com/WalletService/mocks"
	"github.com/WalletService/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTransactionService(t *testing.T) {
	transactionRepository := new(mocks.ITransactionRepository)
	walletService := new(mocks.IWalletService)
	transactionService := NewTransactionService(transactionRepository, walletService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.Transaction{
		TxnType: "credit",
		Amount: 2000,
		WalletID: 1,
	}
	transactionRepository.On("GetTransactionById", 1).Return(&expectedResult1, nil)
	actualResult1, _ := transactionService.GetTransactionService(1)
	assert.Equal(t, &expectedResult1, actualResult1)
}

func TestPostTransactionService(t *testing.T) {
	transactionRepository := new(mocks.ITransactionRepository)
	walletService := new(mocks.IWalletService)
	transactionService := NewTransactionService(transactionRepository, walletService)

	t.Logf("Running Positive Case with valid output..........")
	walletExpected := model.Wallet{
		IsBlock: false,
	}
	expectedResult1 := model.Transaction{
		TxnType: "credit",
		Amount: 2000,
		WalletID: 1,
	}
	walletService.On("GetWalletService", 1, false).Return(&walletExpected, nil)
	walletService.On("UpdateWalletService", &walletExpected).Return(&walletExpected, nil)
	transactionRepository.On("CreateTransaction", &expectedResult1).Return(&expectedResult1, nil)
	actualResult1, _ := transactionService.PostTransactionService(&expectedResult1, 1)
	assert.Equal(t, &expectedResult1, actualResult1)

	t.Logf("Running Negative Case with misspell txntype error..........")
	invalidInput := model.Transaction{
		TxnType: "debit-misspell",
		Amount: 2000,
	}
	_, err := transactionService.PostTransactionService(&invalidInput, 2)
	assert.NotNil(t, err)
}
