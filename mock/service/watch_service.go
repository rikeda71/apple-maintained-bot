// Code generated by MockGen. DO NOT EDIT.
// Source: watch_service.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/s14t284/apple-maitained-bot/domain/model"
)

// MockWatchService is a mock of WatchService interface.
type MockWatchService struct {
	ctrl     *gomock.Controller
	recorder *MockWatchServiceMockRecorder
}

// MockWatchServiceMockRecorder is the mock recorder for MockWatchService.
type MockWatchServiceMockRecorder struct {
	mock *MockWatchService
}

// NewMockWatchService creates a new mock instance.
func NewMockWatchService(ctrl *gomock.Controller) *MockWatchService {
	mock := &MockWatchService{ctrl: ctrl}
	mock.recorder = &MockWatchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchService) EXPECT() *MockWatchServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockWatchService) Add(watch *model.Watch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", watch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockWatchServiceMockRecorder) Add(watch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockWatchService)(nil).Add), watch)
}

// Find mocks base method.
func (m *MockWatchService) Find(param *model.WatchRequestParam) (model.Watches, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", param)
	ret0, _ := ret[0].(model.Watches)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockWatchServiceMockRecorder) Find(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockWatchService)(nil).Find), param)
}

// FindAll mocks base method.
func (m *MockWatchService) FindAll() (model.Watches, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(model.Watches)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockWatchServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockWatchService)(nil).FindAll))
}

// FindByURL mocks base method.
func (m *MockWatchService) FindByURL(url string) (*model.Watch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByURL", url)
	ret0, _ := ret[0].(*model.Watch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByURL indicates an expected call of FindByURL.
func (mr *MockWatchServiceMockRecorder) FindByURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByURL", reflect.TypeOf((*MockWatchService)(nil).FindByURL), url)
}

// IsExist mocks base method.
func (m *MockWatchService) IsExist(watch *model.Watch) (bool, uint, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExist", watch)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// IsExist indicates an expected call of IsExist.
func (mr *MockWatchServiceMockRecorder) IsExist(watch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExist", reflect.TypeOf((*MockWatchService)(nil).IsExist), watch)
}

// Remove mocks base method.
func (m *MockWatchService) Remove(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockWatchServiceMockRecorder) Remove(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockWatchService)(nil).Remove), id)
}

// Update mocks base method.
func (m *MockWatchService) Update(watch *model.Watch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", watch)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWatchServiceMockRecorder) Update(watch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWatchService)(nil).Update), watch)
}

// UpdateAllSoldTemporary mocks base method.
func (m *MockWatchService) UpdateAllSoldTemporary() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllSoldTemporary")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAllSoldTemporary indicates an expected call of UpdateAllSoldTemporary.
func (mr *MockWatchServiceMockRecorder) UpdateAllSoldTemporary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllSoldTemporary", reflect.TypeOf((*MockWatchService)(nil).UpdateAllSoldTemporary))
}