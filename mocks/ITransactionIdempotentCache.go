// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	model "github.com/WalletService/model"
	mock "github.com/stretchr/testify/mock"
)

// ITransactionIdempotentCache is an autogenerated mock type for the ITransactionIdempotentCache type
type ITransactionIdempotentCache struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *ITransactionIdempotentCache) Get(key string) *model.Transaction {
	ret := _m.Called(key)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(string) *model.Transaction); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	return r0
}

// GetIdempotencyKey provides a mock function with given fields: r
func (_m *ITransactionIdempotentCache) GetIdempotencyKey(r *http.Request) (string, error) {
	ret := _m.Called(r)

	var r0 string
	if rf, ok := ret.Get(0).(func(*http.Request) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value
func (_m *ITransactionIdempotentCache) Set(key string, value *model.Transaction) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *model.Transaction) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
