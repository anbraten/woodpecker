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

// GetReposRepoIDPullRequestsReader is a Reader for the GetReposRepoIDPullRequests structure.
type GetReposRepoIDPullRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReposRepoIDPullRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReposRepoIDPullRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /repos/{repo_id}/pull_requests] GetReposRepoIDPullRequests", response, response.Code())
	}
}

// NewGetReposRepoIDPullRequestsOK creates a GetReposRepoIDPullRequestsOK with default headers values
func NewGetReposRepoIDPullRequestsOK() *GetReposRepoIDPullRequestsOK {
	return &GetReposRepoIDPullRequestsOK{}
}

/*
GetReposRepoIDPullRequestsOK describes a response with status code 200, with default header values.

OK
*/
type GetReposRepoIDPullRequestsOK struct {
	Payload []*models.PullRequest
}

// IsSuccess returns true when this get repos repo Id pull requests o k response has a 2xx status code
func (o *GetReposRepoIDPullRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get repos repo Id pull requests o k response has a 3xx status code
func (o *GetReposRepoIDPullRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get repos repo Id pull requests o k response has a 4xx status code
func (o *GetReposRepoIDPullRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get repos repo Id pull requests o k response has a 5xx status code
func (o *GetReposRepoIDPullRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get repos repo Id pull requests o k response a status code equal to that given
func (o *GetReposRepoIDPullRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get repos repo Id pull requests o k response
func (o *GetReposRepoIDPullRequestsOK) Code() int {
	return 200
}

func (o *GetReposRepoIDPullRequestsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{repo_id}/pull_requests][%d] getReposRepoIdPullRequestsOK %s", 200, payload)
}

func (o *GetReposRepoIDPullRequestsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{repo_id}/pull_requests][%d] getReposRepoIdPullRequestsOK %s", 200, payload)
}

func (o *GetReposRepoIDPullRequestsOK) GetPayload() []*models.PullRequest {
	return o.Payload
}

func (o *GetReposRepoIDPullRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
