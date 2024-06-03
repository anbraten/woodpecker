// Code generated by go-swagger; DO NOT EDIT.

package process_profiling_and_debugging

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

// NewGetDebugPprofGoroutineParams creates a new GetDebugPprofGoroutineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDebugPprofGoroutineParams() *GetDebugPprofGoroutineParams {
	return &GetDebugPprofGoroutineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugPprofGoroutineParamsWithTimeout creates a new GetDebugPprofGoroutineParams object
// with the ability to set a timeout on a request.
func NewGetDebugPprofGoroutineParamsWithTimeout(timeout time.Duration) *GetDebugPprofGoroutineParams {
	return &GetDebugPprofGoroutineParams{
		timeout: timeout,
	}
}

// NewGetDebugPprofGoroutineParamsWithContext creates a new GetDebugPprofGoroutineParams object
// with the ability to set a context for a request.
func NewGetDebugPprofGoroutineParamsWithContext(ctx context.Context) *GetDebugPprofGoroutineParams {
	return &GetDebugPprofGoroutineParams{
		Context: ctx,
	}
}

// NewGetDebugPprofGoroutineParamsWithHTTPClient creates a new GetDebugPprofGoroutineParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDebugPprofGoroutineParamsWithHTTPClient(client *http.Client) *GetDebugPprofGoroutineParams {
	return &GetDebugPprofGoroutineParams{
		HTTPClient: client,
	}
}

/*
GetDebugPprofGoroutineParams contains all the parameters to send to the API endpoint

	for the get debug pprof goroutine operation.

	Typically these are written to a http.Request.
*/
type GetDebugPprofGoroutineParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* Debug.

	   Use debug=2 as a query parameter to export in the same format as an un-recovered panic

	   Default: 1
	*/
	Debug *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get debug pprof goroutine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugPprofGoroutineParams) WithDefaults() *GetDebugPprofGoroutineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get debug pprof goroutine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugPprofGoroutineParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")

		debugDefault = int64(1)
	)

	val := GetDebugPprofGoroutineParams{
		Authorization: authorizationDefault,
		Debug:         &debugDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) WithTimeout(timeout time.Duration) *GetDebugPprofGoroutineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) WithContext(ctx context.Context) *GetDebugPprofGoroutineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) WithHTTPClient(client *http.Client) *GetDebugPprofGoroutineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) WithAuthorization(authorization string) *GetDebugPprofGoroutineParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithDebug adds the debug to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) WithDebug(debug *int64) *GetDebugPprofGoroutineParams {
	o.SetDebug(debug)
	return o
}

// SetDebug adds the debug to the get debug pprof goroutine params
func (o *GetDebugPprofGoroutineParams) SetDebug(debug *int64) {
	o.Debug = debug
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugPprofGoroutineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Debug != nil {

		// query param debug
		var qrDebug int64

		if o.Debug != nil {
			qrDebug = *o.Debug
		}
		qDebug := swag.FormatInt64(qrDebug)
		if qDebug != "" {

			if err := r.SetQueryParam("debug", qDebug); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
