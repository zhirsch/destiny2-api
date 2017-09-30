// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition This is the position and other data related to nodes in the activity graph that you can click to launch activities. An Activity Graph node will only have one active Activity at a time, which will determine the activity to be launched (and, unless overrideDisplay information is provided, will also determine the tooltip and other UI related to the node)
// swagger:model Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition

type DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition struct {

	// activities
	Activities DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionActivities `json:"activities"`

	// featuring states
	FeaturingStates DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinitionFeaturingStates `json:"featuringStates"`

	// An identifier for the Activity Graph Node, only guaranteed to be unique within its parent Activity Graph.
	NodeID uint32 `json:"nodeId,omitempty"`

	// The node *may* have display properties that override the active Activity's display properties.
	OverrideDisplay *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"overrideDisplay,omitempty"`

	// The position on the map for this node.
	Position *DestinyDefinitionsCommonDestinyPositionDefinition `json:"position,omitempty"`
}

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition activities false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition featuringStates false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition nodeId false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition overrideDisplay false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeDefinition position false */

// Validate validates this destiny definitions director destiny activity graph node definition
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverrideDisplay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition) validateOverrideDisplay(formats strfmt.Registry) error {

	if swag.IsZero(m.OverrideDisplay) { // not required
		return nil
	}

	if m.OverrideDisplay != nil {

		if err := m.OverrideDisplay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrideDisplay")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition) validatePosition(formats strfmt.Registry) error {

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
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDirectorDestinyActivityGraphNodeDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
