// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewGetOrgLookupOrgFullNameParams creates a new GetOrgLookupOrgFullNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgLookupOrgFullNameParams() *GetOrgLookupOrgFullNameParams {
	return &GetOrgLookupOrgFullNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgLookupOrgFullNameParamsWithTimeout creates a new GetOrgLookupOrgFullNameParams object
// with the ability to set a timeout on a request.
func NewGetOrgLookupOrgFullNameParamsWithTimeout(timeout time.Duration) *GetOrgLookupOrgFullNameParams {
	return &GetOrgLookupOrgFullNameParams{
		timeout: timeout,
	}
}

// NewGetOrgLookupOrgFullNameParamsWithContext creates a new GetOrgLookupOrgFullNameParams object
// with the ability to set a context for a request.
func NewGetOrgLookupOrgFullNameParamsWithContext(ctx context.Context) *GetOrgLookupOrgFullNameParams {
	return &GetOrgLookupOrgFullNameParams{
		Context: ctx,
	}
}

// NewGetOrgLookupOrgFullNameParamsWithHTTPClient creates a new GetOrgLookupOrgFullNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgLookupOrgFullNameParamsWithHTTPClient(client *http.Client) *GetOrgLookupOrgFullNameParams {
	return &GetOrgLookupOrgFullNameParams{
		HTTPClient: client,
	}
}

/*
GetOrgLookupOrgFullNameParams contains all the parameters to send to the API endpoint

	for the get org lookup org full name operation.

	Typically these are written to a http.Request.
*/
type GetOrgLookupOrgFullNameParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* OrgFullName.

	   the organizations full name / slug
	*/
	OrgFullName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get org lookup org full name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgLookupOrgFullNameParams) WithDefaults() *GetOrgLookupOrgFullNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org lookup org full name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgLookupOrgFullNameParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := GetOrgLookupOrgFullNameParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) WithTimeout(timeout time.Duration) *GetOrgLookupOrgFullNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) WithContext(ctx context.Context) *GetOrgLookupOrgFullNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) WithHTTPClient(client *http.Client) *GetOrgLookupOrgFullNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) WithAuthorization(authorization string) *GetOrgLookupOrgFullNameParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithOrgFullName adds the orgFullName to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) WithOrgFullName(orgFullName string) *GetOrgLookupOrgFullNameParams {
	o.SetOrgFullName(orgFullName)
	return o
}

// SetOrgFullName adds the orgFullName to the get org lookup org full name params
func (o *GetOrgLookupOrgFullNameParams) SetOrgFullName(orgFullName string) {
	o.OrgFullName = orgFullName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgLookupOrgFullNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param org_full_name
	if err := r.SetPathParam("org_full_name", o.OrgFullName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
