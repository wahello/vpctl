// Code generated by go-swagger; DO NOT EDIT.

package server_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// DescribeServerGroupInUpstreamReader is a Reader for the DescribeServerGroupInUpstream structure.
type DescribeServerGroupInUpstreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeServerGroupInUpstreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeServerGroupInUpstreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDescribeServerGroupInUpstreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeServerGroupInUpstreamOK creates a DescribeServerGroupInUpstreamOK with default headers values
func NewDescribeServerGroupInUpstreamOK() *DescribeServerGroupInUpstreamOK {
	return &DescribeServerGroupInUpstreamOK{}
}

/*DescribeServerGroupInUpstreamOK handles this case with default header values.

ok
*/
type DescribeServerGroupInUpstreamOK struct {
	Payload *vproxy_client_model.ServerGroupInUpstreamDetail
}

func (o *DescribeServerGroupInUpstreamOK) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/server-group/{sg}/detail][%d] describeServerGroupInUpstreamOK  %+v", 200, o.Payload)
}

func (o *DescribeServerGroupInUpstreamOK) GetPayload() *vproxy_client_model.ServerGroupInUpstreamDetail {
	return o.Payload
}

func (o *DescribeServerGroupInUpstreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.ServerGroupInUpstreamDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerGroupInUpstreamNotFound creates a DescribeServerGroupInUpstreamNotFound with default headers values
func NewDescribeServerGroupInUpstreamNotFound() *DescribeServerGroupInUpstreamNotFound {
	return &DescribeServerGroupInUpstreamNotFound{}
}

/*DescribeServerGroupInUpstreamNotFound handles this case with default header values.

not found
*/
type DescribeServerGroupInUpstreamNotFound struct {
}

func (o *DescribeServerGroupInUpstreamNotFound) Error() string {
	return fmt.Sprintf("[GET /upstream/{ups}/server-group/{sg}/detail][%d] describeServerGroupInUpstreamNotFound ", 404)
}

func (o *DescribeServerGroupInUpstreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
