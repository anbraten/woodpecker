// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new agents API client with basic auth credentials.
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

// New creates a new agents API client with a bearer token for authentication.
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
Client for agents API
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
	DeleteAgentsAgent(params *DeleteAgentsAgentParams, opts ...ClientOption) (*DeleteAgentsAgentOK, error)

	GetAgents(params *GetAgentsParams, opts ...ClientOption) (*GetAgentsOK, error)

	GetAgentsAgent(params *GetAgentsAgentParams, opts ...ClientOption) (*GetAgentsAgentOK, error)

	GetAgentsAgentTasks(params *GetAgentsAgentTasksParams, opts ...ClientOption) (*GetAgentsAgentTasksOK, error)

	PatchAgentsAgent(params *PatchAgentsAgentParams, opts ...ClientOption) (*PatchAgentsAgentOK, error)

	PostAgents(params *PostAgentsParams, opts ...ClientOption) (*PostAgentsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAgentsAgent deletes an agent
*/
func (a *Client) DeleteAgentsAgent(params *DeleteAgentsAgentParams, opts ...ClientOption) (*DeleteAgentsAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAgentsAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAgentsAgent",
		Method:             "DELETE",
		PathPattern:        "/agents/{agent}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAgentsAgentReader{formats: a.formats},
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
	success, ok := result.(*DeleteAgentsAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAgentsAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAgents lists agents
*/
func (a *Client) GetAgents(params *GetAgentsParams, opts ...ClientOption) (*GetAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgents",
		Method:             "GET",
		PathPattern:        "/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAgentsReader{formats: a.formats},
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
	success, ok := result.(*GetAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAgentsAgent gets an agent
*/
func (a *Client) GetAgentsAgent(params *GetAgentsAgentParams, opts ...ClientOption) (*GetAgentsAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentsAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgentsAgent",
		Method:             "GET",
		PathPattern:        "/agents/{agent}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAgentsAgentReader{formats: a.formats},
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
	success, ok := result.(*GetAgentsAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAgentsAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAgentsAgentTasks lists agent tasks
*/
func (a *Client) GetAgentsAgentTasks(params *GetAgentsAgentTasksParams, opts ...ClientOption) (*GetAgentsAgentTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentsAgentTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgentsAgentTasks",
		Method:             "GET",
		PathPattern:        "/agents/{agent}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAgentsAgentTasksReader{formats: a.formats},
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
	success, ok := result.(*GetAgentsAgentTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAgentsAgentTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAgentsAgent updates an agent
*/
func (a *Client) PatchAgentsAgent(params *PatchAgentsAgentParams, opts ...ClientOption) (*PatchAgentsAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAgentsAgentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAgentsAgent",
		Method:             "PATCH",
		PathPattern:        "/agents/{agent}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAgentsAgentReader{formats: a.formats},
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
	success, ok := result.(*PatchAgentsAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAgentsAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAgents creates a new agent

Creates a new agent with a random token
*/
func (a *Client) PostAgents(params *PostAgentsParams, opts ...ClientOption) (*PostAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAgents",
		Method:             "POST",
		PathPattern:        "/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAgentsReader{formats: a.formats},
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
	success, ok := result.(*PostAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAgents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
