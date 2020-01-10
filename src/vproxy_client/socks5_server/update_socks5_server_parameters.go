// Code generated by go-swagger; DO NOT EDIT.

package socks5_server

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

// NewUpdateSocks5ServerParams creates a new UpdateSocks5ServerParams object
// with the default values initialized.
func NewUpdateSocks5ServerParams() *UpdateSocks5ServerParams {
	var ()
	return &UpdateSocks5ServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSocks5ServerParamsWithTimeout creates a new UpdateSocks5ServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSocks5ServerParamsWithTimeout(timeout time.Duration) *UpdateSocks5ServerParams {
	var ()
	return &UpdateSocks5ServerParams{

		timeout: timeout,
	}
}

// NewUpdateSocks5ServerParamsWithContext creates a new UpdateSocks5ServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSocks5ServerParamsWithContext(ctx context.Context) *UpdateSocks5ServerParams {
	var ()
	return &UpdateSocks5ServerParams{

		Context: ctx,
	}
}

// NewUpdateSocks5ServerParamsWithHTTPClient creates a new UpdateSocks5ServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSocks5ServerParamsWithHTTPClient(client *http.Client) *UpdateSocks5ServerParams {
	var ()
	return &UpdateSocks5ServerParams{
		HTTPClient: client,
	}
}

/*UpdateSocks5ServerParams contains all the parameters to send to the API endpoint
for the update socks5 server operation typically these are written to a http.Request
*/
type UpdateSocks5ServerParams struct {

	/*Body*/
	Body *vproxy_client_model.Socks5ServerUpdate
	/*Socks5
	  name of the socks5-server

	*/
	Socks5 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update socks5 server params
func (o *UpdateSocks5ServerParams) WithTimeout(timeout time.Duration) *UpdateSocks5ServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update socks5 server params
func (o *UpdateSocks5ServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update socks5 server params
func (o *UpdateSocks5ServerParams) WithContext(ctx context.Context) *UpdateSocks5ServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update socks5 server params
func (o *UpdateSocks5ServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update socks5 server params
func (o *UpdateSocks5ServerParams) WithHTTPClient(client *http.Client) *UpdateSocks5ServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update socks5 server params
func (o *UpdateSocks5ServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update socks5 server params
func (o *UpdateSocks5ServerParams) WithBody(body *vproxy_client_model.Socks5ServerUpdate) *UpdateSocks5ServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update socks5 server params
func (o *UpdateSocks5ServerParams) SetBody(body *vproxy_client_model.Socks5ServerUpdate) {
	o.Body = body
}

// WithSocks5 adds the socks5 to the update socks5 server params
func (o *UpdateSocks5ServerParams) WithSocks5(socks5 string) *UpdateSocks5ServerParams {
	o.SetSocks5(socks5)
	return o
}

// SetSocks5 adds the socks5 to the update socks5 server params
func (o *UpdateSocks5ServerParams) SetSocks5(socks5 string) {
	o.Socks5 = socks5
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSocks5ServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param socks5
	if err := r.SetPathParam("socks5", o.Socks5); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
