// Code generated by go-swagger; DO NOT EDIT.

package repository_secrets

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

	"go.woodpecker-ci.org/woodpecker/v2/woodpecker-go/models"
)

// NewPostReposRepoIDSecretsParams creates a new PostReposRepoIDSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostReposRepoIDSecretsParams() *PostReposRepoIDSecretsParams {
	return &PostReposRepoIDSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostReposRepoIDSecretsParamsWithTimeout creates a new PostReposRepoIDSecretsParams object
// with the ability to set a timeout on a request.
func NewPostReposRepoIDSecretsParamsWithTimeout(timeout time.Duration) *PostReposRepoIDSecretsParams {
	return &PostReposRepoIDSecretsParams{
		timeout: timeout,
	}
}

// NewPostReposRepoIDSecretsParamsWithContext creates a new PostReposRepoIDSecretsParams object
// with the ability to set a context for a request.
func NewPostReposRepoIDSecretsParamsWithContext(ctx context.Context) *PostReposRepoIDSecretsParams {
	return &PostReposRepoIDSecretsParams{
		Context: ctx,
	}
}

// NewPostReposRepoIDSecretsParamsWithHTTPClient creates a new PostReposRepoIDSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostReposRepoIDSecretsParamsWithHTTPClient(client *http.Client) *PostReposRepoIDSecretsParams {
	return &PostReposRepoIDSecretsParams{
		HTTPClient: client,
	}
}

/*
PostReposRepoIDSecretsParams contains all the parameters to send to the API endpoint

	for the post repos repo ID secrets operation.

	Typically these are written to a http.Request.
*/
type PostReposRepoIDSecretsParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* RepoID.

	   the repository id
	*/
	RepoID int64

	/* Secret.

	   the new secret
	*/
	Secret *models.Secret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post repos repo ID secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReposRepoIDSecretsParams) WithDefaults() *PostReposRepoIDSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post repos repo ID secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostReposRepoIDSecretsParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := PostReposRepoIDSecretsParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithTimeout(timeout time.Duration) *PostReposRepoIDSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithContext(ctx context.Context) *PostReposRepoIDSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithHTTPClient(client *http.Client) *PostReposRepoIDSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithAuthorization(authorization string) *PostReposRepoIDSecretsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRepoID adds the repoID to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithRepoID(repoID int64) *PostReposRepoIDSecretsParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetRepoID(repoID int64) {
	o.RepoID = repoID
}

// WithSecret adds the secret to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) WithSecret(secret *models.Secret) *PostReposRepoIDSecretsParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the post repos repo ID secrets params
func (o *PostReposRepoIDSecretsParams) SetSecret(secret *models.Secret) {
	o.Secret = secret
}

// WriteToRequest writes these params to a swagger request
func (o *PostReposRepoIDSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param repo_id
	if err := r.SetPathParam("repo_id", swag.FormatInt64(o.RepoID)); err != nil {
		return err
	}
	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
