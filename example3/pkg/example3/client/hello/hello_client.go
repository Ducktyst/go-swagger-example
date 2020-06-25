// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hello API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hello API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	HelloWorld(params *HelloWorldParams) (*HelloWorldOK, error)

	HelloWorldFull(params *HelloWorldFullParams) (*HelloWorldFullOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  HelloWorld examples route

  Some description
*/
func (a *Client) HelloWorld(params *HelloWorldParams) (*HelloWorldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloWorldParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HelloWorld",
		Method:             "GET",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HelloWorldReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloWorldOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HelloWorld: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HelloWorldFull examples route

  Some description
*/
func (a *Client) HelloWorldFull(params *HelloWorldFullParams) (*HelloWorldFullOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelloWorldFullParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HelloWorldFull",
		Method:             "GET",
		PathPattern:        "/hello-world",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HelloWorldFullReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelloWorldFullOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for HelloWorldFull: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
