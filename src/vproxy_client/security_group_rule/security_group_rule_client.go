// Code generated by go-swagger; DO NOT EDIT.

package security_group_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new security group rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security group rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddSecurityGroupRule adds security group rule into security group
*/
func (a *Client) AddSecurityGroupRule(params *AddSecurityGroupRuleParams) (*AddSecurityGroupRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSecurityGroupRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSecurityGroupRule",
		Method:             "POST",
		PathPattern:        "/security-group/{secg}/security-group-rule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSecurityGroupRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddSecurityGroupRuleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addSecurityGroupRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DescribeSecurityGroupRule gets detailed info of one security group rule in security group
*/
func (a *Client) DescribeSecurityGroupRule(params *DescribeSecurityGroupRuleParams) (*DescribeSecurityGroupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeSecurityGroupRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "describeSecurityGroupRule",
		Method:             "GET",
		PathPattern:        "/security-group/{secg}/security-group-rule/{secgr}/detail",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeSecurityGroupRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DescribeSecurityGroupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for describeSecurityGroupRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSecurityGroupRule gets security group rule in security group
*/
func (a *Client) GetSecurityGroupRule(params *GetSecurityGroupRuleParams) (*GetSecurityGroupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityGroupRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSecurityGroupRule",
		Method:             "GET",
		PathPattern:        "/security-group/{secg}/security-group-rule/{secgr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSecurityGroupRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSecurityGroupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSecurityGroupRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListSecurityGroupRule retrieves security group rule list from security group
*/
func (a *Client) ListSecurityGroupRule(params *ListSecurityGroupRuleParams) (*ListSecurityGroupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSecurityGroupRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSecurityGroupRule",
		Method:             "GET",
		PathPattern:        "/security-group/{secg}/security-group-rule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSecurityGroupRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSecurityGroupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSecurityGroupRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveSecurityGroupRule removes security group rule from security group
*/
func (a *Client) RemoveSecurityGroupRule(params *RemoveSecurityGroupRuleParams) (*RemoveSecurityGroupRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveSecurityGroupRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeSecurityGroupRule",
		Method:             "DELETE",
		PathPattern:        "/security-group/{secg}/security-group-rule/{secgr}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveSecurityGroupRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveSecurityGroupRuleNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeSecurityGroupRule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
