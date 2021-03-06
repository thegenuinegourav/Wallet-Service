// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "github.com/WalletService/model"
	mock "github.com/stretchr/testify/mock"
)

// IWalletRepository is an autogenerated mock type for the IWalletRepository type
type IWalletRepository struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: wallet
func (_m *IWalletRepository) CreateWallet(wallet *model.Wallet) (*model.Wallet, error) {
	ret := _m.Called(wallet)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(*model.Wallet) *model.Wallet); ok {
		r0 = rf(wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Wallet) error); ok {
		r1 = rf(wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletById provides a mock function with given fields: id
func (_m *IWalletRepository) GetWalletById(id int) (*model.Wallet, error) {
	ret := _m.Called(id)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(int) *model.Wallet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletByUserId provides a mock function with given fields: userID
func (_m *IWalletRepository) GetWalletByUserId(userID int) (*model.Wallet, error) {
	ret := _m.Called(userID)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(int) *model.Wallet); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWallet provides a mock function with given fields: wallet
func (_m *IWalletRepository) UpdateWallet(wallet *model.Wallet) (*model.Wallet, error) {
	ret := _m.Called(wallet)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(*model.Wallet) *model.Wallet); ok {
		r0 = rf(wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Wallet) error); ok {
		r1 = rf(wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
