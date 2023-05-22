// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	approve "github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/user/approve"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetSubmissionAprrove provides a mock function with given fields: userID, search
func (_m *UseCase) GetSubmissionAprrove(userID string, search approve.GetAllQueryParams) ([]approve.Core, error) {
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

// GetSubmissionById provides a mock function with given fields: userID, id
func (_m *UseCase) GetSubmissionById(userID string, id int) (approve.Core, error) {
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

// UpdateApprove provides a mock function with given fields: userID, id, updateInput
func (_m *UseCase) UpdateApprove(userID string, id int, updateInput approve.Core) error {
	ret := _m.Called(userID, id, updateInput)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, approve.Core) error); ok {
		r0 = rf(userID, id, updateInput)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}