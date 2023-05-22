// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	approve "github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/approve"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetSubmissionByHyperApproval provides a mock function with given fields: userID, id, token
func (_m *UseCase) GetSubmissionByHyperApproval(userID string, id int, token string) (approve.GetSubmissionByIDCore, error) {
	ret := _m.Called(userID, id, token)

	var r0 approve.GetSubmissionByIDCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, string) (approve.GetSubmissionByIDCore, error)); ok {
		return rf(userID, id, token)
	}
	if rf, ok := ret.Get(0).(func(string, int, string) approve.GetSubmissionByIDCore); ok {
		r0 = rf(userID, id, token)
	} else {
		r0 = ret.Get(0).(approve.GetSubmissionByIDCore)
	}

	if rf, ok := ret.Get(1).(func(string, int, string) error); ok {
		r1 = rf(userID, id, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByHyperApproval provides a mock function with given fields: userID, updateInput
func (_m *UseCase) UpdateByHyperApproval(userID string, updateInput approve.Core) error {
	ret := _m.Called(userID, updateInput)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, approve.Core) error); ok {
		r0 = rf(userID, updateInput)
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