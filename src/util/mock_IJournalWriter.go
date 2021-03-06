// Code generated by mockery v1.0.0. DO NOT EDIT.

package util

import (
	journald "github.com/ssgreg/journald"
	mock "github.com/stretchr/testify/mock"
)

// MockIJournalWriter is an autogenerated mock type for the IJournalWriter type
type MockIJournalWriter struct {
	mock.Mock
}

// Send provides a mock function with given fields: msg, p, fields
func (_m *MockIJournalWriter) Send(msg string, p journald.Priority, fields map[string]interface{}) error {
	ret := _m.Called(msg, p, fields)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, journald.Priority, map[string]interface{}) error); ok {
		r0 = rf(msg, p, fields)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
