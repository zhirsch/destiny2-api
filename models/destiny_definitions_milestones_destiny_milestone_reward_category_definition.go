// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition The definition of a category of rewards, that contains many individual rewards.
// swagger:model Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition

type DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition struct {

	// Identifies the reward category. Only guaranteed unique within this specific component!
	CategoryHash uint32 `json:"categoryHash,omitempty"`

	// The string identifier for the category, if you want to use it for some end. Guaranteed unique within the specific component.
	CategoryIdentifier string `json:"categoryIdentifier,omitempty"`

	// Hopefully this is obvious by now.
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// If you want to use BNet's recommended order for rendering categories programmatically, use this value and compare it to other categories to determine the order in which they should be rendered. I don't feel great about putting this here, I won't lie.
	Order int32 `json:"order,omitempty"`

	// reward entries
	RewardEntries DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinitionRewardEntries `json:"rewardEntries,omitempty"`
}

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition categoryHash false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition categoryIdentifier false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition displayProperties false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition order false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneRewardCategoryDefinition rewardEntries false */

// Validate validates this destiny definitions milestones destiny milestone reward category definition
func (m *DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsMilestonesDestinyMilestoneRewardCategoryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}