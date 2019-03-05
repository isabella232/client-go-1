// Code generated by go-swagger; DO NOT EDIT.

package oidc

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

// NewGetAuthOidcCallbackParams creates a new GetAuthOidcCallbackParams object
// with the default values initialized.
func NewGetAuthOidcCallbackParams() *GetAuthOidcCallbackParams {

	return &GetAuthOidcCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthOidcCallbackParamsWithTimeout creates a new GetAuthOidcCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthOidcCallbackParamsWithTimeout(timeout time.Duration) *GetAuthOidcCallbackParams {

	return &GetAuthOidcCallbackParams{

		timeout: timeout,
	}
}

// NewGetAuthOidcCallbackParamsWithContext creates a new GetAuthOidcCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthOidcCallbackParamsWithContext(ctx context.Context) *GetAuthOidcCallbackParams {

	return &GetAuthOidcCallbackParams{

		Context: ctx,
	}
}

// NewGetAuthOidcCallbackParamsWithHTTPClient creates a new GetAuthOidcCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthOidcCallbackParamsWithHTTPClient(client *http.Client) *GetAuthOidcCallbackParams {

	return &GetAuthOidcCallbackParams{
		HTTPClient: client,
	}
}

/*GetAuthOidcCallbackParams contains all the parameters to send to the API endpoint
for the get auth oidc callback operation typically these are written to a http.Request
*/
type GetAuthOidcCallbackParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) WithTimeout(timeout time.Duration) *GetAuthOidcCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) WithContext(ctx context.Context) *GetAuthOidcCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) WithHTTPClient(client *http.Client) *GetAuthOidcCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth oidc callback params
func (o *GetAuthOidcCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthOidcCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}