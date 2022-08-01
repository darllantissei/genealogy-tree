// Code generated by MockGen. DO NOT EDIT.
// Source: /home/darllan/go/src/github.com/darllantissei/genealogy-tree/application/person/interface.go

// Package mock_person is a generated GoMock package.
package mock_person

import (
	context "context"
	reflect "reflect"

	model "github.com/darllantissei/genealogy-tree/application/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIPersonService is a mock of IPersonService interface.
type MockIPersonService struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonServiceMockRecorder
}

// MockIPersonServiceMockRecorder is the mock recorder for MockIPersonService.
type MockIPersonServiceMockRecorder struct {
	mock *MockIPersonService
}

// NewMockIPersonService creates a new mock instance.
func NewMockIPersonService(ctrl *gomock.Controller) *MockIPersonService {
	mock := &MockIPersonService{ctrl: ctrl}
	mock.recorder = &MockIPersonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPersonService) EXPECT() *MockIPersonServiceMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockIPersonService) Fetch(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockIPersonServiceMockRecorder) Fetch(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockIPersonService)(nil).Fetch), ctx, prsn)
}

// Record mocks base method.
func (m *MockIPersonService) Record(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Record", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Record indicates an expected call of Record.
func (mr *MockIPersonServiceMockRecorder) Record(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockIPersonService)(nil).Record), ctx, prsn)
}

// MockIPersonWriteDB is a mock of IPersonWriteDB interface.
type MockIPersonWriteDB struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonWriteDBMockRecorder
}

// MockIPersonWriteDBMockRecorder is the mock recorder for MockIPersonWriteDB.
type MockIPersonWriteDBMockRecorder struct {
	mock *MockIPersonWriteDB
}

// NewMockIPersonWriteDB creates a new mock instance.
func NewMockIPersonWriteDB(ctrl *gomock.Controller) *MockIPersonWriteDB {
	mock := &MockIPersonWriteDB{ctrl: ctrl}
	mock.recorder = &MockIPersonWriteDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPersonWriteDB) EXPECT() *MockIPersonWriteDBMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockIPersonWriteDB) Insert(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockIPersonWriteDBMockRecorder) Insert(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIPersonWriteDB)(nil).Insert), ctx, prsn)
}

// Update mocks base method.
func (m *MockIPersonWriteDB) Update(ctx context.Context, prsn model.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, prsn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPersonWriteDBMockRecorder) Update(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPersonWriteDB)(nil).Update), ctx, prsn)
}

// MockIPersonReadDB is a mock of IPersonReadDB interface.
type MockIPersonReadDB struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonReadDBMockRecorder
}

// MockIPersonReadDBMockRecorder is the mock recorder for MockIPersonReadDB.
type MockIPersonReadDBMockRecorder struct {
	mock *MockIPersonReadDB
}

// NewMockIPersonReadDB creates a new mock instance.
func NewMockIPersonReadDB(ctrl *gomock.Controller) *MockIPersonReadDB {
	mock := &MockIPersonReadDB{ctrl: ctrl}
	mock.recorder = &MockIPersonReadDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPersonReadDB) EXPECT() *MockIPersonReadDBMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIPersonReadDB) Get(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIPersonReadDBMockRecorder) Get(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPersonReadDB)(nil).Get), ctx, prsn)
}

// MockIPersonPersistenceDB is a mock of IPersonPersistenceDB interface.
type MockIPersonPersistenceDB struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonPersistenceDBMockRecorder
}

// MockIPersonPersistenceDBMockRecorder is the mock recorder for MockIPersonPersistenceDB.
type MockIPersonPersistenceDBMockRecorder struct {
	mock *MockIPersonPersistenceDB
}

// NewMockIPersonPersistenceDB creates a new mock instance.
func NewMockIPersonPersistenceDB(ctrl *gomock.Controller) *MockIPersonPersistenceDB {
	mock := &MockIPersonPersistenceDB{ctrl: ctrl}
	mock.recorder = &MockIPersonPersistenceDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPersonPersistenceDB) EXPECT() *MockIPersonPersistenceDBMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIPersonPersistenceDB) Get(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIPersonPersistenceDBMockRecorder) Get(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPersonPersistenceDB)(nil).Get), ctx, prsn)
}

// Insert mocks base method.
func (m *MockIPersonPersistenceDB) Insert(ctx context.Context, prsn model.Person) (model.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, prsn)
	ret0, _ := ret[0].(model.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockIPersonPersistenceDBMockRecorder) Insert(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIPersonPersistenceDB)(nil).Insert), ctx, prsn)
}

// Update mocks base method.
func (m *MockIPersonPersistenceDB) Update(ctx context.Context, prsn model.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, prsn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPersonPersistenceDBMockRecorder) Update(ctx, prsn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPersonPersistenceDB)(nil).Update), ctx, prsn)
}
