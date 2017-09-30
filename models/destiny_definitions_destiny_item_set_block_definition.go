// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemSetBlockDefinition Primarily for Quests, this is the definition of properties related to the item if it is a quest and its various quest steps.
// swagger:model Destiny.Definitions.DestinyItemSetBlockDefinition

type DestinyDefinitionsDestinyItemSetBlockDefinition struct {

	// item list
	ItemList DestinyDefinitionsDestinyItemSetBlockDefinitionItemList `json:"itemList"`

	// If true, items in the set can only be added in increasing order, and adding an item will remove any previous item. For Quests, this is by necessity true. Only one quest step is present at a time, and previous steps are removed as you advance in the quest.
	RequireOrderedSetItemAdd bool `json:"requireOrderedSetItemAdd,omitempty"`

	// If true, the UI should treat this quest as "featured"
	SetIsFeatured bool `json:"setIsFeatured,omitempty"`

	// A string identifier we can use to attempt to identify the category of the Quest.
	SetType string `json:"setType,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemSetBlockDefinition itemList false */

/* polymorph Destiny.Definitions.DestinyItemSetBlockDefinition requireOrderedSetItemAdd false */

/* polymorph Destiny.Definitions.DestinyItemSetBlockDefinition setIsFeatured false */

/* polymorph Destiny.Definitions.DestinyItemSetBlockDefinition setType false */

// Validate validates this destiny definitions destiny item set block definition
func (m *DestinyDefinitionsDestinyItemSetBlockDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSetBlockDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSetBlockDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemSetBlockDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
