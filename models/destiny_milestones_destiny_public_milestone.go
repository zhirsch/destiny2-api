// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyPublicMilestone Information about milestones, presented in a character state-agnostic manner. Combine this data with DestinyMilestoneDefinition to get a full picture of the milestone, which is basically a checklist of things to do in the game. Think of this as GetPublicAdvisors 3.0, for those who used the Destiny 1 API.
// swagger:model Destiny.Milestones.DestinyPublicMilestone

type DestinyMilestonesDestinyPublicMilestone struct {

	// available quests
	AvailableQuests DestinyMilestonesDestinyPublicMilestoneAvailableQuests `json:"availableQuests"`

	// If known, this is the date when the Milestone will expire/recycle/end.
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone.
	MilestoneHash uint32 `json:"milestoneHash,omitempty"`

	// If known, this is the date when the Milestone started/became active.
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant.
	VendorHashes []uint32 `json:"vendorHashes"`
}

/* polymorph Destiny.Milestones.DestinyPublicMilestone availableQuests false */

/* polymorph Destiny.Milestones.DestinyPublicMilestone endDate false */

/* polymorph Destiny.Milestones.DestinyPublicMilestone milestoneHash false */

/* polymorph Destiny.Milestones.DestinyPublicMilestone startDate false */

/* polymorph Destiny.Milestones.DestinyPublicMilestone vendorHashes false */

// Validate validates this destiny milestones destiny public milestone
func (m *DestinyMilestonesDestinyPublicMilestone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVendorHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyMilestonesDestinyPublicMilestone) validateVendorHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.VendorHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestone) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyPublicMilestone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}