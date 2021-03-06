// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneQuestRewardsDefinitionItems The items that represent your reward for completing the quest.
// Be warned, these could be "dummy" items: items that are only used to render a good-looking in-game tooltip, but aren't the actual items themselves.
// For instance, when experience is given there's often a dummy item representing "experience", with quantity being the amount of experience you got. We don't have a programmatic association between those and whatever Progression is actually getting that experience... yet.
// swagger:model destinyDefinitionsMilestonesDestinyMilestoneQuestRewardsDefinitionItems

type DestinyDefinitionsMilestonesDestinyMilestoneQuestRewardsDefinitionItems []*DestinyDefinitionsMilestonesDestinyMilestoneQuestRewardItem

// Validate validates this destiny definitions milestones destiny milestone quest rewards definition items
func (m DestinyDefinitionsMilestonesDestinyMilestoneQuestRewardsDefinitionItems) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
