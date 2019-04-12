/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type EdgelbV2Service struct {
	Marathon EdgelbV2ServiceMarathon `json:"marathon,omitempty"`
	Mesos    EdgelbV2ServiceMesos    `json:"mesos,omitempty"`
	Endpoint EdgelbV2Endpoint        `json:"endpoint,omitempty"`
}
