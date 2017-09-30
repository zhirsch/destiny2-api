// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyActivityGraphListEntryDefinition Destinations and Activities may have default Activity Graphs that should be shown when you bring up the Director and are playing in either.
// This contract defines the graph referred to and the gating for when it is relevant.
// swagger:model Destiny.Definitions.DestinyActivityGraphListEntryDefinition

type DestinyDefinitionsDestinyActivityGraphListEntryDefinition struct {

	// The hash identifier of the DestinyActivityGraphDefinition that should be shown when opening the director.
	ActivityGraphHash uint32 `json:"activityGraphHash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyActivityGraphListEntryDefinition activityGraphHash false */

// Validate validates this destiny definitions destiny activity graph list entry definition
func (m *DestinyDefinitionsDestinyActivityGraphListEntryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityGraphListEntryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyActivityGraphListEntryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyActivityGraphListEntryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
