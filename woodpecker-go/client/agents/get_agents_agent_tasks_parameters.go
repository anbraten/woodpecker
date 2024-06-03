// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewGetAgentsAgentTasksParams creates a new GetAgentsAgentTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentsAgentTasksParams() *GetAgentsAgentTasksParams {
	return &GetAgentsAgentTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentsAgentTasksParamsWithTimeout creates a new GetAgentsAgentTasksParams object
// with the ability to set a timeout on a request.
func NewGetAgentsAgentTasksParamsWithTimeout(timeout time.Duration) *GetAgentsAgentTasksParams {
	return &GetAgentsAgentTasksParams{
		timeout: timeout,
	}
}

// NewGetAgentsAgentTasksParamsWithContext creates a new GetAgentsAgentTasksParams object
// with the ability to set a context for a request.
func NewGetAgentsAgentTasksParamsWithContext(ctx context.Context) *GetAgentsAgentTasksParams {
	return &GetAgentsAgentTasksParams{
		Context: ctx,
	}
}

// NewGetAgentsAgentTasksParamsWithHTTPClient creates a new GetAgentsAgentTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentsAgentTasksParamsWithHTTPClient(client *http.Client) *GetAgentsAgentTasksParams {
	return &GetAgentsAgentTasksParams{
		HTTPClient: client,
	}
}

/*
GetAgentsAgentTasksParams contains all the parameters to send to the API endpoint

	for the get agents agent tasks operation.

	Typically these are written to a http.Request.
*/
type GetAgentsAgentTasksParams struct {

	/* Authorization.

	   Insert your personal access token

	   Default: "Bearer \u003cpersonal access token\u003e"
	*/
	Authorization string

	/* Agent.

	   the agent's id
	*/
	Agent int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agents agent tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentsAgentTasksParams) WithDefaults() *GetAgentsAgentTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agents agent tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentsAgentTasksParams) SetDefaults() {
	var (
		authorizationDefault = string("Bearer <personal access token>")
	)

	val := GetAgentsAgentTasksParams{
		Authorization: authorizationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) WithTimeout(timeout time.Duration) *GetAgentsAgentTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) WithContext(ctx context.Context) *GetAgentsAgentTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) WithHTTPClient(client *http.Client) *GetAgentsAgentTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) WithAuthorization(authorization string) *GetAgentsAgentTasksParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAgent adds the agent to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) WithAgent(agent int64) *GetAgentsAgentTasksParams {
	o.SetAgent(agent)
	return o
}

// SetAgent adds the agent to the get agents agent tasks params
func (o *GetAgentsAgentTasksParams) SetAgent(agent int64) {
	o.Agent = agent
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentsAgentTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param agent
	if err := r.SetPathParam("agent", swag.FormatInt64(o.Agent)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
