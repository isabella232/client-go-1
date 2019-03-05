// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LDAPConfiguration l d a p configuration
// swagger:model LDAPConfiguration
type LDAPConfiguration struct {

	// ca certs
	CaCerts string `json:"ca-certs,omitempty"`

	// client cert
	ClientCert string `json:"client-cert,omitempty"`

	// dntemplate
	Dntemplate string `json:"dntemplate,omitempty"`

	// enforce starttls
	// Required: true
	EnforceStarttls *bool `json:"enforce-starttls"`

	// group search
	GroupSearch *LDAPGroupSearchConfig `json:"group-search,omitempty"`

	// host
	// Required: true
	Host *string `json:"host"`

	// lookup dn
	LookupDn string `json:"lookup-dn,omitempty"`

	// lookup password
	LookupPassword string `json:"lookup-password,omitempty"`

	// port
	// Required: true
	Port *int64 `json:"port"`

	// use ldaps
	// Required: true
	UseLdaps *bool `json:"use-ldaps"`

	// user search
	UserSearch *LDAPUserSearchConfig `json:"user-search,omitempty"`
}

// Validate validates this l d a p configuration
func (m *LDAPConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnforceStarttls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseLdaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPConfiguration) validateEnforceStarttls(formats strfmt.Registry) error {

	if err := validate.Required("enforce-starttls", "body", m.EnforceStarttls); err != nil {
		return err
	}

	return nil
}

func (m *LDAPConfiguration) validateGroupSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupSearch) { // not required
		return nil
	}

	if m.GroupSearch != nil {
		if err := m.GroupSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group-search")
			}
			return err
		}
	}

	return nil
}

func (m *LDAPConfiguration) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *LDAPConfiguration) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *LDAPConfiguration) validateUseLdaps(formats strfmt.Registry) error {

	if err := validate.Required("use-ldaps", "body", m.UseLdaps); err != nil {
		return err
	}

	return nil
}

func (m *LDAPConfiguration) validateUserSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.UserSearch) { // not required
		return nil
	}

	if m.UserSearch != nil {
		if err := m.UserSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user-search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPConfiguration) UnmarshalBinary(b []byte) error {
	var res LDAPConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}