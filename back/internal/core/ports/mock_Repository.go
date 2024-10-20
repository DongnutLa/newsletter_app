// Code generated by mockery v2.28.1. DO NOT EDIT.

package ports

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository[T interface{}, Q interface{}] struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, opts
func (_m *MockRepository[T, Q]) DeleteOne(ctx context.Context, opts DeleteOpts) (bool, error) {
	ret := _m.Called(ctx, opts)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DeleteOpts) (bool, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DeleteOpts) bool); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, DeleteOpts) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, opts, result, returnCount
func (_m *MockRepository[T, Q]) FindMany(ctx context.Context, opts FindManyOpts, result *[]T, returnCount bool) (*int64, error) {
	ret := _m.Called(ctx, opts, result, returnCount)

	var r0 *int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, FindManyOpts, *[]T, bool) (*int64, error)); ok {
		return rf(ctx, opts, result, returnCount)
	}
	if rf, ok := ret.Get(0).(func(context.Context, FindManyOpts, *[]T, bool) *int64); ok {
		r0 = rf(ctx, opts, result, returnCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, FindManyOpts, *[]T, bool) error); ok {
		r1 = rf(ctx, opts, result, returnCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, opts, result
func (_m *MockRepository[T, Q]) FindOne(ctx context.Context, opts FindOneOpts, result *T) error {
	ret := _m.Called(ctx, opts, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, FindOneOpts, *T) error); ok {
		r0 = rf(ctx, opts, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertOne provides a mock function with given fields: ctx, entity
func (_m *MockRepository[T, Q]) InsertOne(ctx context.Context, entity T) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, T) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOne provides a mock function with given fields: ctx, opts
func (_m *MockRepository[T, Q]) UpdateOne(ctx context.Context, opts UpdateOpts) (*T, error) {
	ret := _m.Called(ctx, opts)

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, UpdateOpts) (*T, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, UpdateOpts) *T); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, UpdateOpts) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository[T interface{}, Q interface{}](t mockConstructorTestingTNewMockRepository) *MockRepository[T, Q] {
	mock := &MockRepository[T, Q]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
