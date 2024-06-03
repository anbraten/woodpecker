// Code generated by go-swagger; DO NOT EDIT.

package pipeline_queues

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

// NewGetQueueInfoParams creates a new GetQueueInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueueInfoParams() *GetQueueInfoParams {
	return &GetQueueInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueueInfoParamsWithTimeout creates a new GetQueueInfoParams object
// with the ability to set a timeout on a request.
func NewGetQueueInfoParamsWithTimeout(timeout time.Duration) *GetQueueInfoParams {
	return &GetQueueInfoParams{
		timeout: timeout,
	}
}

// NewGetQueueInfoParamsWithContext creates a new GetQueueInfoParams object
// with the ability to set a context for a request.
func NewGetQueueInfoParamsWithContext(ctx context.Context) *GetQueueInfoParams {
	return &GetQueueInfoParams{
		Context: ctx,
	}
}

// NewGetQueueInfoParamsWithHTTPClient creates a new GetQueueInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueueInfoParamsWithHTTPClient(client *http.Client) *GetQueueInfoParams {
	return &GetQueueInfoParams{
		HTTPClient: client,
	}
}

/*
GetQueueInfoParams contains all the parameters to send to the API endpoint

	for the get queue info operation.

	Typically these are written to a http.Request.
*/
type GetQueueInfoParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get queue info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueueInfoParams) WithDefaults() *GetQueueInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get queue info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueueInfoParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := GetQueueInfoParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get queue info params
func (o *GetQueueInfoParams) WithTimeout(timeout time.Duration) *GetQueueInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queue info params
func (o *GetQueueInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queue info params
func (o *GetQueueInfoParams) WithContext(ctx context.Context) *GetQueueInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queue info params
func (o *GetQueueInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queue info params
func (o *GetQueueInfoParams) WithHTTPClient(client *http.Client) *GetQueueInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queue info params
func (o *GetQueueInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get queue info params
func (o *GetQueueInfoParams) WithAuthorization(authorization string) *GetQueueInfoParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get queue info params
func (o *GetQueueInfoParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueueInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
