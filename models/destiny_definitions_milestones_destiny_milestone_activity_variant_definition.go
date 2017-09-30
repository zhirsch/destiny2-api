// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition Represents a variant on an activity for a Milestone: a specific difficulty tier, or a specific activity variant for example.
// These will often have more specific details, such as an associated Guided Game, progression steps, tier-specific rewards, and custom values.
// swagger:model Destiny.Definitions.Milestones.DestinyMilestoneActivityVariantDefinition

type DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition struct {

	// The hash to use for looking up the variant Activity's definition (DestinyActivityDefinition), where you can find its distinguishing characteristics such as difficulty level and recommended light level.
	// Frequently, that will be the only distinguishing characteristics in practice, which is somewhat of a bummer.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// If you care to do so, render the variants in the order prescribed by this value.
	// When you combine live Milestone data with the definition, the order becomes more useful because you'll be cross-referencing between the definition and live data.
	Order int32 `json:"order,omitempty"`
}

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneActivityVariantDefinition activityHash false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneActivityVariantDefinition order false */

// Validate validates this destiny definitions milestones destiny milestone activity variant definition
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsMilestonesDestinyMilestoneActivityVariantDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}