// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyActivityModeDefinition destiny definitions destiny activity mode definition
// swagger:model Destiny.Definitions.DestinyActivityModeDefinition

type DestinyDefinitionsDestinyActivityModeDefinition struct {

	// activity mode category
	ActivityModeCategory DestinyDestinyActivityModeCategory `json:"activityModeCategory,omitempty"`

	// activity mode mappings
	ActivityModeMappings map[string]DestinyHistoricalStatsDefinitionsDestinyActivityModeType `json:"activityModeMappings,omitempty"`

	// If FALSE, we want to ignore this type when we're showing activity modes in BNet UI. It will still be returned in case 3rd parties want to use it for any purpose.
	Display bool `json:"display,omitempty"`

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// friendly name
	FriendlyName string `json:"friendlyName,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// is aggregate mode
	IsAggregateMode bool `json:"isAggregateMode,omitempty"`

	// is team based
	IsTeamBased bool `json:"isTeamBased,omitempty"`

	// mode type
	ModeType DestinyHistoricalStatsDefinitionsDestinyActivityModeType `json:"modeType,omitempty"`

	// The relative ordering of activity modes.
	Order int32 `json:"order,omitempty"`

	// parent hashes
	ParentHashes []uint32 `json:"parentHashes"`

	// pgcr image
	PgcrImage string `json:"pgcrImage,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition activityModeCategory false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition activityModeMappings false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition display false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition friendlyName false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition hash false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition index false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition isAggregateMode false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition isTeamBased false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition modeType false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition order false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition parentHashes false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition pgcrImage false */

/* polymorph Destiny.Definitions.DestinyActivityModeDefinition redacted false */

// Validate validates this destiny definitions destiny activity mode definition
func (m *DestinyDefinitionsDestinyActivityModeDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityModeCategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParentHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyActivityModeDefinition) validateActivityModeCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityModeCategory) { // not required
		return nil
	}

	if err := m.ActivityModeCategory.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("activityModeCategory")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyActivityModeDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

func (m *DestinyDefinitionsDestinyActivityModeDefinition) validateModeType(formats strfmt.Registry) error {

	if swag.IsZero(m.ModeType) { // not required
		return nil
	}

	if err := m.ModeType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("modeType")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyActivityModeDefinition) validateParentHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityModeDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityModeDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyActivityModeDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
