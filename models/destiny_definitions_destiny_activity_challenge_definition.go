// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyActivityChallengeDefinition Represents a reference to a Challenge, which for now is just an Objective.
// swagger:model Destiny.Definitions.DestinyActivityChallengeDefinition

type DestinyDefinitionsDestinyActivityChallengeDefinition struct {

	// The hash for the Objective that matches this challenge. Use it to look up the DestinyObjectiveDefinition.
	ObjectiveHash uint32 `json:"objectiveHash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyActivityChallengeDefinition objectiveHash false */

// Validate validates this destiny definitions destiny activity challenge definition
func (m *DestinyDefinitionsDestinyActivityChallengeDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityChallengeDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityChallengeDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyActivityChallengeDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}