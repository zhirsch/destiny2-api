// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition destiny definitions destiny vendor display properties definition
// swagger:model Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition

type DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition struct {

	// description
	Description string `json:"description,omitempty"`

	// has icon
	HasIcon bool `json:"hasIcon,omitempty"`

	// Note that "icon" is sometimes misleading, and should be interpreted in the context of the entity. For instance, in Destiny 1 the DestinyRecordBookDefinition's icon was a big picture of a book.
	// But usually, it will be a small square image that you can use as... well, an icon.
	Icon string `json:"icon,omitempty"`

	// large icon
	LargeIcon string `json:"largeIcon,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// requirements display
	RequirementsDisplay DestinyDefinitionsDestinyVendorDisplayPropertiesDefinitionRequirementsDisplay `json:"requirementsDisplay"`

	// subtitle
	Subtitle string `json:"subtitle,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition description false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition hasIcon false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition icon false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition largeIcon false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition name false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition requirementsDisplay false */

/* polymorph Destiny.Definitions.DestinyVendorDisplayPropertiesDefinition subtitle false */

// Validate validates this destiny definitions destiny vendor display properties definition
func (m *DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyVendorDisplayPropertiesDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
