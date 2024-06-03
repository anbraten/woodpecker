// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPostReposRepoIDPipelinesNumberApproveParams creates a new PostReposRepoIDPipelinesNumberApproveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostReposRepoIDPipelinesNumberApproveParams() *PostReposRepoIDPipelinesNumberApproveParams {
	return &PostReposRepoIDPipelinesNumberApproveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostReposRepoIDPipelinesNumberApproveParamsWithTimeout creates a new PostReposRepoIDPipelinesNumberApproveParams object
// with the ability to set a timeout on a request.
func NewPostReposRepoIDPipelinesNumberApproveParamsWithTimeout(timeout time.Duration) *PostReposRepoIDPipelinesNumberApproveParams {
	return &PostReposRepoIDPipelinesNumberApproveParams{
		timeout: timeout,
	}
}

// NewPostReposRepoIDPipelinesNumberApproveParamsWithContext creates a new PostReposRepoIDPipelinesNumberApproveParams object
// with the ability to set a context for a request.
func NewPostReposRepoIDPipelinesNumberApproveParamsWithContext(ctx context.Context) *PostReposRepoIDPipelinesNumberApproveParams {
	return &PostReposRepoIDPipelinesNumberApproveParams{
		Context: ctx,
	}
}

// NewPostReposRepoIDPipelinesNumberApproveParamsWithHTTPClient creates a new PostReposRepoIDPipelinesNumberApproveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostReposRepoIDPipelinesNumberApproveParamsWithHTTPClient(client *http.Client) *PostReposRepoIDPipelinesNumberApproveParams {
	return &PostReposRepoIDPipelinesNumberApproveParams{
		HTTPClient: client,
	}
}

/*
PostReposRepoIDPipelinesNumberApproveParams contains all the parameters to send to the API endpoint

	for the post repos repo ID pipelines number approve operation.

	Typically these are written to a http.Request.
*/
type PostReposRepoIDPipelinesNumberApproveParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* Number.

	   the number of the pipeline
	*/
	Number int64

	/* RepoID.

	   the repository id
	*/
	RepoID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repos repo ID pipelines number approve params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithDefaults() *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repos repo ID pipelines number approve params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := PostReposRepoIDPipelinesNumberApproveParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithTimeout(timeout time.Duration) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithContext(ctx context.Context) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithHTTPClient(client *http.Client) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithAuthorization(authorization string) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNumber adds the number to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithNumber(number int64) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetNumber(number int64) {
	o.Number = number
}

// WithRepoID adds the repoID to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) WithRepoID(repoID int64) *PostReposRepoIDPipelinesNumberApproveParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the post repos repo ID pipelines number approve params
func (o *PostReposRepoIDPipelinesNumberApproveParams) SetRepoID(repoID int64) {
	o.RepoID = repoID
}

// WriteToRequest writes these params to a swagger request
func (o *PostReposRepoIDPipelinesNumberApproveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", swag.FormatInt64(o.Number)); err != nil {
		return err
	}

	// path param repo_id
	if err := r.SetPathParam("repo_id", swag.FormatInt64(o.RepoID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
