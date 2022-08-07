// Code generated by MockGen. DO NOT EDIT.
// Source: /home/darllan/go/src/github.com/darllantissei/genealogy-tree/application/common/interface.go

// Package mock_common is a generated GoMock package.
package mock_common

import (
	context "context"
	reflect "reflect"

	model "github.com/darllantissei/genealogy-tree/application/model"
	gomock "github.com/golang/mock/gomock"
)

// MockICommonService is a mock of ICommonService interface.
type MockICommonService struct {
	ctrl     *gomock.Controller
	recorder *MockICommonServiceMockRecorder
}

// MockICommonServiceMockRecorder is the mock recorder for MockICommonService.
type MockICommonServiceMockRecorder struct {
	mock *MockICommonService
}

// NewMockICommonService creates a new mock instance.
func NewMockICommonService(ctrl *gomock.Controller) *MockICommonService {
	mock := &MockICommonService{ctrl: ctrl}
	mock.recorder = &MockICommonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommonService) EXPECT() *MockICommonServiceMockRecorder {
	return m.recorder
}

// BuildError mocks base method.
func (m *MockICommonService) BuildError(ctx context.Context, msgErr []string) model.Returns {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildError", ctx, msgErr)
	ret0, _ := ret[0].(model.Returns)
	return ret0
}

// BuildError indicates an expected call of BuildError.
func (mr *MockICommonServiceMockRecorder) BuildError(ctx, msgErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildError", reflect.TypeOf((*MockICommonService)(nil).BuildError), ctx, msgErr)
}

// SliceExists mocks base method.
func (m *MockICommonService) SliceExists(slice, item interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliceExists", slice, item)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SliceExists indicates an expected call of SliceExists.
func (mr *MockICommonServiceMockRecorder) SliceExists(slice, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliceExists", reflect.TypeOf((*MockICommonService)(nil).SliceExists), slice, item)
}
