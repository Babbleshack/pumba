// Code generated by mockery v1.0.0
package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, random
func (_m *Command) Run(ctx context.Context, random bool) error {
	ret := _m.Called(ctx, random)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, bool) error); ok {
		r0 = rf(ctx, random)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}