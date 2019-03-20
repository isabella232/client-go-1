/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

// Setting this to the empty object is enough to redirect all traffic from HTTP (this frontend) to HTTPS (port 443).
type V2FrontendRedirectToHttps struct {
	// One may additionally set a whitelist of fields that must be matched to allow HTTP.
	Except []V2FrontendRedirectToHttpsExcept `json:"except,omitempty"`
}