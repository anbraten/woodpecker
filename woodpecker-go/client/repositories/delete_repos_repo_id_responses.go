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

// DeleteReposRepoIDReader is a Reader for the DeleteReposRepoID structure.
type DeleteReposRepoIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReposRepoIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteReposRepoIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{repo_id}] DeleteReposRepoID", response, response.Code())
	}
}

// NewDeleteReposRepoIDOK creates a DeleteReposRepoIDOK with default headers values
func NewDeleteReposRepoIDOK() *DeleteReposRepoIDOK {
	return &DeleteReposRepoIDOK{}
}

/*
DeleteReposRepoIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteReposRepoIDOK struct {
	Payload *models.Repo
}

// IsSuccess returns true when this delete repos repo Id o k response has a 2xx status code
func (o *DeleteReposRepoIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete repos repo Id o k response has a 3xx status code
func (o *DeleteReposRepoIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repos repo Id o k response has a 4xx status code
func (o *DeleteReposRepoIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete repos repo Id o k response has a 5xx status code
func (o *DeleteReposRepoIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repos repo Id o k response a status code equal to that given
func (o *DeleteReposRepoIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete repos repo Id o k response
func (o *DeleteReposRepoIDOK) Code() int {
	return 200
}

func (o *DeleteReposRepoIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{repo_id}][%d] deleteReposRepoIdOK %s", 200, payload)
}

func (o *DeleteReposRepoIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{repo_id}][%d] deleteReposRepoIdOK %s", 200, payload)
}

func (o *DeleteReposRepoIDOK) GetPayload() *models.Repo {
	return o.Payload
}

func (o *DeleteReposRepoIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
