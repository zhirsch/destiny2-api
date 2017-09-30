// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition Guided Game information for this activity.
// swagger:model Destiny.Definitions.DestinyActivityGuidedBlockDefinition

type DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition struct {

	// If -1, the guided group cannot be disbanded. Otherwise, take the total # of players in the activity and subtract this number: that is the total # of votes needed for the guided group to disband.
	GUIDEDDisbandCount int32 `json:"guidedDisbandCount,omitempty"`

	// The maximum amount of people that can be in the waiting lobby.
	GUIDEDMaxLobbySize int32 `json:"guidedMaxLobbySize,omitempty"`

	// The minimum amount of people that can be in the waiting lobby.
	GUIDEDMinLobbySize int32 `json:"guidedMinLobbySize,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyActivityGuidedBlockDefinition guidedDisbandCount false */

/* polymorph Destiny.Definitions.DestinyActivityGuidedBlockDefinition guidedMaxLobbySize false */

/* polymorph Destiny.Definitions.DestinyActivityGuidedBlockDefinition guidedMinLobbySize false */

// Validate validates this destiny definitions destiny activity guided block definition
func (m *DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyActivityGUIDEDBlockDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}