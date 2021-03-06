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

// RemoveDNSServerReader is a Reader for the RemoveDNSServer structure.
type RemoveDNSServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveDNSServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveDNSServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveDNSServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveDNSServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveDNSServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveDNSServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveDNSServerNoContent creates a RemoveDNSServerNoContent with default headers values
func NewRemoveDNSServerNoContent() *RemoveDNSServerNoContent {
	return &RemoveDNSServerNoContent{}
}

/*RemoveDNSServerNoContent handles this case with default header values.

ok
*/
type RemoveDNSServerNoContent struct {
}

func (o *RemoveDNSServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dns-server/{dns}][%d] removeDnsServerNoContent ", 204)
}

func (o *RemoveDNSServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveDNSServerBadRequest creates a RemoveDNSServerBadRequest with default headers values
func NewRemoveDNSServerBadRequest() *RemoveDNSServerBadRequest {
	return &RemoveDNSServerBadRequest{}
}

/*RemoveDNSServerBadRequest handles this case with default header values.

invalid input parameters
*/
type RemoveDNSServerBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *RemoveDNSServerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dns-server/{dns}][%d] removeDnsServerBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveDNSServerBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *RemoveDNSServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveDNSServerNotFound creates a RemoveDNSServerNotFound with default headers values
func NewRemoveDNSServerNotFound() *RemoveDNSServerNotFound {
	return &RemoveDNSServerNotFound{}
}

/*RemoveDNSServerNotFound handles this case with default header values.

resource not found
*/
type RemoveDNSServerNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *RemoveDNSServerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /dns-server/{dns}][%d] removeDnsServerNotFound  %+v", 404, o.Payload)
}

func (o *RemoveDNSServerNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *RemoveDNSServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveDNSServerConflict creates a RemoveDNSServerConflict with default headers values
func NewRemoveDNSServerConflict() *RemoveDNSServerConflict {
	return &RemoveDNSServerConflict{}
}

/*RemoveDNSServerConflict handles this case with default header values.

conflict
*/
type RemoveDNSServerConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *RemoveDNSServerConflict) Error() string {
	return fmt.Sprintf("[DELETE /dns-server/{dns}][%d] removeDnsServerConflict  %+v", 409, o.Payload)
}

func (o *RemoveDNSServerConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *RemoveDNSServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveDNSServerInternalServerError creates a RemoveDNSServerInternalServerError with default headers values
func NewRemoveDNSServerInternalServerError() *RemoveDNSServerInternalServerError {
	return &RemoveDNSServerInternalServerError{}
}

/*RemoveDNSServerInternalServerError handles this case with default header values.

internal error
*/
type RemoveDNSServerInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *RemoveDNSServerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dns-server/{dns}][%d] removeDnsServerInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveDNSServerInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *RemoveDNSServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
