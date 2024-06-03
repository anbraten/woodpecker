// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostUserTokenReader is a Reader for the PostUserToken structure.
type PostUserTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /user/token] PostUserToken", response, response.Code())
	}
}

// NewPostUserTokenOK creates a PostUserTokenOK with default headers values
func NewPostUserTokenOK() *PostUserTokenOK {
	return &PostUserTokenOK{}
}

/*
PostUserTokenOK describes a response with status code 200, with default header values.

OK
*/
type PostUserTokenOK struct {
}

// IsSuccess returns true when this post user token o k response has a 2xx status code
func (o *PostUserTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post user token o k response has a 3xx status code
func (o *PostUserTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user token o k response has a 4xx status code
func (o *PostUserTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user token o k response has a 5xx status code
func (o *PostUserTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post user token o k response a status code equal to that given
func (o *PostUserTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post user token o k response
func (o *PostUserTokenOK) Code() int {
	return 200
}

func (o *PostUserTokenOK) Error() string {
	return fmt.Sprintf("[POST /user/token][%d] postUserTokenOK", 200)
}

func (o *PostUserTokenOK) String() string {
	return fmt.Sprintf("[POST /user/token][%d] postUserTokenOK", 200)
}

func (o *PostUserTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
