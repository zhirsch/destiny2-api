// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyPublicMilestoneActivity A milestone may have one or more conceptual Activities associated with it, and each of those conceptual activities could have a variety of variants, modes, tiers, what-have-you. Our attempts to determine what qualifies as a conceptual activity are, unfortunately, janky. So if you see missing modes or modes that don't seem appropriate to you, let us know and I'll buy you a beer if we ever meet up in person.
// swagger:model Destiny.Milestones.DestinyPublicMilestoneActivity

type DestinyMilestonesDestinyPublicMilestoneActivity struct {

	// The hash identifier of the activity that's been chosen to be considered the canonical "conceptual" activity definition. This may have many variants, defined herein.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// The activity may have 0-to-many modifiers: if it does, this will contain the hashes to the DestinyActivityModifierDefinition that defines the modifier being applied.
	ModifierHashes []uint32 `json:"modifierHashes"`

	// variants
	Variants DestinyMilestonesDestinyPublicMilestoneActivityVariants `json:"variants"`
}

/* polymorph Destiny.Milestones.DestinyPublicMilestoneActivity activityHash false */

/* polymorph Destiny.Milestones.DestinyPublicMilestoneActivity modifierHashes false */

/* polymorph Destiny.Milestones.DestinyPublicMilestoneActivity variants false */

// Validate validates this destiny milestones destiny public milestone activity
func (m *DestinyMilestonesDestinyPublicMilestoneActivity) Validate(formats strfmt.Registry) error {
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

func (m *DestinyMilestonesDestinyPublicMilestoneActivity) validateModifierHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifierHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestoneActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestoneActivity) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyPublicMilestoneActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}