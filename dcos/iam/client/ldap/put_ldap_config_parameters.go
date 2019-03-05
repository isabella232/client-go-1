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

	models "github.com/dcos/client-go/dcos/iam/models"
)

// NewPutLdapConfigParams creates a new PutLdapConfigParams object
// with the default values initialized.
func NewPutLdapConfigParams() *PutLdapConfigParams {
	var ()
	return &PutLdapConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLdapConfigParamsWithTimeout creates a new PutLdapConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLdapConfigParamsWithTimeout(timeout time.Duration) *PutLdapConfigParams {
	var ()
	return &PutLdapConfigParams{

		timeout: timeout,
	}
}

// NewPutLdapConfigParamsWithContext creates a new PutLdapConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLdapConfigParamsWithContext(ctx context.Context) *PutLdapConfigParams {
	var ()
	return &PutLdapConfigParams{

		Context: ctx,
	}
}

// NewPutLdapConfigParamsWithHTTPClient creates a new PutLdapConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLdapConfigParamsWithHTTPClient(client *http.Client) *PutLdapConfigParams {
	var ()
	return &PutLdapConfigParams{
		HTTPClient: client,
	}
}

/*PutLdapConfigParams contains all the parameters to send to the API endpoint
for the put ldap config operation typically these are written to a http.Request
*/
type PutLdapConfigParams struct {

	/*LDAPConfiguration
	  JSON object containing the LDAP configuration details.

	*/
	LDAPConfiguration *models.LDAPConfiguration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put ldap config params
func (o *PutLdapConfigParams) WithTimeout(timeout time.Duration) *PutLdapConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ldap config params
func (o *PutLdapConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ldap config params
func (o *PutLdapConfigParams) WithContext(ctx context.Context) *PutLdapConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ldap config params
func (o *PutLdapConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ldap config params
func (o *PutLdapConfigParams) WithHTTPClient(client *http.Client) *PutLdapConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ldap config params
func (o *PutLdapConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLDAPConfiguration adds the lDAPConfiguration to the put ldap config params
func (o *PutLdapConfigParams) WithLDAPConfiguration(lDAPConfiguration *models.LDAPConfiguration) *PutLdapConfigParams {
	o.SetLDAPConfiguration(lDAPConfiguration)
	return o
}

// SetLDAPConfiguration adds the lDAPConfiguration to the put ldap config params
func (o *PutLdapConfigParams) SetLDAPConfiguration(lDAPConfiguration *models.LDAPConfiguration) {
	o.LDAPConfiguration = lDAPConfiguration
}

// WriteToRequest writes these params to a swagger request
func (o *PutLdapConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LDAPConfiguration != nil {
		if err := r.SetBodyParam(o.LDAPConfiguration); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}