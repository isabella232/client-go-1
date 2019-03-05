// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewDeleteLdapConfigParams creates a new DeleteLdapConfigParams object
// with the default values initialized.
func NewDeleteLdapConfigParams() *DeleteLdapConfigParams {

	return &DeleteLdapConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLdapConfigParamsWithTimeout creates a new DeleteLdapConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLdapConfigParamsWithTimeout(timeout time.Duration) *DeleteLdapConfigParams {

	return &DeleteLdapConfigParams{

		timeout: timeout,
	}
}

// NewDeleteLdapConfigParamsWithContext creates a new DeleteLdapConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLdapConfigParamsWithContext(ctx context.Context) *DeleteLdapConfigParams {

	return &DeleteLdapConfigParams{

		Context: ctx,
	}
}

// NewDeleteLdapConfigParamsWithHTTPClient creates a new DeleteLdapConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLdapConfigParamsWithHTTPClient(client *http.Client) *DeleteLdapConfigParams {

	return &DeleteLdapConfigParams{
		HTTPClient: client,
	}
}

/*DeleteLdapConfigParams contains all the parameters to send to the API endpoint
for the delete ldap config operation typically these are written to a http.Request
*/
type DeleteLdapConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ldap config params
func (o *DeleteLdapConfigParams) WithTimeout(timeout time.Duration) *DeleteLdapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ldap config params
func (o *DeleteLdapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ldap config params
func (o *DeleteLdapConfigParams) WithContext(ctx context.Context) *DeleteLdapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ldap config params
func (o *DeleteLdapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ldap config params
func (o *DeleteLdapConfigParams) WithHTTPClient(client *http.Client) *DeleteLdapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ldap config params
func (o *DeleteLdapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLdapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}