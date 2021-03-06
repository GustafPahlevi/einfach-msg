// Code generated by MockGen. DO NOT EDIT.
// Source: mongo.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/GustafPahlevi/go-simple-svc/model"

	gomock "github.com/golang/mock/gomock"
)

// MockCollection is a mock of Collection interface
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
}

// MockCollectionMockRecorder is the mock recorder for MockCollection
type MockCollectionMockRecorder struct {
	mock *MockCollection
}

// NewMockCollection creates a new mock instance
func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockCollection) Insert(request model.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionMockRecorder) Insert(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollection)(nil).Insert), request)
}

// Get mocks base method
func (m *MockCollection) Get() ([]*model.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]*model.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCollectionMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCollection)(nil).Get))
}
