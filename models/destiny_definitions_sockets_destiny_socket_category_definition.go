// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsSocketsDestinySocketCategoryDefinition Sockets on an item are organized into Categories visually.
// You can find references to the socket category defined on an item's DestinyInventoryItemDefinition.sockets.socketCategories property.
// This has the display information for rendering the categories' header, and a hint for how the UI should handle showing this category.
// swagger:model Destiny.Definitions.Sockets.DestinySocketCategoryDefinition

type DestinyDefinitionsSocketsDestinySocketCategoryDefinition struct {

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`

	// A string hinting to the game's UI system about how the sockets in this category should be displayed.
	// BNet doesn't use it: it's up to you to find valid values and make your own special UI if you want to honor this category style.
	UICategoryStyle uint32 `json:"uiCategoryStyle,omitempty"`
}

/* polymorph Destiny.Definitions.Sockets.DestinySocketCategoryDefinition displayProperties false */

/* polymorph Destiny.Definitions.Sockets.DestinySocketCategoryDefinition hash false */

/* polymorph Destiny.Definitions.Sockets.DestinySocketCategoryDefinition index false */

/* polymorph Destiny.Definitions.Sockets.DestinySocketCategoryDefinition redacted false */

/* polymorph Destiny.Definitions.Sockets.DestinySocketCategoryDefinition uiCategoryStyle false */

// Validate validates this destiny definitions sockets destiny socket category definition
func (m *DestinyDefinitionsSocketsDestinySocketCategoryDefinition) Validate(formats strfmt.Registry) error {
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

func (m *DestinyDefinitionsSocketsDestinySocketCategoryDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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
func (m *DestinyDefinitionsSocketsDestinySocketCategoryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsSocketsDestinySocketCategoryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsSocketsDestinySocketCategoryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}