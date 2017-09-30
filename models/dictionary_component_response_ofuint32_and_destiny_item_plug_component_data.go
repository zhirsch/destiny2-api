// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData dictionary component response ofuint32 and destiny item plug component data
// swagger:model dictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData

type DictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData map[string]DestinyComponentsItemsDestinyItemPlugComponent

// Validate validates this dictionary component response ofuint32 and destiny item plug component data
func (m DictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DictionaryComponentResponseOfuint32AndDestinyItemPlugComponentData(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}