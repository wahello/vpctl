// Code generated by go-swagger; DO NOT EDIT.

package dns_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// AddDNSServerReader is a Reader for the AddDNSServer structure.
type AddDNSServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDNSServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddDNSServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDNSServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddDNSServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddDNSServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDNSServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDNSServerNoContent creates a AddDNSServerNoContent with default headers values
func NewAddDNSServerNoContent() *AddDNSServerNoContent {
	return &AddDNSServerNoContent{}
}

/*AddDNSServerNoContent handles this case with default header values.

ok
*/
type AddDNSServerNoContent struct {
}

func (o *AddDNSServerNoContent) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNoContent ", 204)
}

func (o *AddDNSServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddDNSServerBadRequest creates a AddDNSServerBadRequest with default headers values
func NewAddDNSServerBadRequest() *AddDNSServerBadRequest {
	return &AddDNSServerBadRequest{}
}

/*AddDNSServerBadRequest handles this case with default header values.

invalid input parameters
*/
type AddDNSServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *AddDNSServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *AddDNSServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *AddDNSServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerNotFound creates a AddDNSServerNotFound with default headers values
func NewAddDNSServerNotFound() *AddDNSServerNotFound {
	return &AddDNSServerNotFound{}
}

/*AddDNSServerNotFound handles this case with default header values.

resource not found
*/
type AddDNSServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *AddDNSServerNotFound) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *AddDNSServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *AddDNSServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerConflict creates a AddDNSServerConflict with default headers values
func NewAddDNSServerConflict() *AddDNSServerConflict {
	return &AddDNSServerConflict{}
}

/*AddDNSServerConflict handles this case with default header values.

conflict
*/
type AddDNSServerConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *AddDNSServerConflict) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerConflict  %+v", 409, o.Payload)
}

func (o *AddDNSServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *AddDNSServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDNSServerInternalServerError creates a AddDNSServerInternalServerError with default headers values
func NewAddDNSServerInternalServerError() *AddDNSServerInternalServerError {
	return &AddDNSServerInternalServerError{}
}

/*AddDNSServerInternalServerError handles this case with default header values.

internal error
*/
type AddDNSServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *AddDNSServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dns-server][%d] addDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDNSServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *AddDNSServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
