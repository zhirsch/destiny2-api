// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemSocketCategoryDefinition Sockets are grouped into categories in the UI. These define which category and which sockets are under that category.
// swagger:model Destiny.Definitions.DestinyItemSocketCategoryDefinition

type DestinyDefinitionsDestinyItemSocketCategoryDefinition struct {

	// The hash for the Socket Category: a quick way to go get the header display information for the category. Use it to look up DestinySocketCategoryDefinition info.
	SocketCategoryHash uint32 `json:"socketCategoryHash,omitempty"`

	// Use these indexes to look up the sockets in the "sockets.socketEntries" property on the item definition. These are the indexes under the category, in game-rendered order.
	SocketIndexes []int32 `json:"socketIndexes"`
}

/* polymorph Destiny.Definitions.DestinyItemSocketCategoryDefinition socketCategoryHash false */

/* polymorph Destiny.Definitions.DestinyItemSocketCategoryDefinition socketIndexes false */

// Validate validates this destiny definitions destiny item socket category definition
func (m *DestinyDefinitionsDestinyItemSocketCategoryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSocketIndexes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyItemSocketCategoryDefinition) validateSocketIndexes(formats strfmt.Registry) error {

	if swag.IsZero(m.SocketIndexes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSocketCategoryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSocketCategoryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemSocketCategoryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
