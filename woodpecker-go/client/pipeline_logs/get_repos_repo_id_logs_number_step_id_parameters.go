// Code generated by go-swagger; DO NOT EDIT.

package pipeline_logs

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

// NewGetReposRepoIDLogsNumberStepIDParams creates a new GetReposRepoIDLogsNumberStepIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReposRepoIDLogsNumberStepIDParams() *GetReposRepoIDLogsNumberStepIDParams {
	return &GetReposRepoIDLogsNumberStepIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReposRepoIDLogsNumberStepIDParamsWithTimeout creates a new GetReposRepoIDLogsNumberStepIDParams object
// with the ability to set a timeout on a request.
func NewGetReposRepoIDLogsNumberStepIDParamsWithTimeout(timeout time.Duration) *GetReposRepoIDLogsNumberStepIDParams {
	return &GetReposRepoIDLogsNumberStepIDParams{
		timeout: timeout,
	}
}

// NewGetReposRepoIDLogsNumberStepIDParamsWithContext creates a new GetReposRepoIDLogsNumberStepIDParams object
// with the ability to set a context for a request.
func NewGetReposRepoIDLogsNumberStepIDParamsWithContext(ctx context.Context) *GetReposRepoIDLogsNumberStepIDParams {
	return &GetReposRepoIDLogsNumberStepIDParams{
		Context: ctx,
	}
}

// NewGetReposRepoIDLogsNumberStepIDParamsWithHTTPClient creates a new GetReposRepoIDLogsNumberStepIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReposRepoIDLogsNumberStepIDParamsWithHTTPClient(client *http.Client) *GetReposRepoIDLogsNumberStepIDParams {
	return &GetReposRepoIDLogsNumberStepIDParams{
		HTTPClient: client,
	}
}

/*
GetReposRepoIDLogsNumberStepIDParams contains all the parameters to send to the API endpoint

	for the get repos repo ID logs number step ID operation.

	Typically these are written to a http.Request.
*/
type GetReposRepoIDLogsNumberStepIDParams struct {

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

	/* StepID.

	   the step id
	*/
	StepID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repos repo ID logs number step ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReposRepoIDLogsNumberStepIDParams) WithDefaults() *GetReposRepoIDLogsNumberStepIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repos repo ID logs number step ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReposRepoIDLogsNumberStepIDParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := GetReposRepoIDLogsNumberStepIDParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithTimeout(timeout time.Duration) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithContext(ctx context.Context) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithHTTPClient(client *http.Client) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithAuthorization(authorization string) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithNumber adds the number to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithNumber(number int64) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetNumber(number int64) {
	o.Number = number
}

// WithRepoID adds the repoID to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithRepoID(repoID int64) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetRepoID(repoID int64) {
	o.RepoID = repoID
}

// WithStepID adds the stepID to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) WithStepID(stepID int64) *GetReposRepoIDLogsNumberStepIDParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get repos repo ID logs number step ID params
func (o *GetReposRepoIDLogsNumberStepIDParams) SetStepID(stepID int64) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReposRepoIDLogsNumberStepIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stepID
	if err := r.SetPathParam("stepID", swag.FormatInt64(o.StepID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
