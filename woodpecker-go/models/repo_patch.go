// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepoPatch repo patch
//
// swagger:model RepoPatch
type RepoPatch struct {

	// allow deploy
	AllowDeploy bool `json:"allow_deploy,omitempty"`

	// allow pr
	AllowPr bool `json:"allow_pr,omitempty"`

	// cancel previous pipeline events
	CancelPreviousPipelineEvents []WebhookEvent `json:"cancel_previous_pipeline_events"`

	// config file
	ConfigFile string `json:"config_file,omitempty"`

	// gated
	Gated bool `json:"gated,omitempty"`

	// netrc only trusted
	NetrcOnlyTrusted bool `json:"netrc_only_trusted,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`

	// trusted
	Trusted bool `json:"trusted,omitempty"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this repo patch
func (m *RepoPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancelPreviousPipelineEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoPatch) validateCancelPreviousPipelineEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.CancelPreviousPipelineEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.CancelPreviousPipelineEvents); i++ {

		if err := m.CancelPreviousPipelineEvents[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancel_previous_pipeline_events" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancel_previous_pipeline_events" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this repo patch based on the context it is used
func (m *RepoPatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCancelPreviousPipelineEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepoPatch) contextValidateCancelPreviousPipelineEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CancelPreviousPipelineEvents); i++ {

		if swag.IsZero(m.CancelPreviousPipelineEvents[i]) { // not required
			return nil
		}

		if err := m.CancelPreviousPipelineEvents[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancel_previous_pipeline_events" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cancel_previous_pipeline_events" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepoPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoPatch) UnmarshalBinary(b []byte) error {
	var res RepoPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
