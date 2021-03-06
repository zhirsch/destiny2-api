// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DestinyDefinitionsDestinyActivityPlaylistItemDefinition If the activity is a playlist, this is the definition for a specific entry in the playlist: a single possible combination of Activity and Activity Mode that can be chosen.
// swagger:model Destiny.Definitions.DestinyActivityPlaylistItemDefinition

type DestinyDefinitionsDestinyActivityPlaylistItemDefinition struct {

	// The hash identifier of the Activity that can be played. Use it to look up the DestinyActivityDefinition.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// activity mode hashes
	ActivityModeHashes []uint32 `json:"activityModeHashes"`

	// activity mode types
	ActivityModeTypes []DestinyHistoricalStatsDefinitionsDestinyActivityModeType `json:"activityModeTypes"`

	// direct activity mode hash
	DirectActivityModeHash uint32 `json:"directActivityModeHash,omitempty"`

	// direct activity mode type
	DirectActivityModeType int32 `json:"directActivityModeType,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyActivityPlaylistItemDefinition activityHash false */

/* polymorph Destiny.Definitions.DestinyActivityPlaylistItemDefinition activityModeHashes false */

/* polymorph Destiny.Definitions.DestinyActivityPlaylistItemDefinition activityModeTypes false */

/* polymorph Destiny.Definitions.DestinyActivityPlaylistItemDefinition directActivityModeHash false */

/* polymorph Destiny.Definitions.DestinyActivityPlaylistItemDefinition directActivityModeType false */

// Validate validates this destiny definitions destiny activity playlist item definition
func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityModeHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateActivityModeTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDirectActivityModeType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) validateActivityModeHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityModeHashes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) validateActivityModeTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityModeTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivityModeTypes); i++ {

		if err := m.ActivityModeTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activityModeTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

var destinyDefinitionsDestinyActivityPlaylistItemDefinitionTypeDirectActivityModeTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,2,3,4,5,6,7,9,10,11,12,13,15,16,17,18,19,20,21,22,24,25,26,27,28,29,30,31,32,37,38,39,40]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDefinitionsDestinyActivityPlaylistItemDefinitionTypeDirectActivityModeTypePropEnum = append(destinyDefinitionsDestinyActivityPlaylistItemDefinitionTypeDirectActivityModeTypePropEnum, v)
	}
}

// prop value enum
func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) validateDirectActivityModeTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, destinyDefinitionsDestinyActivityPlaylistItemDefinitionTypeDirectActivityModeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) validateDirectActivityModeType(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectActivityModeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectActivityModeTypeEnum("directActivityModeType", "body", m.DirectActivityModeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityPlaylistItemDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyActivityPlaylistItemDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
