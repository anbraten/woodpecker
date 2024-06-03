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

// GetVersionReader is a Reader for the GetVersion structure.
type GetVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /version] GetVersion", response, response.Code())
	}
}

// NewGetVersionOK creates a GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

/*
GetVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetVersionOK struct {
	Payload *GetVersionOKBody
}

// IsSuccess returns true when this get version o k response has a 2xx status code
func (o *GetVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version o k response has a 3xx status code
func (o *GetVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version o k response has a 4xx status code
func (o *GetVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version o k response has a 5xx status code
func (o *GetVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get version o k response a status code equal to that given
func (o *GetVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get version o k response
func (o *GetVersionOK) Code() int {
	return 200
}

func (o *GetVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getVersionOK %s", 200, payload)
}

func (o *GetVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /version][%d] getVersionOK %s", 200, payload)
}

func (o *GetVersionOK) GetPayload() *GetVersionOKBody {
	return o.Payload
}

func (o *GetVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetVersionOKBody get version o k body
swagger:model GetVersionOKBody
*/
type GetVersionOKBody struct {
	GetVersionOKBodyAllOf0

	// source
	Source string `json:"source,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetVersionOKBody) UnmarshalJSON(raw []byte) error {
	// GetVersionOKBodyAO0
	var getVersionOKBodyAO0 GetVersionOKBodyAllOf0
	if err := swag.ReadJSON(raw, &getVersionOKBodyAO0); err != nil {
		return err
	}
	o.GetVersionOKBodyAllOf0 = getVersionOKBodyAO0

	// GetVersionOKBodyAO1
	var dataGetVersionOKBodyAO1 struct {
		Source string `json:"source,omitempty"`

		Version string `json:"version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetVersionOKBodyAO1); err != nil {
		return err
	}

	o.Source = dataGetVersionOKBodyAO1.Source

	o.Version = dataGetVersionOKBodyAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetVersionOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getVersionOKBodyAO0, err := swag.WriteJSON(o.GetVersionOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getVersionOKBodyAO0)
	var dataGetVersionOKBodyAO1 struct {
		Source string `json:"source,omitempty"`

		Version string `json:"version,omitempty"`
	}

	dataGetVersionOKBodyAO1.Source = o.Source

	dataGetVersionOKBodyAO1.Version = o.Version

	jsonDataGetVersionOKBodyAO1, errGetVersionOKBodyAO1 := swag.WriteJSON(dataGetVersionOKBodyAO1)
	if errGetVersionOKBodyAO1 != nil {
		return nil, errGetVersionOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetVersionOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get version o k body
func (o *GetVersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetVersionOKBodyAllOf0

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get version o k body based on the context it is used
func (o *GetVersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetVersionOKBodyAllOf0
	if err := o.GetVersionOKBodyAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetVersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVersionOKBody) UnmarshalBinary(b []byte) error {
	var res GetVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetVersionOKBodyAllOf0 get version o k body all of0
swagger:model GetVersionOKBodyAllOf0
*/
type GetVersionOKBodyAllOf0 string

// Validate validates this get version o k body all of0
func (o GetVersionOKBodyAllOf0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get version o k body all of0 based on context it is used
func (o GetVersionOKBodyAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
