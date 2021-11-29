// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionRuleTrigger retention rule trigger
//
// swagger:model RetentionRuleTrigger
type RetentionRuleTrigger struct {

	// kind
	Kind string `json:"kind,omitempty"`

	// references
	References interface{} `json:"references,omitempty"`

	// settings
	Settings interface{} `json:"settings,omitempty"`
}

// Validate validates this retention rule trigger
func (m *RetentionRuleTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionRuleTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionRuleTrigger) UnmarshalBinary(b []byte) error {
	var res RetentionRuleTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
