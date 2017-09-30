// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition Nodes can have different visual states. This object represents a single visual state ("highlight type") that a node can be in, and the unlock expression condition to determine whether it should be set.
// swagger:model Destiny.Definitions.Director.DestinyActivityGraphNodeFeaturingStateDefinition

type DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition struct {

	// The node can be highlighted in a variety of ways - the game iterates through these and finds the first FeaturingState that is valid at the present moment given the Game, Account, and Character state, and renders the node in that state. See the ActivityGraphNodeHighlightType enum for possible values.
	HighlightType DestinyActivityGraphNodeHighlightType `json:"highlightType,omitempty"`
}

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeFeaturingStateDefinition highlightType false */

// Validate validates this destiny definitions director destiny activity graph node featuring state definition
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHighlightType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition) validateHighlightType(formats strfmt.Registry) error {

	if swag.IsZero(m.HighlightType) { // not required
		return nil
	}

	if err := m.HighlightType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("highlightType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}