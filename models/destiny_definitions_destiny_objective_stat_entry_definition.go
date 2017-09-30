// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyObjectiveStatEntryDefinition Defines the conditions under which stat modifications will be applied to a Character while participating in an objective.
// swagger:model Destiny.Definitions.DestinyObjectiveStatEntryDefinition

type DestinyDefinitionsDestinyObjectiveStatEntryDefinition struct {

	// The stat being modified, and the value used.
	Stat *DestinyDefinitionsDestinyItemInvestmentStatDefinition `json:"stat,omitempty"`

	// Whether it will be applied as long as the objective is active, when it's completed, or until it's completed.
	Style DestinyDestinyObjectiveGrantStyle `json:"style,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyObjectiveStatEntryDefinition stat false */

/* polymorph Destiny.Definitions.DestinyObjectiveStatEntryDefinition style false */

// Validate validates this destiny definitions destiny objective stat entry definition
func (m *DestinyDefinitionsDestinyObjectiveStatEntryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStyle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyObjectiveStatEntryDefinition) validateStat(formats strfmt.Registry) error {

	if swag.IsZero(m.Stat) { // not required
		return nil
	}

	if m.Stat != nil {

		if err := m.Stat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stat")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyDefinitionsDestinyObjectiveStatEntryDefinition) validateStyle(formats strfmt.Registry) error {

	if swag.IsZero(m.Style) { // not required
		return nil
	}

	if err := m.Style.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("style")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyObjectiveStatEntryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyObjectiveStatEntryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyObjectiveStatEntryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}