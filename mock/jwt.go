// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/davidchristie/identity/jwt (interfaces: JWT)

// Package mock is a generated GoMock package.
package mock

import (
	entity "github.com/davidchristie/identity/entity"
	jwt "github.com/davidchristie/identity/jwt"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJWT is a mock of JWT interface
type MockJWT struct {
	ctrl     *gomock.Controller
	recorder *MockJWTMockRecorder
}

// MockJWTMockRecorder is the mock recorder for MockJWT
type MockJWTMockRecorder struct {
	mock *MockJWT
}

// NewMockJWT creates a new mock instance
func NewMockJWT(ctrl *gomock.Controller) *MockJWT {
	mock := &MockJWT{ctrl: ctrl}
	mock.recorder = &MockJWTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJWT) EXPECT() *MockJWTMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockJWT) Parse(arg0 string) (*entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].(*entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockJWTMockRecorder) Parse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockJWT)(nil).Parse), arg0)
}

// SignedString mocks base method
func (m *MockJWT) SignedString(arg0 *jwt.SignedStringInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignedString", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignedString indicates an expected call of SignedString
func (mr *MockJWTMockRecorder) SignedString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignedString", reflect.TypeOf((*MockJWT)(nil).SignedString), arg0)
}
