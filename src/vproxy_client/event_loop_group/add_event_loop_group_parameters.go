// Code generated by go-swagger; DO NOT EDIT.

package event_loop_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	vproxy_client_model "vproxy_client_model"
)

// NewAddEventLoopGroupParams creates a new AddEventLoopGroupParams object
// with the default values initialized.
func NewAddEventLoopGroupParams() *AddEventLoopGroupParams {
	var ()
	return &AddEventLoopGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEventLoopGroupParamsWithTimeout creates a new AddEventLoopGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEventLoopGroupParamsWithTimeout(timeout time.Duration) *AddEventLoopGroupParams {
	var ()
	return &AddEventLoopGroupParams{

		timeout: timeout,
	}
}

// NewAddEventLoopGroupParamsWithContext creates a new AddEventLoopGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEventLoopGroupParamsWithContext(ctx context.Context) *AddEventLoopGroupParams {
	var ()
	return &AddEventLoopGroupParams{

		Context: ctx,
	}
}

// NewAddEventLoopGroupParamsWithHTTPClient creates a new AddEventLoopGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEventLoopGroupParamsWithHTTPClient(client *http.Client) *AddEventLoopGroupParams {
	var ()
	return &AddEventLoopGroupParams{
		HTTPClient: client,
	}
}

/*AddEventLoopGroupParams contains all the parameters to send to the API endpoint
for the add event loop group operation typically these are written to a http.Request
*/
type AddEventLoopGroupParams struct {

	/*Body*/
	Body *vproxy_client_model.EventLoopGroupCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add event loop group params
func (o *AddEventLoopGroupParams) WithTimeout(timeout time.Duration) *AddEventLoopGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add event loop group params
func (o *AddEventLoopGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add event loop group params
func (o *AddEventLoopGroupParams) WithContext(ctx context.Context) *AddEventLoopGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add event loop group params
func (o *AddEventLoopGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add event loop group params
func (o *AddEventLoopGroupParams) WithHTTPClient(client *http.Client) *AddEventLoopGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add event loop group params
func (o *AddEventLoopGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add event loop group params
func (o *AddEventLoopGroupParams) WithBody(body *vproxy_client_model.EventLoopGroupCreate) *AddEventLoopGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add event loop group params
func (o *AddEventLoopGroupParams) SetBody(body *vproxy_client_model.EventLoopGroupCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddEventLoopGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
