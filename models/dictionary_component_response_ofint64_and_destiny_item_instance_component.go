// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent dictionary component response ofint64 and destiny item instance component
// swagger:model DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent

type DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent struct {

	// data
	Data DictionaryComponentResponseOfint64AndDestinyItemInstanceComponentData `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent data false */

/* polymorph DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent privacy false */

// Validate validates this dictionary component response ofint64 and destiny item instance component
func (m *DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent) Validate(formats strfmt.Registry) error {
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

func (m *DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent) UnmarshalBinary(b []byte) error {
	var res DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
