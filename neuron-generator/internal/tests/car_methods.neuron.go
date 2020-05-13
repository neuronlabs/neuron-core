// Code generated by neuron/generator. DO NOT EDIT.
// This file was generated at:
// Wed, 13 May 2020 12:22:51 +0200

package tests

import (
	"github.com/neuronlabs/neuron/errors"
	"github.com/neuronlabs/neuron/mapping"
)

// Compile time check if Car implements mapping.Model interface.
var _ mapping.Model = &Car{}

// NeuronCollectionName implements mapping.Model interface method.
// Returns the name of the collection for the 'Car'.
func (c *Car) NeuronCollectionName() string {
	return "cars"
}

// IsPrimaryKeyZero implements query.Model interface method.
func (c *Car) IsPrimaryKeyZero() bool {
	return c.ID == nil
}

// GetPrimaryKeyValue implements query.Model interface method.
func (c *Car) GetPrimaryKeyValue() interface{} {
	return c.ID
}

// GetPrimaryKeyHashableValue implements query.Model interface method.
func (c *Car) GetPrimaryKeyHashableValue() interface{} {
	if c.ID == nil {
		return c.ID
	}
	return *c.ID
}

// GetPrimaryKeyZeroValue implements query.Model interface method.
func (c *Car) GetPrimaryKeyZeroValue() interface{} {
	return (*[16]byte)(nil)
}

// SetPrimaryKey implements query.Model interface method.
func (c *Car) SetPrimaryKeyValue(value interface{}) error {
	if value == nil {
		c.ID = nil
		return nil
	}
	if v, ok := value.(*[16]byte); ok {
		c.ID = v
		return nil
	}
	return errors.Newf(mapping.ClassInvalidFieldValue, "provided invalid value: '%T' for the primary field for model: '%T'",
		value, c)
}

// Compile time check if Car implements mapping.Fielder interface.
var _ mapping.Fielder = &Car{}

// GetFieldZeroValue implements mapping.Fielder interface.s
func (c *Car) GetFieldZeroValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 1: // Plates
		return "", nil
	default:
		return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
}

// IsFieldZero implements mapping.Fielder interface.
func (c *Car) IsFieldZero(field *mapping.StructField) (bool, error) {
	switch field.Index[0] {
	case 1: // Plates
		return c.Plates == "", nil
	}
	return false, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
}

// SetFieldZeroValue implements mapping.Fielder interface.s
func (c *Car) SetFieldZeroValue(field *mapping.StructField) error {
	switch field.Index[0] {
	case 1: // Plates
		c.Plates = ""
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field name: '%s'", field.Name())
	}
	return nil
}

// GetHashableFieldValue implements mapping.Fielder interface.
func (c *Car) GetHashableFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 1: // Plates
		return c.Plates, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: '%s'", field.Name(), c)
}

// GetFieldValue implements mapping.Fielder interface.
func (c *Car) GetFieldValue(field *mapping.StructField) (interface{}, error) {
	switch field.Index[0] {
	case 1: // Plates
		return c.Plates, nil
	}
	return nil, errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for given model: '%s'", field.Name(), c)
}

// SetFieldValue implements mapping.Fielder interface.
func (c *Car) SetFieldValue(field *mapping.StructField, value interface{}) (err error) {
	switch field.Index[0] {
	case 1: // Plates
		if v, ok := value.(string); ok {
			c.Plates = v
			return nil
		}
		// Check alternate types for the Plates.
		if v, ok := value.([]byte); ok {
			c.Plates = string(v)
			return nil
		}
		return errors.Newf(mapping.ClassInvalidFieldValue, "provided invalid field type: '%T' for the field: %s", value, field.Name())
	default:
		return errors.Newf(mapping.ClassInvalidModelField, "provided invalid field: '%s' for the model: 'Car'", field.Name())
	}
}
