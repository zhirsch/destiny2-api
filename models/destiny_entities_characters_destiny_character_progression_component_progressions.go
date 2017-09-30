// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyEntitiesCharactersDestinyCharacterProgressionComponentProgressions A Dictionary of all known progressions for the Character, keyed by the Progression's hash.
// Not all progressions have user-facing data, but those who do will have that data contained in the DestinyProgressionDefinition.
// swagger:model destinyEntitiesCharactersDestinyCharacterProgressionComponentProgressions

type DestinyEntitiesCharactersDestinyCharacterProgressionComponentProgressions map[string]DestinyDestinyProgression

// Validate validates this destiny entities characters destiny character progression component progressions
func (m DestinyEntitiesCharactersDestinyCharacterProgressionComponentProgressions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyEntitiesCharactersDestinyCharacterProgressionComponentProgressions(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
