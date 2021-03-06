// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	redis "github.com/go-redis/redis/v8"
	mock "github.com/stretchr/testify/mock"
)

// ICacheEngine is an autogenerated mock type for the ICacheEngine type
type ICacheEngine struct {
	mock.Mock
}

// CheckConnection provides a mock function with given fields:
func (_m *ICacheEngine) CheckConnection() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCacheClient provides a mock function with given fields:
func (_m *ICacheEngine) GetCacheClient() *redis.Client {
	ret := _m.Called()

	var r0 *redis.Client
	if rf, ok := ret.Get(0).(func() *redis.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.Client)
		}
	}

	return r0
}
