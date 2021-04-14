package controller

import (
	"bytes"
	"encoding/json"
	"github.com/WalletService/mocks"
	"github.com/WalletService/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	transactionService := new(mocks.ITransactionService)
	transactionIdempotentCache := new(mocks.ITransactionIdempotentCache)
	transactionController := NewTransactionController(transactionService,transactionIdempotentCache)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.Transaction{
		TxnType: "credit",
		Amount : 3000,
		WalletID: 1,
	}
	transactionService.On("GetTransactionService", 1).Return(&expectedResult1, nil)
	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/transaction/1", nil)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/transaction/{id}", transactionController.GetTransaction)
	r.ServeHTTP(w, req)
	actualResult1 := model.Transaction{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)
}

func TestPostTransaction(t *testing.T) {
	transactionService := new(mocks.ITransactionService)
	transactionIdempotentCache := new(mocks.ITransactionIdempotentCache)
	transactionController := NewTransactionController(transactionService,transactionIdempotentCache)

	t.Logf("Running Positive Case with valid cached output..........")
	input := model.Transaction{
		TxnType: "credit",
		Amount : 3000,
	}
	cachedTransaction := model.Transaction{
		TxnType: "credit",
		Amount : 3000,
		WalletID: 1,
	}
	body, _ := json.Marshal(&input)
	req := httptest.NewRequest("POST", "http://localhost:8080/api/v1/wallet/1/transaction", bytes.NewBuffer(body))
	transactionIdempotentCache.On("GetIdempotencyKey", mock.AnythingOfType("*http.Request")).Return("test-key", nil)
	transactionIdempotentCache.On("Get", "test-key").Return(&cachedTransaction, nil)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet/{id}/transaction", transactionController.PostTransaction)
	r.ServeHTTP(w, req)
	actualResult1 := model.Transaction{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, cachedTransaction,actualResult1)

	t.Logf("Running Positive Case with valid output..........")
	input2 := model.Transaction{
		TxnType: "credit",
		Amount : 3000,
	}
	expectedResult2 := model.Transaction{
		TxnType: "credit",
		Amount : 3000,
		WalletID: 1,
	}
	body2, _ := json.Marshal(&input2)
	req = httptest.NewRequest("POST", "http://localhost:8080/api/v1/wallet/2/transaction", bytes.NewBuffer(body2))
	transactionIdempotentCache.On("GetIdempotencyKey", mock.AnythingOfType("*http.Request")).Return("test-key-2", nil)
	transactionIdempotentCache.On("Get", "test-key-2").Return(nil, nil)
	transactionIdempotentCache.On("Set", mock.Anything, mock.Anything).Return(nil)
	transactionService.On("PostTransactionService", &input2, 2).Return(&expectedResult2, nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/wallet/{id}/transaction", transactionController.PostTransaction)
	r.ServeHTTP(w, req)
	actualResult2 := model.Transaction{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)
}
