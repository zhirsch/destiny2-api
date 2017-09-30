// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent dictionary component response ofuint32 and destiny item plug component
// swagger:model DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent

type DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent struct {

	// data
	Data DictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent data false */

/* polymorph DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent privacy false */

// Validate validates this dictionary component response ofuint32 and destiny item plug component
func (m *DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivacy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent) validatePrivacy(formats strfmt.Registry) error {

	if swag.IsZero(m.Privacy) { // not required
		return nil
	}

	if err := m.Privacy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("privacy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent) UnmarshalBinary(b []byte) error {
	var res DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
