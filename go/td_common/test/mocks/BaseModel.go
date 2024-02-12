// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BaseModel is an autogenerated mock type for the BaseModel type
type BaseModel struct {
	mock.Mock
}

// GetID provides a mock function with given fields:
func (_m *BaseModel) GetID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetID provides a mock function with given fields: id
func (_m *BaseModel) SetID(id string) {
	_m.Called(id)
}

// NewBaseModel creates a new instance of BaseModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaseModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaseModel {
	mock := &BaseModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
