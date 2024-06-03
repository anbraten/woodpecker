// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostReposRepoIDPipelinesNumberCancelReader is a Reader for the PostReposRepoIDPipelinesNumberCancel structure.
type PostReposRepoIDPipelinesNumberCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostReposRepoIDPipelinesNumberCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostReposRepoIDPipelinesNumberCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /repos/{repo_id}/pipelines/{number}/cancel] PostReposRepoIDPipelinesNumberCancel", response, response.Code())
	}
}

// NewPostReposRepoIDPipelinesNumberCancelOK creates a PostReposRepoIDPipelinesNumberCancelOK with default headers values
func NewPostReposRepoIDPipelinesNumberCancelOK() *PostReposRepoIDPipelinesNumberCancelOK {
	return &PostReposRepoIDPipelinesNumberCancelOK{}
}

/*
PostReposRepoIDPipelinesNumberCancelOK describes a response with status code 200, with default header values.

OK
*/
type PostReposRepoIDPipelinesNumberCancelOK struct {
}

// IsSuccess returns true when this post repos repo Id pipelines number cancel o k response has a 2xx status code
func (o *PostReposRepoIDPipelinesNumberCancelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repos repo Id pipelines number cancel o k response has a 3xx status code
func (o *PostReposRepoIDPipelinesNumberCancelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repos repo Id pipelines number cancel o k response has a 4xx status code
func (o *PostReposRepoIDPipelinesNumberCancelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repos repo Id pipelines number cancel o k response has a 5xx status code
func (o *PostReposRepoIDPipelinesNumberCancelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post repos repo Id pipelines number cancel o k response a status code equal to that given
func (o *PostReposRepoIDPipelinesNumberCancelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post repos repo Id pipelines number cancel o k response
func (o *PostReposRepoIDPipelinesNumberCancelOK) Code() int {
	return 200
}

func (o *PostReposRepoIDPipelinesNumberCancelOK) Error() string {
	return fmt.Sprintf("[POST /repos/{repo_id}/pipelines/{number}/cancel][%d] postReposRepoIdPipelinesNumberCancelOK", 200)
}

func (o *PostReposRepoIDPipelinesNumberCancelOK) String() string {
	return fmt.Sprintf("[POST /repos/{repo_id}/pipelines/{number}/cancel][%d] postReposRepoIdPipelinesNumberCancelOK", 200)
}

func (o *PostReposRepoIDPipelinesNumberCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
