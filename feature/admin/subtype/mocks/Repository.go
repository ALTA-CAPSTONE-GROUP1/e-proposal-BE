// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	subtype "github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/subtype"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// InsertSubType provides a mock function with given fields: req
func (_m *Repository) InsertSubType(req subtype.RepoData) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(subtype.RepoData) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}