// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLogLevelReader is a Reader for the GetLogLevel structure.
type GetLogLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /log-level] GetLogLevel", response, response.Code())
	}
}

// NewGetLogLevelOK creates a GetLogLevelOK with default headers values
func NewGetLogLevelOK() *GetLogLevelOK {
	return &GetLogLevelOK{}
}

/*
GetLogLevelOK describes a response with status code 200, with default header values.

OK
*/
type GetLogLevelOK struct {
	Payload *GetLogLevelOKBody
}

// IsSuccess returns true when this get log level o k response has a 2xx status code
func (o *GetLogLevelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log level o k response has a 3xx status code
func (o *GetLogLevelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log level o k response has a 4xx status code
func (o *GetLogLevelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log level o k response has a 5xx status code
func (o *GetLogLevelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log level o k response a status code equal to that given
func (o *GetLogLevelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get log level o k response
func (o *GetLogLevelOK) Code() int {
	return 200
}

func (o *GetLogLevelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log-level][%d] getLogLevelOK %s", 200, payload)
}

func (o *GetLogLevelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /log-level][%d] getLogLevelOK %s", 200, payload)
}

func (o *GetLogLevelOK) GetPayload() *GetLogLevelOKBody {
	return o.Payload
}

func (o *GetLogLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLogLevelOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLogLevelOKBody get log level o k body
swagger:model GetLogLevelOKBody
*/
type GetLogLevelOKBody struct {
	GetLogLevelOKBodyAllOf0

	// log level
	LogLevel string `json:"log-level,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetLogLevelOKBody) UnmarshalJSON(raw []byte) error {
	// GetLogLevelOKBodyAO0
	var getLogLevelOKBodyAO0 GetLogLevelOKBodyAllOf0
	if err := swag.ReadJSON(raw, &getLogLevelOKBodyAO0); err != nil {
		return err
	}
	o.GetLogLevelOKBodyAllOf0 = getLogLevelOKBodyAO0

	// GetLogLevelOKBodyAO1
	var dataGetLogLevelOKBodyAO1 struct {
		LogLevel string `json:"log-level,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetLogLevelOKBodyAO1); err != nil {
		return err
	}

	o.LogLevel = dataGetLogLevelOKBodyAO1.LogLevel

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetLogLevelOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getLogLevelOKBodyAO0, err := swag.WriteJSON(o.GetLogLevelOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getLogLevelOKBodyAO0)
	var dataGetLogLevelOKBodyAO1 struct {
		LogLevel string `json:"log-level,omitempty"`
	}

	dataGetLogLevelOKBodyAO1.LogLevel = o.LogLevel

	jsonDataGetLogLevelOKBodyAO1, errGetLogLevelOKBodyAO1 := swag.WriteJSON(dataGetLogLevelOKBodyAO1)
	if errGetLogLevelOKBodyAO1 != nil {
		return nil, errGetLogLevelOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetLogLevelOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get log level o k body
func (o *GetLogLevelOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetLogLevelOKBodyAllOf0

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get log level o k body based on the context it is used
func (o *GetLogLevelOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetLogLevelOKBodyAllOf0
	if err := o.GetLogLevelOKBodyAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetLogLevelOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLogLevelOKBody) UnmarshalBinary(b []byte) error {
	var res GetLogLevelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLogLevelOKBodyAllOf0 get log level o k body all of0
swagger:model GetLogLevelOKBodyAllOf0
*/
type GetLogLevelOKBodyAllOf0 string

// Validate validates this get log level o k body all of0
func (o GetLogLevelOKBodyAllOf0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get log level o k body all of0 based on context it is used
func (o GetLogLevelOKBodyAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
