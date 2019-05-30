// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	query "github.com/neuronlabs/neuron/query"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx, s
func (_m *Repository) Begin(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close closes the repository connection
func (_m *Repository) Close(ctx context.Context) error {
	return nil
}

// Commit provides a mock function with given fields: ctx, s
func (_m *Repository) Commit(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, s
func (_m *Repository) Create(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, s
func (_m *Repository) Delete(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, s
func (_m *Repository) Get(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, s
func (_m *Repository) List(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Patch provides a mock function with given fields: ctx, s
func (_m *Repository) Patch(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx, s
func (_m *Repository) Rollback(ctx context.Context, s *query.Scope) error {
	ret := _m.Called(ctx, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *query.Scope) error); ok {
		r0 = rf(ctx, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryName provides a mock function with given fields:
func (_m *Repository) RepositoryName() string {
	return "mock"
}
