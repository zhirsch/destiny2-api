// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyLocationDefinition A "Location" is a sort of shortcut for referring to a specific combination of Activity, Destination, Place, and even Bubble or NavPoint within a space.
// Most of this data isn't intrinsically useful to us, but Objectives refer to locations, and through that we can at least infer the Activity, Destination, and Place being referred to by the Objective.
// swagger:model Destiny.Definitions.DestinyLocationDefinition

type DestinyDefinitionsDestinyLocationDefinition struct {

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// location releases
	LocationReleases DestinyDefinitionsDestinyLocationDefinitionLocationReleases `json:"locationReleases"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`

	// If the location has a Vendor on it, this is the hash identifier for that Vendor. Look them up with DestinyVendorDefinition.
	VendorHash uint32 `json:"vendorHash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyLocationDefinition hash false */

/* polymorph Destiny.Definitions.DestinyLocationDefinition index false */

/* polymorph Destiny.Definitions.DestinyLocationDefinition locationReleases false */

/* polymorph Destiny.Definitions.DestinyLocationDefinition redacted false */

/* polymorph Destiny.Definitions.DestinyLocationDefinition vendorHash false */

// Validate validates this destiny definitions destiny location definition
func (m *DestinyDefinitionsDestinyLocationDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyLocationDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyLocationDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyLocationDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
