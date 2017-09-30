// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DictionaryComponentResponseOfint32AndDestinyItemStatsComponentData dictionary component response ofint32 and destiny item stats component data
// swagger:model dictionaryComponentResponseOfint32AndDestinyItemStatsComponentData

type DictionaryComponentResponseOfint32AndDestinyItemStatsComponentData map[string]DestinyEntitiesItemsDestinyItemStatsComponent

// Validate validates this dictionary component response ofint32 and destiny item stats component data
func (m DictionaryComponentResponseOfint32AndDestinyItemStatsComponentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DictionaryComponentResponseOfint32AndDestinyItemStatsComponentData(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}