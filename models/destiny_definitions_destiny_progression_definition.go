// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyProgressionDefinition A "Progression" in Destiny is best explained by an example.
// A Character's "Level" is a progression: it has Experience that can be earned, levels that can be gained, and is evaluated and displayed at various points in the game. A Character's "Faction Reputation" is also a progression for much the same reason.
// Progression is used by a variety of systems, and the definition of a Progression will generally only be useful if combining with live data (such as a character's DestinyCharacterProgressionComponent.progressions property, which holds that character's live Progression states).
// Fundamentally, a Progression measures your "Level" by evaluating the thresholds in its Steps (one step per level, except for the last step which can be repeated indefinitely for "Levels" that have no ceiling) against the total earned "progression points"/experience. (for simplicity purposes, we will henceforth refer to earned progression points as experience, though it need not be a mechanic that in any way resembles Experience in a traditional sense).
// Earned experience is calculated in a variety of ways, determined by the Progression's scope. These go from looking up a stored value to performing exceedingly obtuse calculations. This is why we provide live data in DestinyCharacterProgressionComponent.progressions, so you don't have to worry about those.
// swagger:model Destiny.Definitions.DestinyProgressionDefinition

type DestinyDefinitionsDestinyProgressionDefinition struct {

	// display properties
	DisplayProperties *DestinyDefinitionsDestinyProgressionDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// If the value exists, this is the hash identifier for the Faction that owns this Progression.
	// This is purely for convenience, if you're looking at a progression and want to know if and who it's related to in terms of Faction Reputation.
	FactionHash uint32 `json:"factionHash,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`

	// If this is True, then the progression doesn't have a maximum level.
	RepeatLastStep bool `json:"repeatLastStep,omitempty"`

	// The "Scope" of the progression indicates the source of the progression's live data.
	// See the DestinyProgressionScope enum for more info: but essentially, a Progression can either be backed by a stored value, or it can be a calculated derivative of other values.
	Scope DestinyDestinyProgressionScope `json:"scope,omitempty"`

	// If there's a description of how to earn this progression in the local config, this will be that localized description.
	Source string `json:"source,omitempty"`

	// steps
	Steps DestinyDefinitionsDestinyProgressionDefinitionSteps `json:"steps"`

	// If true, the Progression is something worth showing to users.
	// If false, BNet isn't going to show it. But that doesn't mean you can't. We're all friends here.
	Visible bool `json:"visible,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyProgressionDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition factionHash false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition hash false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition index false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition redacted false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition repeatLastStep false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition scope false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition source false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition steps false */

/* polymorph Destiny.Definitions.DestinyProgressionDefinition visible false */

// Validate validates this destiny definitions destiny progression definition
func (m *DestinyDefinitionsDestinyProgressionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyProgressionDefinition) validateDisplayProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayProperties) { // not required
		return nil
	}

	if m.DisplayProperties != nil {

		if err := m.DisplayProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayProperties")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyDefinitionsDestinyProgressionDefinition) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyProgressionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyProgressionDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyProgressionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
