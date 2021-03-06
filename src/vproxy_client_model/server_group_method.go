// Code generated by go-swagger; DO NOT EDIT.

package vproxy_client_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ServerGroupMethod the loadbalancing method
// swagger:model ServerGroupMethod
type ServerGroupMethod string

const (

	// ServerGroupMethodWrr captures enum value "wrr"
	ServerGroupMethodWrr ServerGroupMethod = "wrr"

	// ServerGroupMethodWlc captures enum value "wlc"
	ServerGroupMethodWlc ServerGroupMethod = "wlc"

	// ServerGroupMethodSource captures enum value "source"
	ServerGroupMethodSource ServerGroupMethod = "source"
)

// for schema
var serverGroupMethodEnum []interface{}

func init() {
	var res []ServerGroupMethod
	if err := json.Unmarshal([]byte(`["wrr","wlc","source"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverGroupMethodEnum = append(serverGroupMethodEnum, v)
	}
}

func (m ServerGroupMethod) validateServerGroupMethodEnum(path, location string, value ServerGroupMethod) error {
	if err := validate.Enum(path, location, value, serverGroupMethodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this server group method
func (m ServerGroupMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServerGroupMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
