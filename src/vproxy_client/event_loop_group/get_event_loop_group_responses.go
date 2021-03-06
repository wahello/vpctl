// Code generated by go-swagger; DO NOT EDIT.

package event_loop_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// GetEventLoopGroupReader is a Reader for the GetEventLoopGroup structure.
type GetEventLoopGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventLoopGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventLoopGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventLoopGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventLoopGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetEventLoopGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEventLoopGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventLoopGroupOK creates a GetEventLoopGroupOK with default headers values
func NewGetEventLoopGroupOK() *GetEventLoopGroupOK {
	return &GetEventLoopGroupOK{}
}

/*GetEventLoopGroupOK handles this case with default header values.

ok
*/
type GetEventLoopGroupOK struct {
	Payload *vproxy_client_model.EventLoopGroup
}

func (o *GetEventLoopGroupOK) Error() string {
	return fmt.Sprintf("[GET /event-loop-group/{elg}][%d] getEventLoopGroupOK  %+v", 200, o.Payload)
}

func (o *GetEventLoopGroupOK) GetPayload() *vproxy_client_model.EventLoopGroup {
	return o.Payload
}

func (o *GetEventLoopGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.EventLoopGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventLoopGroupBadRequest creates a GetEventLoopGroupBadRequest with default headers values
func NewGetEventLoopGroupBadRequest() *GetEventLoopGroupBadRequest {
	return &GetEventLoopGroupBadRequest{}
}

/*GetEventLoopGroupBadRequest handles this case with default header values.

invalid input parameters
*/
type GetEventLoopGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *GetEventLoopGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /event-loop-group/{elg}][%d] getEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetEventLoopGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *GetEventLoopGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventLoopGroupNotFound creates a GetEventLoopGroupNotFound with default headers values
func NewGetEventLoopGroupNotFound() *GetEventLoopGroupNotFound {
	return &GetEventLoopGroupNotFound{}
}

/*GetEventLoopGroupNotFound handles this case with default header values.

resource not found
*/
type GetEventLoopGroupNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *GetEventLoopGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /event-loop-group/{elg}][%d] getEventLoopGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetEventLoopGroupNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *GetEventLoopGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventLoopGroupConflict creates a GetEventLoopGroupConflict with default headers values
func NewGetEventLoopGroupConflict() *GetEventLoopGroupConflict {
	return &GetEventLoopGroupConflict{}
}

/*GetEventLoopGroupConflict handles this case with default header values.

conflict
*/
type GetEventLoopGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *GetEventLoopGroupConflict) Error() string {
	return fmt.Sprintf("[GET /event-loop-group/{elg}][%d] getEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *GetEventLoopGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *GetEventLoopGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventLoopGroupInternalServerError creates a GetEventLoopGroupInternalServerError with default headers values
func NewGetEventLoopGroupInternalServerError() *GetEventLoopGroupInternalServerError {
	return &GetEventLoopGroupInternalServerError{}
}

/*GetEventLoopGroupInternalServerError handles this case with default header values.

internal error
*/
type GetEventLoopGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *GetEventLoopGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /event-loop-group/{elg}][%d] getEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetEventLoopGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *GetEventLoopGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
