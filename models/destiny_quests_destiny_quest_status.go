// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyQuestsDestinyQuestStatus Data regarding the progress of a Quest for a specific character. Quests are composed of multiple steps, each with potentially multiple objectives: this QuestStatus will return Objective data for the *currently active* step in this quest.
// swagger:model Destiny.Quests.DestinyQuestStatus

type DestinyQuestsDestinyQuestStatus struct {

	// Whether or not the whole quest has been completed, regardless of whether or not you have redeemed the rewards for the quest.
	Completed bool `json:"completed,omitempty"`

	// The current Quest Step will be an instanced item in the player's inventory. If you care about that, this is the instance ID of that item.
	ItemInstanceID int64 `json:"itemInstanceId,omitempty"`

	// The hash identifier for the Quest Item. (Note: Quests are defined as Items, and thus you would use this to look up the quest's DestinyInventoryItemDefinition). For information on all steps in the quest, you can then examine its DestinyInventoryItemDefinition.setData property for Quest Steps (which are *also* items). You can use the Item Definition to display human readable data about the overall quest.
	QuestHash uint32 `json:"questHash,omitempty"`

	// Whether or not you have redeemed rewards for this quest.
	Redeemed bool `json:"redeemed,omitempty"`

	// Whether or not you have started this quest.
	Started bool `json:"started,omitempty"`

	// The hash identifier of the current Quest Step, which is also a DestinyInventoryItemDefinition. You can use this to get human readable data about the current step and what to do in that step.
	StepHash uint32 `json:"stepHash,omitempty"`

	// step objectives
	StepObjectives DestinyQuestsDestinyQuestStatusStepObjectives `json:"stepObjectives"`

	// Whether or not the quest is tracked
	Tracked bool `json:"tracked,omitempty"`

	// If the quest has a related Vendor that you should talk to in order to initiate the quest/earn rewards/continue the quest, this will be the hash identifier of that Vendor. Look it up its DestinyVendorDefinition.
	VendorHash uint32 `json:"vendorHash,omitempty"`
}

/* polymorph Destiny.Quests.DestinyQuestStatus completed false */

/* polymorph Destiny.Quests.DestinyQuestStatus itemInstanceId false */

/* polymorph Destiny.Quests.DestinyQuestStatus questHash false */

/* polymorph Destiny.Quests.DestinyQuestStatus redeemed false */

/* polymorph Destiny.Quests.DestinyQuestStatus started false */

/* polymorph Destiny.Quests.DestinyQuestStatus stepHash false */

/* polymorph Destiny.Quests.DestinyQuestStatus stepObjectives false */

/* polymorph Destiny.Quests.DestinyQuestStatus tracked false */

/* polymorph Destiny.Quests.DestinyQuestStatus vendorHash false */

// Validate validates this destiny quests destiny quest status
func (m *DestinyQuestsDestinyQuestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyQuestsDestinyQuestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyQuestsDestinyQuestStatus) UnmarshalBinary(b []byte) error {
	var res DestinyQuestsDestinyQuestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
