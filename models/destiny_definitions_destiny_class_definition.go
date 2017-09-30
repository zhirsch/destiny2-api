// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyClassDefinition Defines a Character Class in Destiny 2. These are types of characters you can play, like Titan, Warlock, and Hunter.
// swagger:model Destiny.Definitions.DestinyClassDefinition

type DestinyDefinitionsDestinyClassDefinition struct {

	// In Destiny 1, we added a convenience Enumeration for referring to classes. We've kept it, though mostly for posterity. This is the enum value for this definition's class.
	ClassType DestinyDestinyClass `json:"classType,omitempty"`

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// A localized string referring to the singular form of the Class's name when referred to in gendered form. Keyed by the DestinyGender.
	GenderedClassNames map[string]string `json:"genderedClassNames,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// If the Class has a Mentor (all classes *should*), this will be the hash identifier for that Vendor if you care.
	MentorVendorHash uint32 `json:"mentorVendorHash,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyClassDefinition classType false */

/* polymorph Destiny.Definitions.DestinyClassDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyClassDefinition genderedClassNames false */

/* polymorph Destiny.Definitions.DestinyClassDefinition hash false */

/* polymorph Destiny.Definitions.DestinyClassDefinition index false */

/* polymorph Destiny.Definitions.DestinyClassDefinition mentorVendorHash false */

/* polymorph Destiny.Definitions.DestinyClassDefinition redacted false */

// Validate validates this destiny definitions destiny class definition
func (m *DestinyDefinitionsDestinyClassDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyClassDefinition) validateClassType(formats strfmt.Registry) error {

	if swag.IsZero(m.ClassType) { // not required
		return nil
	}

	if err := m.ClassType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("classType")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyClassDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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
func (m *DestinyDefinitionsDestinyClassDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyClassDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyClassDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
