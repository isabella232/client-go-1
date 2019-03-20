/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

// This describes what backends to send traffic to. This can be expressed with a variety of filters such as matching on the hostname or the HTTP URL path.
type V2FrontendLinkBackend struct {
	// This is default backend that is routed to if none of the other filters are matched.
	DefaultBackend string `json:"defaultBackend,omitempty"`
	// This is an optional field that specifies a mapping to various backends. These rules are applied in order.
	Map []V2FrontendLinkBackendMap `json:"map,omitempty"`
}