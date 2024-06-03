// Code generated by go-swagger; DO NOT EDIT.

package repository_cron_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new repository cron jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new repository cron jobs API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new repository cron jobs API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for repository cron jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptTextPlain sets the Accept header to "text/plain".
func WithAcceptTextPlain(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/plain"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteReposRepoIDCronCron(params *DeleteReposRepoIDCronCronParams, opts ...ClientOption) (*DeleteReposRepoIDCronCronNoContent, error)

	GetReposRepoIDCron(params *GetReposRepoIDCronParams, opts ...ClientOption) (*GetReposRepoIDCronOK, error)

	GetReposRepoIDCronCron(params *GetReposRepoIDCronCronParams, opts ...ClientOption) (*GetReposRepoIDCronCronOK, error)

	PatchReposRepoIDCronCron(params *PatchReposRepoIDCronCronParams, opts ...ClientOption) (*PatchReposRepoIDCronCronOK, error)

	PostReposRepoIDCron(params *PostReposRepoIDCronParams, opts ...ClientOption) (*PostReposRepoIDCronOK, error)

	PostReposRepoIDCronCron(params *PostReposRepoIDCronCronParams, opts ...ClientOption) (*PostReposRepoIDCronCronOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteReposRepoIDCronCron deletes a cron job
*/
func (a *Client) DeleteReposRepoIDCronCron(params *DeleteReposRepoIDCronCronParams, opts ...ClientOption) (*DeleteReposRepoIDCronCronNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReposRepoIDCronCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteReposRepoIDCronCron",
		Method:             "DELETE",
		PathPattern:        "/repos/{repo_id}/cron/{cron}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReposRepoIDCronCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteReposRepoIDCronCronNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteReposRepoIDCronCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReposRepoIDCron lists cron jobs
*/
func (a *Client) GetReposRepoIDCron(params *GetReposRepoIDCronParams, opts ...ClientOption) (*GetReposRepoIDCronOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReposRepoIDCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetReposRepoIDCron",
		Method:             "GET",
		PathPattern:        "/repos/{repo_id}/cron",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReposRepoIDCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReposRepoIDCronOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetReposRepoIDCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetReposRepoIDCronCron gets a cron job
*/
func (a *Client) GetReposRepoIDCronCron(params *GetReposRepoIDCronCronParams, opts ...ClientOption) (*GetReposRepoIDCronCronOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReposRepoIDCronCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetReposRepoIDCronCron",
		Method:             "GET",
		PathPattern:        "/repos/{repo_id}/cron/{cron}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReposRepoIDCronCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReposRepoIDCronCronOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetReposRepoIDCronCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchReposRepoIDCronCron updates a cron job
*/
func (a *Client) PatchReposRepoIDCronCron(params *PatchReposRepoIDCronCronParams, opts ...ClientOption) (*PatchReposRepoIDCronCronOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchReposRepoIDCronCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchReposRepoIDCronCron",
		Method:             "PATCH",
		PathPattern:        "/repos/{repo_id}/cron/{cron}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchReposRepoIDCronCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchReposRepoIDCronCronOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchReposRepoIDCronCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostReposRepoIDCron creates a cron job
*/
func (a *Client) PostReposRepoIDCron(params *PostReposRepoIDCronParams, opts ...ClientOption) (*PostReposRepoIDCronOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReposRepoIDCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostReposRepoIDCron",
		Method:             "POST",
		PathPattern:        "/repos/{repo_id}/cron",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostReposRepoIDCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostReposRepoIDCronOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostReposRepoIDCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostReposRepoIDCronCron starts a cron job now
*/
func (a *Client) PostReposRepoIDCronCron(params *PostReposRepoIDCronCronParams, opts ...ClientOption) (*PostReposRepoIDCronCronOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostReposRepoIDCronCronParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostReposRepoIDCronCron",
		Method:             "POST",
		PathPattern:        "/repos/{repo_id}/cron/{cron}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostReposRepoIDCronCronReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostReposRepoIDCronCronOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostReposRepoIDCronCron: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
