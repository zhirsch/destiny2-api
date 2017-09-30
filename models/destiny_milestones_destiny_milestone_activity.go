// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyMilestoneActivity Sometimes, we know the specific activity that the Milestone wants you to play. This entity provides additional information about that Activity and all of its variants. (sometimes there's only one variant, but I think you get the point)
// swagger:model Destiny.Milestones.DestinyMilestoneActivity

type DestinyMilestonesDestinyMilestoneActivity struct {

	// The hash of an arbitrarily chosen variant of this activity. We'll go ahead and call that the "canonical" activity, because if you're using this value you should only use it for properties that are common across the variants: things like the name of the activity, it's location, etc... Use this hash to look up the DestinyActivityDefinition of this activity for rendering data.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// If the activity has modifiers, this will be the list of modifiers that all variants have in common. Perform lookups against DestinyActivityModifierDefinition which defines the modifier being applied to get at the modifier data. Note that, in the DestiyActivityDefinition, you will see many more modifiers than this being referred to: those are all *possible* modifiers for the activity, not the active ones. Use only the active ones to match what's really live.
	ModifierHashes []uint32 `json:"modifierHashes"`

	// variants
	Variants DestinyMilestonesDestinyMilestoneActivityVariants `json:"variants"`
}

/* polymorph Destiny.Milestones.DestinyMilestoneActivity activityHash false */

/* polymorph Destiny.Milestones.DestinyMilestoneActivity modifierHashes false */

/* polymorph Destiny.Milestones.DestinyMilestoneActivity variants false */

// Validate validates this destiny milestones destiny milestone activity
func (m *DestinyMilestonesDestinyMilestoneActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifierHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyMilestonesDestinyMilestoneActivity) validateModifierHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifierHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneActivity) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyMilestoneActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}