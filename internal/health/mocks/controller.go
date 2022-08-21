// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockdatabase is a mock of database interface.
type Mockdatabase struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseMockRecorder
}

// MockdatabaseMockRecorder is the mock recorder for Mockdatabase.
type MockdatabaseMockRecorder struct {
	mock *Mockdatabase
}

// NewMockdatabase creates a new mock instance.
func NewMockdatabase(ctrl *gomock.Controller) *Mockdatabase {
	mock := &Mockdatabase{ctrl: ctrl}
	mock.recorder = &MockdatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockdatabase) EXPECT() *MockdatabaseMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *Mockdatabase) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockdatabaseMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*Mockdatabase)(nil).Ping), arg0)
}