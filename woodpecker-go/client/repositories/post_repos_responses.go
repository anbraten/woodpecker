// Code generated by go-swagger; DO NOT EDIT.

package repositories

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

// PostReposReader is a Reader for the PostRepos structure.
type PostReposReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReposReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /repos] PostRepos", response, response.Code())
	}
}

// NewPostReposOK creates a PostReposOK with default headers values
func NewPostReposOK() *PostReposOK {
	return &PostReposOK{}
}

/*
PostReposOK describes a response with status code 200, with default header values.

OK
*/
type PostReposOK struct {
	Payload *models.Repo
}

// IsSuccess returns true when this post repos o k response has a 2xx status code
func (o *PostReposOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repos o k response has a 3xx status code
func (o *PostReposOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repos o k response has a 4xx status code
func (o *PostReposOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repos o k response has a 5xx status code
func (o *PostReposOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post repos o k response a status code equal to that given
func (o *PostReposOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post repos o k response
func (o *PostReposOK) Code() int {
	return 200
}

func (o *PostReposOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos][%d] postReposOK %s", 200, payload)
}

func (o *PostReposOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos][%d] postReposOK %s", 200, payload)
}

func (o *PostReposOK) GetPayload() *models.Repo {
	return o.Payload
}

func (o *PostReposOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
