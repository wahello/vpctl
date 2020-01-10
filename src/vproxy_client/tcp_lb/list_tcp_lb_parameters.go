// Code generated by go-swagger; DO NOT EDIT.

package tcp_lb

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
)

// NewListTCPLbParams creates a new ListTCPLbParams object
// with the default values initialized.
func NewListTCPLbParams() *ListTCPLbParams {

	return &ListTCPLbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTCPLbParamsWithTimeout creates a new ListTCPLbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTCPLbParamsWithTimeout(timeout time.Duration) *ListTCPLbParams {

	return &ListTCPLbParams{

		timeout: timeout,
	}
}

// NewListTCPLbParamsWithContext creates a new ListTCPLbParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTCPLbParamsWithContext(ctx context.Context) *ListTCPLbParams {

	return &ListTCPLbParams{

		Context: ctx,
	}
}

// NewListTCPLbParamsWithHTTPClient creates a new ListTCPLbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTCPLbParamsWithHTTPClient(client *http.Client) *ListTCPLbParams {

	return &ListTCPLbParams{
		HTTPClient: client,
	}
}

/*ListTCPLbParams contains all the parameters to send to the API endpoint
for the list Tcp lb operation typically these are written to a http.Request
*/
type ListTCPLbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list Tcp lb params
func (o *ListTCPLbParams) WithTimeout(timeout time.Duration) *ListTCPLbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list Tcp lb params
func (o *ListTCPLbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list Tcp lb params
func (o *ListTCPLbParams) WithContext(ctx context.Context) *ListTCPLbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list Tcp lb params
func (o *ListTCPLbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list Tcp lb params
func (o *ListTCPLbParams) WithHTTPClient(client *http.Client) *ListTCPLbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list Tcp lb params
func (o *ListTCPLbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTCPLbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
