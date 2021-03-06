// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent dictionary component response ofint32 and destiny item talent grid component
// swagger:model DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent

type DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent struct {

	// data
	Data DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponentData `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent data false */

/* polymorph DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent privacy false */

// Validate validates this dictionary component response ofint32 and destiny item talent grid component
func (m *DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent) Validate(formats strfmt.Registry) error {
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

func (m *DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent) UnmarshalBinary(b []byte) error {
	var res DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
