// Code generated by MockGen. DO NOT EDIT.
// Source: rocket.go

// Package rocket is a generated GoMock package.
package rocket

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteRocket mocks base method.
func (m *MockStore) DeleteRocket(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRocket", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRocket indicates an expected call of DeleteRocket.
func (mr *MockStoreMockRecorder) DeleteRocket(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRocket", reflect.TypeOf((*MockStore)(nil).DeleteRocket), id)
}

// GetRocketByID mocks base method.
func (m *MockStore) GetRocketByID(id string) (Rocket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRocketByID", id)
	ret0, _ := ret[0].(Rocket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRocketByID indicates an expected call of GetRocketByID.
func (mr *MockStoreMockRecorder) GetRocketByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRocketByID", reflect.TypeOf((*MockStore)(nil).GetRocketByID), id)
}

// InsertRocket mocks base method.
func (m *MockStore) InsertRocket(rkt Rocket) (Rocket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRocket", rkt)
	ret0, _ := ret[0].(Rocket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRocket indicates an expected call of InsertRocket.
func (mr *MockStoreMockRecorder) InsertRocket(rkt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRocket", reflect.TypeOf((*MockStore)(nil).InsertRocket), rkt)
}
