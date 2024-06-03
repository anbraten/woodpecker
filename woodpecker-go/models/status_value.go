// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StatusValue status value
//
// swagger:model StatusValue
type StatusValue string

func NewStatusValue(value StatusValue) *StatusValue {
	return &value
}

// Pointer returns a pointer to a freshly-allocated StatusValue.
func (m StatusValue) Pointer() *StatusValue {
	return &m
}

const (

	// StatusValueSkipped captures enum value "skipped"
	StatusValueSkipped StatusValue = "skipped"

	// StatusValuePending captures enum value "pending"
	StatusValuePending StatusValue = "pending"

	// StatusValueRunning captures enum value "running"
	StatusValueRunning StatusValue = "running"

	// StatusValueSuccess captures enum value "success"
	StatusValueSuccess StatusValue = "success"

	// StatusValueFailure captures enum value "failure"
	StatusValueFailure StatusValue = "failure"

	// StatusValueKilled captures enum value "killed"
	StatusValueKilled StatusValue = "killed"

	// StatusValueError captures enum value "error"
	StatusValueError StatusValue = "error"

	// StatusValueBlocked captures enum value "blocked"
	StatusValueBlocked StatusValue = "blocked"

	// StatusValueDeclined captures enum value "declined"
	StatusValueDeclined StatusValue = "declined"

	// StatusValueCreated captures enum value "created"
	StatusValueCreated StatusValue = "created"
)

// for schema
var statusValueEnum []interface{}

func init() {
	var res []StatusValue
	if err := json.Unmarshal([]byte(`["skipped","pending","running","success","failure","killed","error","blocked","declined","created"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusValueEnum = append(statusValueEnum, v)
	}
}

func (m StatusValue) validateStatusValueEnum(path, location string, value StatusValue) error {
	if err := validate.EnumCase(path, location, value, statusValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this status value
func (m StatusValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this status value based on context it is used
func (m StatusValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
