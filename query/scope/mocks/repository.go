// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/kucjac/jsonapi/mapping"
	"github.com/kucjac/jsonapi/repositories"
	mock "github.com/stretchr/testify/mock"
)
import scope "github.com/kucjac/jsonapi/query/scope"

var _ repositories.Repository = &Repository{}

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 *scope.Scope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*scope.Scope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *Repository) Delete(_a0 *scope.Scope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*scope.Scope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *Repository) Get(_a0 *scope.Scope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*scope.Scope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: _a0
func (_m *Repository) List(_a0 *scope.Scope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*scope.Scope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Patch provides a mock function with given fields: _a0
func (_m *Repository) Patch(_a0 *scope.Scope) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*scope.Scope) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (r *Repository) RepositoryName() string {
	return "scope-mock"
}

func (r *Repository) New(*mapping.ModelStruct) interface{} {
	return &Repository{}
}
