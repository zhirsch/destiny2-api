// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponentData dictionary component response ofint64 and destiny item talent grid component data
// swagger:model dictionaryComponentResponseOfint64AndDestinyItemTalentGridComponentData

type DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponentData map[string]DestinyEntitiesItemsDestinyItemTalentGridComponent

// Validate validates this dictionary component response ofint64 and destiny item talent grid component data
func (m DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponentData(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
