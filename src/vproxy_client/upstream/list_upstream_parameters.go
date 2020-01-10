// Code generated by go-swagger; DO NOT EDIT.

package upstream

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

// NewListUpstreamParams creates a new ListUpstreamParams object
// with the default values initialized.
func NewListUpstreamParams() *ListUpstreamParams {

	return &ListUpstreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUpstreamParamsWithTimeout creates a new ListUpstreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUpstreamParamsWithTimeout(timeout time.Duration) *ListUpstreamParams {

	return &ListUpstreamParams{

		timeout: timeout,
	}
}

// NewListUpstreamParamsWithContext creates a new ListUpstreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUpstreamParamsWithContext(ctx context.Context) *ListUpstreamParams {

	return &ListUpstreamParams{

		Context: ctx,
	}
}

// NewListUpstreamParamsWithHTTPClient creates a new ListUpstreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUpstreamParamsWithHTTPClient(client *http.Client) *ListUpstreamParams {

	return &ListUpstreamParams{
		HTTPClient: client,
	}
}

/*ListUpstreamParams contains all the parameters to send to the API endpoint
for the list upstream operation typically these are written to a http.Request
*/
type ListUpstreamParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list upstream params
func (o *ListUpstreamParams) WithTimeout(timeout time.Duration) *ListUpstreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list upstream params
func (o *ListUpstreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list upstream params
func (o *ListUpstreamParams) WithContext(ctx context.Context) *ListUpstreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list upstream params
func (o *ListUpstreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list upstream params
func (o *ListUpstreamParams) WithHTTPClient(client *http.Client) *ListUpstreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list upstream params
func (o *ListUpstreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListUpstreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
