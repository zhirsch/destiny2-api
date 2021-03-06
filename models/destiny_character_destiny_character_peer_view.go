// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyCharacterDestinyCharacterPeerView A minimal view of a character's equipped items, for the purpose of rendering a summary screen or showing the character in 3D.
// swagger:model Destiny.Character.DestinyCharacterPeerView

type DestinyCharacterDestinyCharacterPeerView struct {

	// equipment
	Equipment DestinyCharacterDestinyCharacterPeerViewEquipment `json:"equipment"`
}

/* polymorph Destiny.Character.DestinyCharacterPeerView equipment false */

// Validate validates this destiny character destiny character peer view
func (m *DestinyCharacterDestinyCharacterPeerView) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyCharacterDestinyCharacterPeerView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyCharacterDestinyCharacterPeerView) UnmarshalBinary(b []byte) error {
	var res DestinyCharacterDestinyCharacterPeerView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
