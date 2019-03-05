// Code generated by go-swagger; DO NOT EDIT.

package saml

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

// NewPatchAuthSamlProvidersProviderIDParams creates a new PatchAuthSamlProvidersProviderIDParams object
// with the default values initialized.
func NewPatchAuthSamlProvidersProviderIDParams() *PatchAuthSamlProvidersProviderIDParams {
	var ()
	return &PatchAuthSamlProvidersProviderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAuthSamlProvidersProviderIDParamsWithTimeout creates a new PatchAuthSamlProvidersProviderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAuthSamlProvidersProviderIDParamsWithTimeout(timeout time.Duration) *PatchAuthSamlProvidersProviderIDParams {
	var ()
	return &PatchAuthSamlProvidersProviderIDParams{

		timeout: timeout,
	}
}

// NewPatchAuthSamlProvidersProviderIDParamsWithContext creates a new PatchAuthSamlProvidersProviderIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAuthSamlProvidersProviderIDParamsWithContext(ctx context.Context) *PatchAuthSamlProvidersProviderIDParams {
	var ()
	return &PatchAuthSamlProvidersProviderIDParams{

		Context: ctx,
	}
}

// NewPatchAuthSamlProvidersProviderIDParamsWithHTTPClient creates a new PatchAuthSamlProvidersProviderIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAuthSamlProvidersProviderIDParamsWithHTTPClient(client *http.Client) *PatchAuthSamlProvidersProviderIDParams {
	var ()
	return &PatchAuthSamlProvidersProviderIDParams{
		HTTPClient: client,
	}
}

/*PatchAuthSamlProvidersProviderIDParams contains all the parameters to send to the API endpoint
for the patch auth saml providers provider ID operation typically these are written to a http.Request
*/
type PatchAuthSamlProvidersProviderIDParams struct {

	/*ProviderConfigObject
	  Provider config JSON object

	*/
	ProviderConfigObject *models.SAMLProviderConfig
	/*ProviderID
	  The ID of the provider to modify.

	*/
	ProviderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) WithTimeout(timeout time.Duration) *PatchAuthSamlProvidersProviderIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) WithContext(ctx context.Context) *PatchAuthSamlProvidersProviderIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) WithHTTPClient(client *http.Client) *PatchAuthSamlProvidersProviderIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProviderConfigObject adds the providerConfigObject to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) WithProviderConfigObject(providerConfigObject *models.SAMLProviderConfig) *PatchAuthSamlProvidersProviderIDParams {
	o.SetProviderConfigObject(providerConfigObject)
	return o
}

// SetProviderConfigObject adds the providerConfigObject to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) SetProviderConfigObject(providerConfigObject *models.SAMLProviderConfig) {
	o.ProviderConfigObject = providerConfigObject
}

// WithProviderID adds the providerID to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) WithProviderID(providerID string) *PatchAuthSamlProvidersProviderIDParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the patch auth saml providers provider ID params
func (o *PatchAuthSamlProvidersProviderIDParams) SetProviderID(providerID string) {
	o.ProviderID = providerID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAuthSamlProvidersProviderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProviderConfigObject != nil {
		if err := r.SetBodyParam(o.ProviderConfigObject); err != nil {
			return err
		}
	}

	// path param provider-id
	if err := r.SetPathParam("provider-id", o.ProviderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}