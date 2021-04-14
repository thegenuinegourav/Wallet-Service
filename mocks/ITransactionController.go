// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ITransactionController is an autogenerated mock type for the ITransactionController type
type ITransactionController struct {
	mock.Mock
}

// GetActiveTransactions provides a mock function with given fields: w, r
func (_m *ITransactionController) GetActiveTransactions(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetTransaction provides a mock function with given fields: w, r
func (_m *ITransactionController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetTransactions provides a mock function with given fields: w, r
func (_m *ITransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// GetTransactionsByWalletId provides a mock function with given fields: w, r
func (_m *ITransactionController) GetTransactionsByWalletId(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// PostTransaction provides a mock function with given fields: w, r
func (_m *ITransactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// PutTransaction provides a mock function with given fields: w, r
func (_m *ITransactionController) PutTransaction(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UpdateActiveTransactions provides a mock function with given fields: w, r
func (_m *ITransactionController) UpdateActiveTransactions(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}
