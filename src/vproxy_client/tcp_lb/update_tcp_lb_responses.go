// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// UpdateTCPLbReader is a Reader for the UpdateTCPLb structure.
type UpdateTCPLbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTCPLbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateTCPLbNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTCPLbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTCPLbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTCPLbConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTCPLbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTCPLbNoContent creates a UpdateTCPLbNoContent with default headers values
func NewUpdateTCPLbNoContent() *UpdateTCPLbNoContent {
	return &UpdateTCPLbNoContent{}
}

/*UpdateTCPLbNoContent handles this case with default header values.

ok
*/
type UpdateTCPLbNoContent struct {
}

func (o *UpdateTCPLbNoContent) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNoContent ", 204)
}

func (o *UpdateTCPLbNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTCPLbBadRequest creates a UpdateTCPLbBadRequest with default headers values
func NewUpdateTCPLbBadRequest() *UpdateTCPLbBadRequest {
	return &UpdateTCPLbBadRequest{}
}

/*UpdateTCPLbBadRequest handles this case with default header values.

invalid input parameters
*/
type UpdateTCPLbBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *UpdateTCPLbBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTCPLbBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *UpdateTCPLbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbNotFound creates a UpdateTCPLbNotFound with default headers values
func NewUpdateTCPLbNotFound() *UpdateTCPLbNotFound {
	return &UpdateTCPLbNotFound{}
}

/*UpdateTCPLbNotFound handles this case with default header values.

resource not found
*/
type UpdateTCPLbNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *UpdateTCPLbNotFound) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTCPLbNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *UpdateTCPLbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbConflict creates a UpdateTCPLbConflict with default headers values
func NewUpdateTCPLbConflict() *UpdateTCPLbConflict {
	return &UpdateTCPLbConflict{}
}

/*UpdateTCPLbConflict handles this case with default header values.

conflict
*/
type UpdateTCPLbConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *UpdateTCPLbConflict) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbConflict  %+v", 409, o.Payload)
}

func (o *UpdateTCPLbConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *UpdateTCPLbConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTCPLbInternalServerError creates a UpdateTCPLbInternalServerError with default headers values
func NewUpdateTCPLbInternalServerError() *UpdateTCPLbInternalServerError {
	return &UpdateTCPLbInternalServerError{}
}

/*UpdateTCPLbInternalServerError handles this case with default header values.

internal error
*/
type UpdateTCPLbInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *UpdateTCPLbInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tcp-lb/{tl}][%d] updateTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTCPLbInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *UpdateTCPLbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
