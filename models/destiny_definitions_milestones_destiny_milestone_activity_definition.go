// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition Milestones can have associated activities which provide additional information about the context, challenges, modifiers, state etc... related to this Milestone.
// Information we need to be able to return that data is defined here, along with Tier data to establish a relationship between a conceptual Activity and its difficulty levels and variants.
// swagger:model Destiny.Definitions.Milestones.DestinyMilestoneActivityDefinition

type DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition struct {

	// The "Conceptual" activity hash. Basically, we picked the lowest level activity and are treating it as the canonical definition of the activity for rendering purposes.
	// If you care about the specific difficulty modes and variations, use the activities under "Variants".
	ConceptualActivityHash uint32 `json:"conceptualActivityHash,omitempty"`

	// variants
	Variants DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinitionVariants `json:"variants,omitempty"`
}

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneActivityDefinition conceptualActivityHash false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneActivityDefinition variants false */

// Validate validates this destiny definitions milestones destiny milestone activity definition
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}