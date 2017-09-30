// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SingleComponentResponseOfDestinyInventoryComponent single component response of destiny inventory component
// swagger:model SingleComponentResponseOfDestinyInventoryComponent

type SingleComponentResponseOfDestinyInventoryComponent struct {

	// data
	Data *DestinyEntitiesInventoryDestinyInventoryComponent `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph SingleComponentResponseOfDestinyInventoryComponent data false */

/* polymorph SingleComponentResponseOfDestinyInventoryComponent privacy false */

// Validate validates this single component response of destiny inventory component
func (m *SingleComponentResponseOfDestinyInventoryComponent) Validate(formats strfmt.Registry) error {
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

func (m *SingleComponentResponseOfDestinyInventoryComponent) validateData(formats strfmt.Registry) error {

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

func (m *SingleComponentResponseOfDestinyInventoryComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *SingleComponentResponseOfDestinyInventoryComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleComponentResponseOfDestinyInventoryComponent) UnmarshalBinary(b []byte) error {
	var res SingleComponentResponseOfDestinyInventoryComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
