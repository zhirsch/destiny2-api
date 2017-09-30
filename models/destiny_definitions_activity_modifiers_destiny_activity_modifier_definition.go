// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition Modifiers - in Destiny 1, these were referred to as "Skulls" - are changes that can be applied to an Activity.
// swagger:model Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition

type DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition struct {

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}

/* polymorph Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition displayProperties false */

/* polymorph Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition hash false */

/* polymorph Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition index false */

/* polymorph Destiny.Definitions.ActivityModifiers.DestinyActivityModifierDefinition redacted false */

// Validate validates this destiny definitions activity modifiers destiny activity modifier definition
func (m *DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsActivityModifiersDestinyActivityModifierDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}