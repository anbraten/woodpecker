// Code generated by go-swagger; DO NOT EDIT.

package repository_cron_jobs

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

// NewDeleteReposRepoIDCronCronParams creates a new DeleteReposRepoIDCronCronParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteReposRepoIDCronCronParams() *DeleteReposRepoIDCronCronParams {
	return &DeleteReposRepoIDCronCronParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReposRepoIDCronCronParamsWithTimeout creates a new DeleteReposRepoIDCronCronParams object
// with the ability to set a timeout on a request.
func NewDeleteReposRepoIDCronCronParamsWithTimeout(timeout time.Duration) *DeleteReposRepoIDCronCronParams {
	return &DeleteReposRepoIDCronCronParams{
		timeout: timeout,
	}
}

// NewDeleteReposRepoIDCronCronParamsWithContext creates a new DeleteReposRepoIDCronCronParams object
// with the ability to set a context for a request.
func NewDeleteReposRepoIDCronCronParamsWithContext(ctx context.Context) *DeleteReposRepoIDCronCronParams {
	return &DeleteReposRepoIDCronCronParams{
		Context: ctx,
	}
}

// NewDeleteReposRepoIDCronCronParamsWithHTTPClient creates a new DeleteReposRepoIDCronCronParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteReposRepoIDCronCronParamsWithHTTPClient(client *http.Client) *DeleteReposRepoIDCronCronParams {
	return &DeleteReposRepoIDCronCronParams{
		HTTPClient: client,
	}
}

/*
DeleteReposRepoIDCronCronParams contains all the parameters to send to the API endpoint

	for the delete repos repo ID cron cron operation.

	Typically these are written to a http.Request.
*/
type DeleteReposRepoIDCronCronParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* Cron.

	   the cron job id
	*/
	Cron string

	/* RepoID.

	   the repository id
	*/
	RepoID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete repos repo ID cron cron params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReposRepoIDCronCronParams) WithDefaults() *DeleteReposRepoIDCronCronParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete repos repo ID cron cron params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReposRepoIDCronCronParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := DeleteReposRepoIDCronCronParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithTimeout(timeout time.Duration) *DeleteReposRepoIDCronCronParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithContext(ctx context.Context) *DeleteReposRepoIDCronCronParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithHTTPClient(client *http.Client) *DeleteReposRepoIDCronCronParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithAuthorization(authorization string) *DeleteReposRepoIDCronCronParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCron adds the cron to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithCron(cron string) *DeleteReposRepoIDCronCronParams {
	o.SetCron(cron)
	return o
}

// SetCron adds the cron to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetCron(cron string) {
	o.Cron = cron
}

// WithRepoID adds the repoID to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) WithRepoID(repoID int64) *DeleteReposRepoIDCronCronParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the delete repos repo ID cron cron params
func (o *DeleteReposRepoIDCronCronParams) SetRepoID(repoID int64) {
	o.RepoID = repoID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReposRepoIDCronCronParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param cron
	if err := r.SetPathParam("cron", o.Cron); err != nil {
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
