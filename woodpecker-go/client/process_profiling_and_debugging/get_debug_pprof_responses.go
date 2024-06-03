// Code generated by go-swagger; DO NOT EDIT.

package process_profiling_and_debugging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDebugPprofReader is a Reader for the GetDebugPprof structure.
type GetDebugPprofReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugPprofReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDebugPprofOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /debug/pprof] GetDebugPprof", response, response.Code())
	}
}

// NewGetDebugPprofOK creates a GetDebugPprofOK with default headers values
func NewGetDebugPprofOK() *GetDebugPprofOK {
	return &GetDebugPprofOK{}
}

/*
GetDebugPprofOK describes a response with status code 200, with default header values.

OK
*/
type GetDebugPprofOK struct {
}

// IsSuccess returns true when this get debug pprof o k response has a 2xx status code
func (o *GetDebugPprofOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get debug pprof o k response has a 3xx status code
func (o *GetDebugPprofOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get debug pprof o k response has a 4xx status code
func (o *GetDebugPprofOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get debug pprof o k response has a 5xx status code
func (o *GetDebugPprofOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get debug pprof o k response a status code equal to that given
func (o *GetDebugPprofOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get debug pprof o k response
func (o *GetDebugPprofOK) Code() int {
	return 200
}

func (o *GetDebugPprofOK) Error() string {
	return fmt.Sprintf("[GET /debug/pprof][%d] getDebugPprofOK", 200)
}

func (o *GetDebugPprofOK) String() string {
	return fmt.Sprintf("[GET /debug/pprof][%d] getDebugPprofOK", 200)
}

func (o *GetDebugPprofOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
