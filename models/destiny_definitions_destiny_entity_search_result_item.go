// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyEntitySearchResultItem An individual Destiny Entity returned from the entity search.
// swagger:model Destiny.Definitions.DestinyEntitySearchResultItem

type DestinyDefinitionsDestinyEntitySearchResultItem struct {

	// Basic display properties on the entity, so you don't have to look up the definition to show basic results for the item.
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// The type of entity, returned as a string matching the DestinyDefinition's contract class name. You'll have to have your own mapping from class names to actually looking up those definitions in the manifest databases.
	EntityType string `json:"entityType,omitempty"`

	// The hash identifier of the entity. You will use this to look up the DestinyDefinition relevant for the entity found.
	Hash uint32 `json:"hash,omitempty"`

	// The ranking value for sorting that we calculated using our relevance formula. This will hopefully get better with time and iteration.
	Weight float64 `json:"weight,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyEntitySearchResultItem displayProperties false */

/* polymorph Destiny.Definitions.DestinyEntitySearchResultItem entityType false */

/* polymorph Destiny.Definitions.DestinyEntitySearchResultItem hash false */

/* polymorph Destiny.Definitions.DestinyEntitySearchResultItem weight false */

// Validate validates this destiny definitions destiny entity search result item
func (m *DestinyDefinitionsDestinyEntitySearchResultItem) Validate(formats strfmt.Registry) error {
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

func (m *DestinyDefinitionsDestinyEntitySearchResultItem) validateDisplayProperties(formats strfmt.Registry) error {

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
func (m *DestinyDefinitionsDestinyEntitySearchResultItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyEntitySearchResultItem) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyEntitySearchResultItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
