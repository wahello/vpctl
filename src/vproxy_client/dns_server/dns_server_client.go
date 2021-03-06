// Code generated by go-swagger; DO NOT EDIT.

package dns_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dns server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dns server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddDNSServer adds dns server
*/
func (a *Client) AddDNSServer(params *AddDNSServerParams) (*AddDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addDnsServer",
		Method:             "POST",
		PathPattern:        "/dns-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeDNSServer gets detailed info of one dns server
*/
func (a *Client) DescribeDNSServer(params *DescribeDNSServerParams) (*DescribeDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "describeDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server/{dns}/detail",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDNSServer gets dns server
*/
func (a *Client) GetDNSServer(params *GetDNSServerParams) (*GetDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListDNSServer retrieves dns server list
*/
func (a *Client) ListDNSServer(params *ListDNSServerParams) (*ListDNSServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDnsServer",
		Method:             "GET",
		PathPattern:        "/dns-server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDNSServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveDNSServer removes dns server
*/
func (a *Client) RemoveDNSServer(params *RemoveDNSServerParams) (*RemoveDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeDnsServer",
		Method:             "DELETE",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDNSServer updates dns server
*/
func (a *Client) UpdateDNSServer(params *UpdateDNSServerParams) (*UpdateDNSServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDNSServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDnsServer",
		Method:             "PUT",
		PathPattern:        "/dns-server/{dns}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDNSServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDNSServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDnsServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
