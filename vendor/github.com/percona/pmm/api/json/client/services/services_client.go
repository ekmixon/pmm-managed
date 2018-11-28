// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddMySQLService adds my SQL service adds my SQL service
*/
func (a *Client) AddMySQLService(params *AddMySQLServiceParams) (*AddMySQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLService",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Services/AddMySQLService",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMySQLServiceOK), nil

}

/*
ChangeMySQLService changes my SQL service changes my SQL service
*/
func (a *Client) ChangeMySQLService(params *ChangeMySQLServiceParams) (*ChangeMySQLServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMySQLServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeMySQLService",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Services/ChangeMySQLService",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeMySQLServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeMySQLServiceOK), nil

}

/*
GetService gets service returns a single service by ID
*/
func (a *Client) GetService(params *GetServiceParams) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetService",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Services/GetService",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceOK), nil

}

/*
ListServices lists services returns a list of all services
*/
func (a *Client) ListServices(params *ListServicesParams) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListServices",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Services/ListServices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListServicesOK), nil

}

/*
RemoveService removes service removes service without any agents
*/
func (a *Client) RemoveService(params *RemoveServiceParams) (*RemoveServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveService",
		Method:             "POST",
		PathPattern:        "/v0/inventory/Services/RemoveService",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveServiceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
