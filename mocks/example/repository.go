// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"

	request "github.com/wudtichaikarun/poc-go-template-02/services/entities/request"
)

// ExampleRepository is an autogenerated mock type for the ExampleRepository type
type ExampleRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *ExampleRepository) Create(_a0 request.ExampleReq) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.ExampleReq) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *ExampleRepository) List() ([]models.Example, error) {
	ret := _m.Called()

	var r0 []models.Example
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Example, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Example); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Example)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExampleRepository creates a new instance of ExampleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExampleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExampleRepository {
	mock := &ExampleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
