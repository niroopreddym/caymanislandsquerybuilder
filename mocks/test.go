// Code generated by MockGen. DO NOT EDIT.
// Source: querybuilder_interface.go

// Package mock_querybuilder is a generated GoMock package.
package mock_querybuilder

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIQueryBuilder is a mock of IQueryBuilder interface.
type MockIQueryBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockIQueryBuilderMockRecorder
}

// MockIQueryBuilderMockRecorder is the mock recorder for MockIQueryBuilder.
type MockIQueryBuilderMockRecorder struct {
	mock *MockIQueryBuilder
}

// NewMockIQueryBuilder creates a new mock instance.
func NewMockIQueryBuilder(ctrl *gomock.Controller) *MockIQueryBuilder {
	mock := &MockIQueryBuilder{ctrl: ctrl}
	mock.recorder = &MockIQueryBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIQueryBuilder) EXPECT() *MockIQueryBuilderMockRecorder {
	return m.recorder
}

// GetQueryPattern mocks base method.
func (m *MockIQueryBuilder) GetQueryPattern(queryData []byte) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryPattern", queryData)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetQueryPattern indicates an expected call of GetQueryPattern.
func (mr *MockIQueryBuilderMockRecorder) GetQueryPattern(queryData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryPattern", reflect.TypeOf((*MockIQueryBuilder)(nil).GetQueryPattern), queryData)
}