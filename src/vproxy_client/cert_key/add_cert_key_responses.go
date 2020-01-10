// Code generated by go-swagger; DO NOT EDIT.

package cert_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddCertKeyReader is a Reader for the AddCertKey structure.
type AddCertKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCertKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddCertKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddCertKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddCertKeyNoContent creates a AddCertKeyNoContent with default headers values
func NewAddCertKeyNoContent() *AddCertKeyNoContent {
	return &AddCertKeyNoContent{}
}

/*AddCertKeyNoContent handles this case with default header values.

ok
*/
type AddCertKeyNoContent struct {
}

func (o *AddCertKeyNoContent) Error() string {
	return fmt.Sprintf("[POST /cert-key][%d] addCertKeyNoContent ", 204)
}

func (o *AddCertKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddCertKeyBadRequest creates a AddCertKeyBadRequest with default headers values
func NewAddCertKeyBadRequest() *AddCertKeyBadRequest {
	return &AddCertKeyBadRequest{}
}

/*AddCertKeyBadRequest handles this case with default header values.

Invalid input
*/
type AddCertKeyBadRequest struct {
}

func (o *AddCertKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /cert-key][%d] addCertKeyBadRequest ", 400)
}

func (o *AddCertKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
