// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-tenant-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tenant "github.com/ozoncp/ocp-tenant-api/internal/tenant"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddTenants mocks base method.
func (m *MockRepo) AddTenants(arg0 []tenant.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTenants", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTenants indicates an expected call of AddTenants.
func (mr *MockRepoMockRecorder) AddTenants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTenants", reflect.TypeOf((*MockRepo)(nil).AddTenants), arg0)
}

// DescribeTenant mocks base method.
func (m *MockRepo) DescribeTenant(arg0 uint64) (*tenant.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTenant", arg0)
	ret0, _ := ret[0].(*tenant.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTenant indicates an expected call of DescribeTenant.
func (mr *MockRepoMockRecorder) DescribeTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTenant", reflect.TypeOf((*MockRepo)(nil).DescribeTenant), arg0)
}

// ListTenants mocks base method.
func (m *MockRepo) ListTenants(arg0, arg1 uint64) ([]tenant.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenants", arg0, arg1)
	ret0, _ := ret[0].([]tenant.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenants indicates an expected call of ListTenants.
func (mr *MockRepoMockRecorder) ListTenants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenants", reflect.TypeOf((*MockRepo)(nil).ListTenants), arg0, arg1)
}

// RemoveTenant mocks base method.
func (m *MockRepo) RemoveTenant(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTenant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTenant indicates an expected call of RemoveTenant.
func (mr *MockRepoMockRecorder) RemoveTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTenant", reflect.TypeOf((*MockRepo)(nil).RemoveTenant), arg0)
}