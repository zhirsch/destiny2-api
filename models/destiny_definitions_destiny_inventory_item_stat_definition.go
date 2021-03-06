// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyInventoryItemStatDefinition Defines a specific stat value on an item, and the minimum/maximum range that we could compute for the item based on our heuristics for how the item might be generated.
// Not guaranteed to match real-world instances of the item, but should hopefully at least be close. If it's not close, let us know on the Bungie API forums.
// swagger:model Destiny.Definitions.DestinyInventoryItemStatDefinition

type DestinyDefinitionsDestinyInventoryItemStatDefinition struct {

	// The maximum possible value for this stat that we think the item can roll.
	Maximum int32 `json:"maximum,omitempty"`

	// The minimum possible value for this stat that we think the item can roll.
	Minimum int32 `json:"minimum,omitempty"`

	// The hash for the DestinyStatDefinition representing this stat.
	StatHash uint32 `json:"statHash,omitempty"`

	// This value represents the stat value assuming the minimum possible roll but accounting for any mandatory bonuses that should be applied to the stat on item creation.
	// In Destiny 1, this was different from the "minimum" value because there were certain conditions where an item could be theoretically lower level/value than the initial roll.
	// In Destiny 2, this is not possible unless Talent Grids begin to be used again for these purposes or some other system change occurs... thus in practice, value and minimum should be the same in Destiny 2. Good riddance.
	Value int32 `json:"value,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyInventoryItemStatDefinition maximum false */

/* polymorph Destiny.Definitions.DestinyInventoryItemStatDefinition minimum false */

/* polymorph Destiny.Definitions.DestinyInventoryItemStatDefinition statHash false */

/* polymorph Destiny.Definitions.DestinyInventoryItemStatDefinition value false */

// Validate validates this destiny definitions destiny inventory item stat definition
func (m *DestinyDefinitionsDestinyInventoryItemStatDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyInventoryItemStatDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyInventoryItemStatDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyInventoryItemStatDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
