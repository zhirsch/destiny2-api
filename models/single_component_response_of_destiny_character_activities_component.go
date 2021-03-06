// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SingleComponentResponseOfDestinyCharacterActivitiesComponent single component response of destiny character activities component
// swagger:model SingleComponentResponseOfDestinyCharacterActivitiesComponent

type SingleComponentResponseOfDestinyCharacterActivitiesComponent struct {

	// data
	Data *DestinyEntitiesCharactersDestinyCharacterActivitiesComponent `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph SingleComponentResponseOfDestinyCharacterActivitiesComponent data false */

/* polymorph SingleComponentResponseOfDestinyCharacterActivitiesComponent privacy false */

// Validate validates this single component response of destiny character activities component
func (m *SingleComponentResponseOfDestinyCharacterActivitiesComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrivacy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SingleComponentResponseOfDestinyCharacterActivitiesComponent) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *SingleComponentResponseOfDestinyCharacterActivitiesComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *SingleComponentResponseOfDestinyCharacterActivitiesComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleComponentResponseOfDestinyCharacterActivitiesComponent) UnmarshalBinary(b []byte) error {
	var res SingleComponentResponseOfDestinyCharacterActivitiesComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
