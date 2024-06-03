// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// NewDeleteOrgsIDParams creates a new DeleteOrgsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgsIDParams() *DeleteOrgsIDParams {
	return &DeleteOrgsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgsIDParamsWithTimeout creates a new DeleteOrgsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgsIDParamsWithTimeout(timeout time.Duration) *DeleteOrgsIDParams {
	return &DeleteOrgsIDParams{
		timeout: timeout,
	}
}

// NewDeleteOrgsIDParamsWithContext creates a new DeleteOrgsIDParams object
// with the ability to set a context for a request.
func NewDeleteOrgsIDParamsWithContext(ctx context.Context) *DeleteOrgsIDParams {
	return &DeleteOrgsIDParams{
		Context: ctx,
	}
}

// NewDeleteOrgsIDParamsWithHTTPClient creates a new DeleteOrgsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgsIDParamsWithHTTPClient(client *http.Client) *DeleteOrgsIDParams {
	return &DeleteOrgsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgsIDParams contains all the parameters to send to the API endpoint

	for the delete orgs ID operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgsIDParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* ID.

	   the org's id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete orgs ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgsIDParams) WithDefaults() *DeleteOrgsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete orgs ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgsIDParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := DeleteOrgsIDParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete orgs ID params
func (o *DeleteOrgsIDParams) WithTimeout(timeout time.Duration) *DeleteOrgsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orgs ID params
func (o *DeleteOrgsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orgs ID params
func (o *DeleteOrgsIDParams) WithContext(ctx context.Context) *DeleteOrgsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orgs ID params
func (o *DeleteOrgsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orgs ID params
func (o *DeleteOrgsIDParams) WithHTTPClient(client *http.Client) *DeleteOrgsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orgs ID params
func (o *DeleteOrgsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete orgs ID params
func (o *DeleteOrgsIDParams) WithAuthorization(authorization string) *DeleteOrgsIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete orgs ID params
func (o *DeleteOrgsIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithID adds the id to the delete orgs ID params
func (o *DeleteOrgsIDParams) WithID(id string) *DeleteOrgsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete orgs ID params
func (o *DeleteOrgsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
