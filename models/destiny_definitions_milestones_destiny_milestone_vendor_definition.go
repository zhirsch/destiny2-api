// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition If the Milestone or a component has vendors whose inventories could/should be displayed that are relevant to it, this will return the vendor in question.
// It also contains information we need to determine whether that vendor is actually relevant at the moment, given the user's current state.
// swagger:model Destiny.Definitions.Milestones.DestinyMilestoneVendorDefinition

type DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition struct {

	// The hash of the vendor whose wares should be shown as associated with the Milestone.
	VendorHash uint32 `json:"vendorHash,omitempty"`
}

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneVendorDefinition vendorHash false */

// Validate validates this destiny definitions milestones destiny milestone vendor definition
func (m *DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsMilestonesDestinyMilestoneVendorDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}