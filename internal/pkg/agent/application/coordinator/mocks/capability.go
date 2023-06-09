// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by mockery v2.20.0. DO NOT EDIT.

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package mocks

import mock "github.com/stretchr/testify/mock"

// Capability is an autogenerated mock type for the Capability type
type Capability struct {
	mock.Mock
}

type Capability_Expecter struct {
	mock *mock.Mock
}

func (_m *Capability) EXPECT() *Capability_Expecter {
	return &Capability_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: _a0
func (_m *Capability) Apply(_a0 interface{}) (interface{}, error) {
	ret := _m.Called(_a0)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Capability_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type Capability_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - _a0 interface{}
func (_e *Capability_Expecter) Apply(_a0 interface{}) *Capability_Apply_Call {
	return &Capability_Apply_Call{Call: _e.mock.On("Apply", _a0)}
}

func (_c *Capability_Apply_Call) Run(run func(_a0 interface{})) *Capability_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Capability_Apply_Call) Return(_a0 interface{}, _a1 error) *Capability_Apply_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Capability_Apply_Call) RunAndReturn(run func(interface{}) (interface{}, error)) *Capability_Apply_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCapability interface {
	mock.TestingT
	Cleanup(func())
}

// NewCapability creates a new instance of Capability. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCapability(t mockConstructorTestingTNewCapability) *Capability {
	mock := &Capability{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
