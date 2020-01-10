// Code generated by go-swagger; DO NOT EDIT.

package security_group_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// ListSecurityGroupRuleReader is a Reader for the ListSecurityGroupRule structure.
type ListSecurityGroupRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityGroupRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityGroupRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSecurityGroupRuleOK creates a ListSecurityGroupRuleOK with default headers values
func NewListSecurityGroupRuleOK() *ListSecurityGroupRuleOK {
	return &ListSecurityGroupRuleOK{}
}

/*ListSecurityGroupRuleOK handles this case with default header values.

ok
*/
type ListSecurityGroupRuleOK struct {
	Payload []*vproxy_client_model.SecurityGroupRule
}

func (o *ListSecurityGroupRuleOK) Error() string {
	return fmt.Sprintf("[GET /security-group/{secg}/security-group-rule][%d] listSecurityGroupRuleOK  %+v", 200, o.Payload)
}

func (o *ListSecurityGroupRuleOK) GetPayload() []*vproxy_client_model.SecurityGroupRule {
	return o.Payload
}

func (o *ListSecurityGroupRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
