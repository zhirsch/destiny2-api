// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition These Art Elements are meant to represent one-off visual effects overlaid on the map. Currently, we do not have a pipeline to import the assets for these overlays, so this info exists as a placeholder for when such a pipeline exists (if it ever will)
// swagger:model Destiny.Definitions.Director.DestinyActivityGraphArtElementDefinition

type DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition struct {

	// The position on the map of the art element.
	Position *DestinyDefinitionsCommonDestinyPositionDefinition `json:"position,omitempty"`
}

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphArtElementDefinition position false */

// Validate validates this destiny definitions director destiny activity graph art element definition
func (m *DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {

		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDirectorDestinyActivityGraphArtElementDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
