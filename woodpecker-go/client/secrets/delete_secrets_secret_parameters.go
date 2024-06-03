// Code generated by go-swagger; DO NOT EDIT.

package secrets

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
)

// NewDeleteSecretsSecretParams creates a new DeleteSecretsSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSecretsSecretParams() *DeleteSecretsSecretParams {
	return &DeleteSecretsSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecretsSecretParamsWithTimeout creates a new DeleteSecretsSecretParams object
// with the ability to set a timeout on a request.
func NewDeleteSecretsSecretParamsWithTimeout(timeout time.Duration) *DeleteSecretsSecretParams {
	return &DeleteSecretsSecretParams{
		timeout: timeout,
	}
}

// NewDeleteSecretsSecretParamsWithContext creates a new DeleteSecretsSecretParams object
// with the ability to set a context for a request.
func NewDeleteSecretsSecretParamsWithContext(ctx context.Context) *DeleteSecretsSecretParams {
	return &DeleteSecretsSecretParams{
		Context: ctx,
	}
}

// NewDeleteSecretsSecretParamsWithHTTPClient creates a new DeleteSecretsSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSecretsSecretParamsWithHTTPClient(client *http.Client) *DeleteSecretsSecretParams {
	return &DeleteSecretsSecretParams{
		HTTPClient: client,
	}
}

/*
DeleteSecretsSecretParams contains all the parameters to send to the API endpoint

	for the delete secrets secret operation.

	Typically these are written to a http.Request.
*/
type DeleteSecretsSecretParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* Secret.

	   the secret's name
	*/
	Secret string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete secrets secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecretsSecretParams) WithDefaults() *DeleteSecretsSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete secrets secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSecretsSecretParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := DeleteSecretsSecretParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete secrets secret params
func (o *DeleteSecretsSecretParams) WithTimeout(timeout time.Duration) *DeleteSecretsSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete secrets secret params
func (o *DeleteSecretsSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete secrets secret params
func (o *DeleteSecretsSecretParams) WithContext(ctx context.Context) *DeleteSecretsSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete secrets secret params
func (o *DeleteSecretsSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete secrets secret params
func (o *DeleteSecretsSecretParams) WithHTTPClient(client *http.Client) *DeleteSecretsSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete secrets secret params
func (o *DeleteSecretsSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete secrets secret params
func (o *DeleteSecretsSecretParams) WithAuthorization(authorization string) *DeleteSecretsSecretParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete secrets secret params
func (o *DeleteSecretsSecretParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithSecret adds the secret to the delete secrets secret params
func (o *DeleteSecretsSecretParams) WithSecret(secret string) *DeleteSecretsSecretParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the delete secrets secret params
func (o *DeleteSecretsSecretParams) SetSecret(secret string) {
	o.Secret = secret
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecretsSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param secret
	if err := r.SetPathParam("secret", o.Secret); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
