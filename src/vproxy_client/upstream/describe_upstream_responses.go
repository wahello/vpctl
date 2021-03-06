// Code generated by go-swagger; DO NOT EDIT.

package upstream

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// DescribeUpstreamReader is a Reader for the DescribeUpstream structure.
type DescribeUpstreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeUpstreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeUpstreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeUpstreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeUpstreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeUpstreamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeUpstreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeUpstreamOK creates a DescribeUpstreamOK with default headers values
func NewDescribeUpstreamOK() *DescribeUpstreamOK {
	return &DescribeUpstreamOK{}
}

/*DescribeUpstreamOK handles this case with default header values.

ok
*/
type DescribeUpstreamOK struct {
	Payload *vproxy_client_model.UpstreamDetail
}

func (o *DescribeUpstreamOK) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamOK  %+v", 200, o.Payload)
}

func (o *DescribeUpstreamOK) GetPayload() *vproxy_client_model.UpstreamDetail {
	return o.Payload
}

func (o *DescribeUpstreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.UpstreamDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamBadRequest creates a DescribeUpstreamBadRequest with default headers values
func NewDescribeUpstreamBadRequest() *DescribeUpstreamBadRequest {
	return &DescribeUpstreamBadRequest{}
}

/*DescribeUpstreamBadRequest handles this case with default header values.

invalid input parameters
*/
type DescribeUpstreamBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *DescribeUpstreamBadRequest) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeUpstreamBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeUpstreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamNotFound creates a DescribeUpstreamNotFound with default headers values
func NewDescribeUpstreamNotFound() *DescribeUpstreamNotFound {
	return &DescribeUpstreamNotFound{}
}

/*DescribeUpstreamNotFound handles this case with default header values.

resource not found
*/
type DescribeUpstreamNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *DescribeUpstreamNotFound) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamNotFound  %+v", 404, o.Payload)
}

func (o *DescribeUpstreamNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeUpstreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamConflict creates a DescribeUpstreamConflict with default headers values
func NewDescribeUpstreamConflict() *DescribeUpstreamConflict {
	return &DescribeUpstreamConflict{}
}

/*DescribeUpstreamConflict handles this case with default header values.

conflict
*/
type DescribeUpstreamConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *DescribeUpstreamConflict) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamConflict  %+v", 409, o.Payload)
}

func (o *DescribeUpstreamConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeUpstreamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeUpstreamInternalServerError creates a DescribeUpstreamInternalServerError with default headers values
func NewDescribeUpstreamInternalServerError() *DescribeUpstreamInternalServerError {
	return &DescribeUpstreamInternalServerError{}
}

/*DescribeUpstreamInternalServerError handles this case with default header values.

internal error
*/
type DescribeUpstreamInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *DescribeUpstreamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/detail][%d] describeUpstreamInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeUpstreamInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeUpstreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
