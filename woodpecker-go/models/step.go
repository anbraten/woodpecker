// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Step step
//
// swagger:model Step
type Step struct {

	// end time
	EndTime int64 `json:"end_time,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// exit code
	ExitCode int64 `json:"exit_code,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pid
	Pid int64 `json:"pid,omitempty"`

	// pipeline id
	PipelineID int64 `json:"pipeline_id,omitempty"`

	// ppid
	Ppid int64 `json:"ppid,omitempty"`

	// start time
	StartTime int64 `json:"start_time,omitempty"`

	// state
	State StatusValue `json:"state,omitempty"`

	// type
	Type StepType `json:"type,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this step
func (m *Step) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Step) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Step) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this step based on the context it is used
func (m *Step) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Step) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Step) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Step) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Step) UnmarshalBinary(b []byte) error {
	var res Step
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
