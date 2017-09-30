// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemTranslationBlockDefinition This Block defines the rendering data associated with the item, if any.
// swagger:model Destiny.Definitions.DestinyItemTranslationBlockDefinition

type DestinyDefinitionsDestinyItemTranslationBlockDefinition struct {

	// arrangements
	Arrangements DestinyDefinitionsDestinyItemTranslationBlockDefinitionArrangements `json:"arrangements"`

	// custom dyes
	CustomDyes DestinyDefinitionsDestinyItemTranslationBlockDefinitionCustomDyes `json:"customDyes"`

	// default dyes
	DefaultDyes DestinyDefinitionsDestinyItemTranslationBlockDefinitionDefaultDyes `json:"defaultDyes"`

	// has geometry
	HasGeometry bool `json:"hasGeometry,omitempty"`

	// locked dyes
	LockedDyes DestinyDefinitionsDestinyItemTranslationBlockDefinitionLockedDyes `json:"lockedDyes"`

	// weapon pattern hash
	WeaponPatternHash uint32 `json:"weaponPatternHash,omitempty"`

	// weapon pattern identifier
	WeaponPatternIdentifier string `json:"weaponPatternIdentifier,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition arrangements false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition customDyes false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition defaultDyes false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition hasGeometry false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition lockedDyes false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition weaponPatternHash false */

/* polymorph Destiny.Definitions.DestinyItemTranslationBlockDefinition weaponPatternIdentifier false */

// Validate validates this destiny definitions destiny item translation block definition
func (m *DestinyDefinitionsDestinyItemTranslationBlockDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemTranslationBlockDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemTranslationBlockDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemTranslationBlockDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}