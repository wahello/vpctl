// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// ListSecurityGroupReader is a Reader for the ListSecurityGroup structure.
type ListSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSecurityGroupOK creates a ListSecurityGroupOK with default headers values
func NewListSecurityGroupOK() *ListSecurityGroupOK {
	return &ListSecurityGroupOK{}
}

/*ListSecurityGroupOK handles this case with default header values.

ok
*/
type ListSecurityGroupOK struct {
	Payload []*vproxy_client_model.SecurityGroup
}

func (o *ListSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /security-group][%d] listSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupOK) GetPayload() []*vproxy_client_model.SecurityGroup {
	return o.Payload
}

func (o *ListSecurityGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
