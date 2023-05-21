// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// UploadInterface is an autogenerated mock type for the UploadInterface type
type UploadInterface struct {
	mock.Mock
}

// UploadFile provides a mock function with given fields: fileContents, path
func (_m *UploadInterface) UploadFile(fileContents *multipart.FileHeader, path string) ([]string, error) {
	ret := _m.Called(fileContents, path)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, string) ([]string, error)); ok {
		return rf(fileContents, path)
	}
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, string) []string); ok {
		r0 = rf(fileContents, path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*multipart.FileHeader, string) error); ok {
		r1 = rf(fileContents, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUploadInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploadInterface creates a new instance of UploadInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploadInterface(t mockConstructorTestingTNewUploadInterface) *UploadInterface {
	mock := &UploadInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
