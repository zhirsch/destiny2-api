// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent dictionary component response ofint32 and destiny item sockets component
// swagger:model DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent

type DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent struct {

	// data
	Data DictionaryComponentResponseOfint32AndDestinyItemSocketsComponentData `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent data false */

/* polymorph DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent privacy false */

// Validate validates this dictionary component response ofint32 and destiny item sockets component
func (m *DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent) Validate(formats strfmt.Registry) error {
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

func (m *DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent) UnmarshalBinary(b []byte) error {
	var res DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}