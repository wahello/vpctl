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

// DescribeTCPLbReader is a Reader for the DescribeTCPLb structure.
type DescribeTCPLbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeTCPLbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeTCPLbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDescribeTCPLbBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDescribeTCPLbNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDescribeTCPLbConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDescribeTCPLbInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeTCPLbOK creates a DescribeTCPLbOK with default headers values
func NewDescribeTCPLbOK() *DescribeTCPLbOK {
	return &DescribeTCPLbOK{}
}

/*DescribeTCPLbOK handles this case with default header values.

ok
*/
type DescribeTCPLbOK struct {
	Payload *vproxy_client_model.TCPLbDetail
}

func (o *DescribeTCPLbOK) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}/detail][%d] describeTcpLbOK  %+v", 200, o.Payload)
}

func (o *DescribeTCPLbOK) GetPayload() *vproxy_client_model.TCPLbDetail {
	return o.Payload
}

func (o *DescribeTCPLbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.TCPLbDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeTCPLbBadRequest creates a DescribeTCPLbBadRequest with default headers values
func NewDescribeTCPLbBadRequest() *DescribeTCPLbBadRequest {
	return &DescribeTCPLbBadRequest{}
}

/*DescribeTCPLbBadRequest handles this case with default header values.

invalid input parameters
*/
type DescribeTCPLbBadRequest struct {
	Payload *vproxy_client_model.Error400
}

func (o *DescribeTCPLbBadRequest) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}/detail][%d] describeTcpLbBadRequest  %+v", 400, o.Payload)
}

func (o *DescribeTCPLbBadRequest) GetPayload() *vproxy_client_model.Error400 {
	return o.Payload
}

func (o *DescribeTCPLbBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeTCPLbNotFound creates a DescribeTCPLbNotFound with default headers values
func NewDescribeTCPLbNotFound() *DescribeTCPLbNotFound {
	return &DescribeTCPLbNotFound{}
}

/*DescribeTCPLbNotFound handles this case with default header values.

resource not found
*/
type DescribeTCPLbNotFound struct {
	Payload *vproxy_client_model.Error404
}

func (o *DescribeTCPLbNotFound) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}/detail][%d] describeTcpLbNotFound  %+v", 404, o.Payload)
}

func (o *DescribeTCPLbNotFound) GetPayload() *vproxy_client_model.Error404 {
	return o.Payload
}

func (o *DescribeTCPLbNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error404)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeTCPLbConflict creates a DescribeTCPLbConflict with default headers values
func NewDescribeTCPLbConflict() *DescribeTCPLbConflict {
	return &DescribeTCPLbConflict{}
}

/*DescribeTCPLbConflict handles this case with default header values.

conflict
*/
type DescribeTCPLbConflict struct {
	Payload *vproxy_client_model.Error409
}

func (o *DescribeTCPLbConflict) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}/detail][%d] describeTcpLbConflict  %+v", 409, o.Payload)
}

func (o *DescribeTCPLbConflict) GetPayload() *vproxy_client_model.Error409 {
	return o.Payload
}

func (o *DescribeTCPLbConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error409)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeTCPLbInternalServerError creates a DescribeTCPLbInternalServerError with default headers values
func NewDescribeTCPLbInternalServerError() *DescribeTCPLbInternalServerError {
	return &DescribeTCPLbInternalServerError{}
}

/*DescribeTCPLbInternalServerError handles this case with default header values.

internal error
*/
type DescribeTCPLbInternalServerError struct {
	Payload *vproxy_client_model.Error500
}

func (o *DescribeTCPLbInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tcp-lb/{tl}/detail][%d] describeTcpLbInternalServerError  %+v", 500, o.Payload)
}

func (o *DescribeTCPLbInternalServerError) GetPayload() *vproxy_client_model.Error500 {
	return o.Payload
}

func (o *DescribeTCPLbInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.Error500)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
