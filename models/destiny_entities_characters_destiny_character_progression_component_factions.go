// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyEntitiesCharactersDestinyCharacterProgressionComponentFactions A dictionary of all known Factions, keyed by the Faction's hash. It contains data about this character's status with the faction.
// swagger:model destinyEntitiesCharactersDestinyCharacterProgressionComponentFactions

type DestinyEntitiesCharactersDestinyCharacterProgressionComponentFactions map[string]DestinyProgressionDestinyFactionProgression

// Validate validates this destiny entities characters destiny character progression component factions
func (m DestinyEntitiesCharactersDestinyCharacterProgressionComponentFactions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyEntitiesCharactersDestinyCharacterProgressionComponentFactions(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
