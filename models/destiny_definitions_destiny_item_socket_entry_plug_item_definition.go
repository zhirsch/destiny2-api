// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition The definition of a known, reusable plug that can be applied to a socket.
// swagger:model Destiny.Definitions.DestinyItemSocketEntryPlugItemDefinition

type DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition struct {

	// The hash identifier of a DestinyInventoryItemDefinition representing the plug that can be inserted.
	PlugItemHash uint32 `json:"plugItemHash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemSocketEntryPlugItemDefinition plugItemHash false */

// Validate validates this destiny definitions destiny item socket entry plug item definition
func (m *DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}