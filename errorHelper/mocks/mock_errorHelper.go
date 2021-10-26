// Code generated by MockGen. DO NOT EDIT.
// Source: ../errorHelper/errorHelper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// As mocks base method.
func (m *MockHelper) As(err error, target interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "As", err, target)
	ret0, _ := ret[0].(bool)
	return ret0
}

// As indicates an expected call of As.
func (mr *MockHelperMockRecorder) As(err, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "As", reflect.TypeOf((*MockHelper)(nil).As), err, target)
}

// Cause mocks base method.
func (m *MockHelper) Cause(err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cause", err)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cause indicates an expected call of Cause.
func (mr *MockHelperMockRecorder) Cause(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cause", reflect.TypeOf((*MockHelper)(nil).Cause), err)
}

// Is mocks base method.
func (m *MockHelper) Is(err, target error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Is", err, target)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Is indicates an expected call of Is.
func (mr *MockHelperMockRecorder) Is(err, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Is", reflect.TypeOf((*MockHelper)(nil).Is), err, target)
}

// New mocks base method.
func (m *MockHelper) New(message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// New indicates an expected call of New.
func (mr *MockHelperMockRecorder) New(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockHelper)(nil).New), message)
}

// SprintErrorWithStack mocks base method.
func (m *MockHelper) SprintErrorWithStack(err error) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SprintErrorWithStack", err)
	ret0, _ := ret[0].(string)
	return ret0
}

// SprintErrorWithStack indicates an expected call of SprintErrorWithStack.
func (mr *MockHelperMockRecorder) SprintErrorWithStack(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SprintErrorWithStack", reflect.TypeOf((*MockHelper)(nil).SprintErrorWithStack), err)
}

// WithStack mocks base method.
func (m *MockHelper) WithStack(err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStack", err)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithStack indicates an expected call of WithStack.
func (mr *MockHelperMockRecorder) WithStack(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStack", reflect.TypeOf((*MockHelper)(nil).WithStack), err)
}

// Wrap mocks base method.
func (m *MockHelper) Wrap(err error, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wrap", err, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wrap indicates an expected call of Wrap.
func (mr *MockHelperMockRecorder) Wrap(err, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wrap", reflect.TypeOf((*MockHelper)(nil).Wrap), err, message)
}

// Wrapf mocks base method.
func (m *MockHelper) Wrapf(err error, format string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Wrapf", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wrapf indicates an expected call of Wrapf.
func (mr *MockHelperMockRecorder) Wrapf(err, format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wrapf", reflect.TypeOf((*MockHelper)(nil).Wrapf), varargs...)
}