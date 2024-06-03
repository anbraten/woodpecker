// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"go.woodpecker-ci.org/woodpecker/v2/woodpecker-go/models"
)

// PostUsersReader is a Reader for the PostUsers structure.
type PostUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /users] PostUsers", response, response.Code())
	}
}

// NewPostUsersOK creates a PostUsersOK with default headers values
func NewPostUsersOK() *PostUsersOK {
	return &PostUsersOK{}
}

/*
PostUsersOK describes a response with status code 200, with default header values.

OK
*/
type PostUsersOK struct {
	Payload *models.User
}

// IsSuccess returns true when this post users o k response has a 2xx status code
func (o *PostUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post users o k response has a 3xx status code
func (o *PostUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users o k response has a 4xx status code
func (o *PostUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users o k response has a 5xx status code
func (o *PostUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post users o k response a status code equal to that given
func (o *PostUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post users o k response
func (o *PostUsersOK) Code() int {
	return 200
}

func (o *PostUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersOK %s", 200, payload)
}

func (o *PostUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users][%d] postUsersOK %s", 200, payload)
}

func (o *PostUsersOK) GetPayload() *models.User {
	return o.Payload
}

func (o *PostUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
