// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	position "github.com/ALTA-CAPSTONE-GROUP1/e-proposal-BE/feature/admin/position"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// AddPositionLogic provides a mock function with given fields: newPosition
func (_m *UseCase) AddPositionLogic(newPosition position.Core) error {
	ret := _m.Called(newPosition)

	var r0 error
	if rf, ok := ret.Get(0).(func(position.Core) error); ok {
		r0 = rf(newPosition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPositionsLogic provides a mock function with given fields: limit, offset, search
func (_m *UseCase) GetPositionsLogic(limit int, offset int, search string) ([]position.Core, error) {
	ret := _m.Called(limit, offset, search)

	var r0 []position.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]position.Core, error)); ok {
		return rf(limit, offset, search)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []position.Core); ok {
		r0 = rf(limit, offset, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]position.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(limit, offset, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
