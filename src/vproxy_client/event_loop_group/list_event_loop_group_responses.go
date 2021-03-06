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

// ListEventLoopGroupReader is a Reader for the ListEventLoopGroup structure.
type ListEventLoopGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEventLoopGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEventLoopGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListEventLoopGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListEventLoopGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListEventLoopGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListEventLoopGroupOK creates a ListEventLoopGroupOK with default headers values
func NewListEventLoopGroupOK() *ListEventLoopGroupOK {
	return &ListEventLoopGroupOK{}
}

/*ListEventLoopGroupOK handles this case with default header values.

ok
*/
type ListEventLoopGroupOK struct {
	Payload []*vproxy_client_model.EventLoopGroup
}

func (o *ListEventLoopGroupOK) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupOK  %+v", 200, o.Payload)
}

func (o *ListEventLoopGroupOK) GetPayload() []*vproxy_client_model.EventLoopGroup {
	return o.Payload
}

func (o *ListEventLoopGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupBadRequest creates a ListEventLoopGroupBadRequest with default headers values
func NewListEventLoopGroupBadRequest() *ListEventLoopGroupBadRequest {
	return &ListEventLoopGroupBadRequest{}
}

/*ListEventLoopGroupBadRequest handles this case with default header values.

invalid input parameters
*/
type ListEventLoopGroupBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *ListEventLoopGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListEventLoopGroupBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *ListEventLoopGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupConflict creates a ListEventLoopGroupConflict with default headers values
func NewListEventLoopGroupConflict() *ListEventLoopGroupConflict {
	return &ListEventLoopGroupConflict{}
}

/*ListEventLoopGroupConflict handles this case with default header values.

conflict
*/
type ListEventLoopGroupConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *ListEventLoopGroupConflict) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupConflict  %+v", 409, o.Payload)
}

func (o *ListEventLoopGroupConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *ListEventLoopGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventLoopGroupInternalServerError creates a ListEventLoopGroupInternalServerError with default headers values
func NewListEventLoopGroupInternalServerError() *ListEventLoopGroupInternalServerError {
	return &ListEventLoopGroupInternalServerError{}
}

/*ListEventLoopGroupInternalServerError handles this case with default header values.

internal error
*/
type ListEventLoopGroupInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *ListEventLoopGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /event-loop-group][%d] listEventLoopGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListEventLoopGroupInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *ListEventLoopGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
