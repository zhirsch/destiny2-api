// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SingleComponentResponseOfDestinyVendorCategoriesComponent single component response of destiny vendor categories component
// swagger:model SingleComponentResponseOfDestinyVendorCategoriesComponent

type SingleComponentResponseOfDestinyVendorCategoriesComponent struct {

	// data
	Data *DestinyEntitiesVendorsDestinyVendorCategoriesComponent `json:"data,omitempty"`

	// privacy
	Privacy ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

/* polymorph SingleComponentResponseOfDestinyVendorCategoriesComponent data false */

/* polymorph SingleComponentResponseOfDestinyVendorCategoriesComponent privacy false */

// Validate validates this single component response of destiny vendor categories component
func (m *SingleComponentResponseOfDestinyVendorCategoriesComponent) Validate(formats strfmt.Registry) error {
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

func (m *SingleComponentResponseOfDestinyVendorCategoriesComponent) validateData(formats strfmt.Registry) error {

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

func (m *SingleComponentResponseOfDestinyVendorCategoriesComponent) validatePrivacy(formats strfmt.Registry) error {

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
func (m *SingleComponentResponseOfDestinyVendorCategoriesComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleComponentResponseOfDestinyVendorCategoriesComponent) UnmarshalBinary(b []byte) error {
	var res SingleComponentResponseOfDestinyVendorCategoriesComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
