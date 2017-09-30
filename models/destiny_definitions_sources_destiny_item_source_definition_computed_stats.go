// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyDefinitionsSourcesDestinyItemSourceDefinitionComputedStats The stats computed for this level/quality range.
// swagger:model destinyDefinitionsSourcesDestinyItemSourceDefinitionComputedStats

type DestinyDefinitionsSourcesDestinyItemSourceDefinitionComputedStats map[string]DestinyDefinitionsDestinyInventoryItemStatDefinition

// Validate validates this destiny definitions sources destiny item source definition computed stats
func (m DestinyDefinitionsSourcesDestinyItemSourceDefinitionComputedStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyDefinitionsSourcesDestinyItemSourceDefinitionComputedStats(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}