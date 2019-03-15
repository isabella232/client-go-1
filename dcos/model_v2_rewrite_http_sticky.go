/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

// Sticky sessions via a cookie. To use the default values (recommended), set this field to the empty object.
type V2RewriteHttpSticky struct {
	Enabled   bool   `json:"enabled,omitempty"`
	CustomStr string `json:"customStr,omitempty"`
}
