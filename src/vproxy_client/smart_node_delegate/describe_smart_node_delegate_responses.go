// Code generated by go-swagger; DO NOT EDIT.

package smart_node_delegate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// DescribeSmartNodeDelegateReader is a Reader for the DescribeSmartNodeDelegate structure.
type DescribeSmartNodeDelegateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeSmartNodeDelegateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeSmartNodeDelegateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDescribeSmartNodeDelegateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeSmartNodeDelegateOK creates a DescribeSmartNodeDelegateOK with default headers values
func NewDescribeSmartNodeDelegateOK() *DescribeSmartNodeDelegateOK {
	return &DescribeSmartNodeDelegateOK{}
}

/*DescribeSmartNodeDelegateOK handles this case with default header values.

ok
*/
type DescribeSmartNodeDelegateOK struct {
	Payload *vproxy_client_model.SmartNodeDelegateDetail
}

func (o *DescribeSmartNodeDelegateOK) Error() string {
	return fmt.Sprintf("[GET /smart-node-delegate/{snd}/detail][%d] describeSmartNodeDelegateOK  %+v", 200, o.Payload)
}

func (o *DescribeSmartNodeDelegateOK) GetPayload() *vproxy_client_model.SmartNodeDelegateDetail {
	return o.Payload
}

func (o *DescribeSmartNodeDelegateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(vproxy_client_model.SmartNodeDelegateDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeSmartNodeDelegateNotFound creates a DescribeSmartNodeDelegateNotFound with default headers values
func NewDescribeSmartNodeDelegateNotFound() *DescribeSmartNodeDelegateNotFound {
	return &DescribeSmartNodeDelegateNotFound{}
}

/*DescribeSmartNodeDelegateNotFound handles this case with default header values.

not found
*/
type DescribeSmartNodeDelegateNotFound struct {
}

func (o *DescribeSmartNodeDelegateNotFound) Error() string {
	return fmt.Sprintf("[GET /smart-node-delegate/{snd}/detail][%d] describeSmartNodeDelegateNotFound ", 404)
}

func (o *DescribeSmartNodeDelegateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
