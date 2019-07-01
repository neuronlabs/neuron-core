// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/neuronlabs/neuron/mapping"
	"github.com/neuronlabs/neuron/repository"
)

func init() {
	repository.RegisterFactory(&Factory{})
}

// Factory is an autogenerated mock type for the Factory type.
type Factory struct {
	mock.Mock
}

// New provides a mock function with given fields: model.
func (_m *Factory) New(structer repository.ModelStructer, model *mapping.ModelStruct) (repository.Repository, error) {
	return &Repository{}, nil
}

// DriverName provides a mock repository name.
func (_m *Factory) DriverName() string {
	return "mocks"
}

// Close implements repository.Closer interface.
func (_m *Factory) Close(ctx context.Context, done chan<- interface{}) {
	done <- struct{}{}
}
