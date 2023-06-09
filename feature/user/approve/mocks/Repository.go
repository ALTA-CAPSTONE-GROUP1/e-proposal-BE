// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	approve "github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/user/approve"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// SelectSubmissionAprrove provides a mock function with given fields: userID, search
func (_m *Repository) SelectSubmissionAprrove(userID string, search approve.GetAllQueryParams) ([]approve.Core, error) {
	ret := _m.Called(userID, search)

	var r0 []approve.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string, approve.GetAllQueryParams) ([]approve.Core, error)); ok {
		return rf(userID, search)
	}
	if rf, ok := ret.Get(0).(func(string, approve.GetAllQueryParams) []approve.Core); ok {
		r0 = rf(userID, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]approve.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string, approve.GetAllQueryParams) error); ok {
		r1 = rf(userID, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectSubmissionById provides a mock function with given fields: userID, id
func (_m *Repository) SelectSubmissionById(userID string, id int) (approve.Core, error) {
	ret := _m.Called(userID, id)

	var r0 approve.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (approve.Core, error)); ok {
		return rf(userID, id)
	}
	if rf, ok := ret.Get(0).(func(string, int) approve.Core); ok {
		r0 = rf(userID, id)
	} else {
		r0 = ret.Get(0).(approve.Core)
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(userID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateApprove provides a mock function with given fields: userID, id, input
func (_m *Repository) UpdateApprove(userID string, id int, input approve.Core) error {
	ret := _m.Called(userID, id, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, approve.Core) error); ok {
		r0 = rf(userID, id, input)
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
