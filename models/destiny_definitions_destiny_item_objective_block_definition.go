// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemObjectiveBlockDefinition An item can have objectives on it. In practice, these are the exclusive purview of "Quest Step" items: DestinyInventoryItemDefinitions that represent a specific step in a Quest.
// Quest steps have 1:M objectives that we end up processing and returning in live data as DestinyQuestStatus data, and other useful information.
// swagger:model Destiny.Definitions.DestinyItemObjectiveBlockDefinition

type DestinyDefinitionsDestinyItemObjectiveBlockDefinition struct {

	// For every entry in objectiveHashes, there is a corresponding entry in this array at the same index. If the objective is meant to be associated with a specific DestinyActivityDefinition, there will be a valid hash at that index. Otherwise, it will be invalid (0).
	DisplayActivityHashes []uint32 `json:"displayActivityHashes"`

	// The localized string for narrative text related to this quest step, if any.
	Narrative string `json:"narrative,omitempty"`

	// The hashes to Objectives (DestinyObjectiveDefinition) that are part of this Quest Step, in the order that they should be rendered.
	ObjectiveHashes []uint32 `json:"objectiveHashes"`

	// The localized string describing an action to be performed associated with the objectives, if any.
	ObjectiveVerbName string `json:"objectiveVerbName,omitempty"`

	// A hashed value for the questTypeIdentifier, because apparently I like to be redundant.
	QuestTypeHash uint32 `json:"questTypeHash,omitempty"`

	// The identifier for the type of quest being performed, if any. Not associated with any fixed definition, yet.
	QuestTypeIdentifier string `json:"questTypeIdentifier,omitempty"`

	// The hash for the DestinyInventoryItemDefinition representing the Quest to which this Quest Step belongs.
	QuestlineItemHash uint32 `json:"questlineItemHash,omitempty"`

	// If True, all objectives must be completed for the step to be completed. If False, any one objective can be completed for the step to be completed.
	RequireFullObjectiveCompletion bool `json:"requireFullObjectiveCompletion,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition displayActivityHashes false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition narrative false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition objectiveHashes false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition objectiveVerbName false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition questTypeHash false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition questTypeIdentifier false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition questlineItemHash false */

/* polymorph Destiny.Definitions.DestinyItemObjectiveBlockDefinition requireFullObjectiveCompletion false */

// Validate validates this destiny definitions destiny item objective block definition
func (m *DestinyDefinitionsDestinyItemObjectiveBlockDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayActivityHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateObjectiveHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyItemObjectiveBlockDefinition) validateDisplayActivityHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayActivityHashes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyItemObjectiveBlockDefinition) validateObjectiveHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectiveHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemObjectiveBlockDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemObjectiveBlockDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemObjectiveBlockDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}