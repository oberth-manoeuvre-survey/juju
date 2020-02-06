// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/txn (interfaces: Runner)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	txn "github.com/juju/txn"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// MaybePruneTransactions mocks base method
func (m *MockRunner) MaybePruneTransactions(arg0 txn.PruneOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePruneTransactions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybePruneTransactions indicates an expected call of MaybePruneTransactions
func (mr *MockRunnerMockRecorder) MaybePruneTransactions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePruneTransactions", reflect.TypeOf((*MockRunner)(nil).MaybePruneTransactions), arg0)
}

// ResumeTransactions mocks base method
func (m *MockRunner) ResumeTransactions() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeTransactions")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeTransactions indicates an expected call of ResumeTransactions
func (mr *MockRunnerMockRecorder) ResumeTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeTransactions", reflect.TypeOf((*MockRunner)(nil).ResumeTransactions))
}

// Run mocks base method
func (m *MockRunner) Run(arg0 txn.TransactionSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), arg0)
}

// RunTransaction mocks base method
func (m *MockRunner) RunTransaction(arg0 *txn.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction
func (mr *MockRunnerMockRecorder) RunTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockRunner)(nil).RunTransaction), arg0)
}
