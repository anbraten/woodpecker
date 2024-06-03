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

// LogEntryType log entry type
//
// swagger:model LogEntryType
type LogEntryType int64

// for schema
var logEntryTypeEnum []interface{}

func init() {
	var res []LogEntryType
	if err := json.Unmarshal([]byte(`[0,1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logEntryTypeEnum = append(logEntryTypeEnum, v)
	}
}

func (m LogEntryType) validateLogEntryTypeEnum(path, location string, value LogEntryType) error {
	if err := validate.EnumCase(path, location, value, logEntryTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this log entry type
func (m LogEntryType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLogEntryTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this log entry type based on context it is used
func (m LogEntryType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
