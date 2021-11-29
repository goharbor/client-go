// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OIDCUserInfo OIDC user info
//
// swagger:model OIDCUserInfo
type OIDCUserInfo struct {

	// The creation time of the OIDC user info record.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// the ID of the OIDC info record
	ID int64 `json:"id,omitempty"`

	// the secret of the OIDC user that can be used for CLI to push/pull artifacts
	Secret string `json:"secret,omitempty"`

	// the concatenation of sub and issuer in the ID token
	Subiss string `json:"subiss,omitempty"`

	// The update time of the OIDC user info record.
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// the ID of the user
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this OIDC user info
func (m *OIDCUserInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OIDCUserInfo) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OIDCUserInfo) validateUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OIDCUserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCUserInfo) UnmarshalBinary(b []byte) error {
	var res OIDCUserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
