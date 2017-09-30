// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsMilestonesDestinyMilestoneDefinition Milestones are an in-game concept where they're attempting to tell you what you can do next in-game.
// If that sounds a lot like Advisors in Destiny 1, it is! So we threw out Advisors in the Destiny 2 API and tacked all of the data we would have put on Advisors onto Milestones instead.
// Each Milestone represents something going on in the game right now:
// - A "ritual activity" you can perform, like nightfall
// - A "special event" that may have activities related to it, like Taco Tuesday (there's no Taco Tuesday in Destiny 2)
// - A checklist you can fulfill, like helping your Clan complete all of its weekly objectives
// - A tutorial quest you can play through, like the introduction to the Crucible.
// Most of these milestones appear in game as well. Some of them are BNet only, because we're so extra. You're welcome.
// swagger:model Destiny.Definitions.Milestones.DestinyMilestoneDefinition

type DestinyDefinitionsMilestonesDestinyMilestoneDefinition struct {

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// If the milestone has a friendly identifier for association with other features - such as Recruiting - that identifier can be found here. This is "friendly" in that it looks better in a URL than whatever the identifier for the Milestone actually is.
	FriendlyName string `json:"friendlyName,omitempty"`

	// A shortcut for clients - and the server - to understand whether we can predict the start and end dates for this event. In practice, there are multiple ways that an event could have predictable date ranges, but not all events will be able to be predicted via any mechanism (for instance, events that are manually triggered on and off)
	HasPredictableDates bool `json:"hasPredictableDates,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// A custom image someone made just for the milestone. Isn't that special?
	Image string `json:"image,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// Some milestones are explicit objectives that you can see and interact with in the game. Some milestones are more conceptual, built by BNet to help advise you on activities and events that happen in-game but that aren't explicitly shown in game as Milestones. If this is TRUE, you can see this as a milestone in the game. If this is FALSE, it's an event or activity you can participate in, but you won't see it as a Milestone in the game's UI.
	IsInGameMilestone bool `json:"isInGameMilestone,omitempty"`

	// An enumeration listing one of the possible types of milestones. Check out the DestinyMilestoneType enum for more info!
	MilestoneType DestinyDefinitionsMilestonesDestinyMilestoneType `json:"milestoneType,omitempty"`

	// quests
	Quests DestinyDefinitionsMilestonesDestinyMilestoneDefinitionQuests `json:"quests,omitempty"`

	// If True, then the Milestone has been integrated with BNet's recruiting feature.
	Recruitable bool `json:"recruitable,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`

	// rewards
	Rewards DestinyDefinitionsMilestonesDestinyMilestoneDefinitionRewards `json:"rewards,omitempty"`

	// If TRUE, this entry should be returned in the list of milestones for the "Explore Destiny" (i.e. new BNet homepage) features of Bungie.net (as long as the underlying event is active)
	ShowInExplorer bool `json:"showInExplorer,omitempty"`

	// values
	Values DestinyDefinitionsMilestonesDestinyMilestoneDefinitionValues `json:"values,omitempty"`

	// vendors
	Vendors DestinyDefinitionsMilestonesDestinyMilestoneDefinitionVendors `json:"vendors"`
}

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition displayProperties false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition friendlyName false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition hasPredictableDates false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition hash false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition image false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition index false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition isInGameMilestone false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition milestoneType false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition quests false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition recruitable false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition redacted false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition rewards false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition showInExplorer false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition values false */

/* polymorph Destiny.Definitions.Milestones.DestinyMilestoneDefinition vendors false */

// Validate validates this destiny definitions milestones destiny milestone definition
func (m *DestinyDefinitionsMilestonesDestinyMilestoneDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMilestoneType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsMilestonesDestinyMilestoneDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

func (m *DestinyDefinitionsMilestonesDestinyMilestoneDefinition) validateMilestoneType(formats strfmt.Registry) error {

	if swag.IsZero(m.MilestoneType) { // not required
		return nil
	}

	if err := m.MilestoneType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("milestoneType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsMilestonesDestinyMilestoneDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsMilestonesDestinyMilestoneDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
