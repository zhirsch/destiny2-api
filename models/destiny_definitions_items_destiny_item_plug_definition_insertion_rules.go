// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsItemsDestinyItemPlugDefinitionInsertionRules The rules around when this plug can be inserted into a socket, aside from the socket's individual restrictions.
// The live data DestinyItemPlugComponent.insertFailIndexes will be an index into this array, so you can pull out the failure strings appropriate for the user.
// swagger:model destinyDefinitionsItemsDestinyItemPlugDefinitionInsertionRules

type DestinyDefinitionsItemsDestinyItemPlugDefinitionInsertionRules []*DestinyDefinitionsItemsDestinyPlugRuleDefinition

// Validate validates this destiny definitions items destiny item plug definition insertion rules
func (m DestinyDefinitionsItemsDestinyItemPlugDefinitionInsertionRules) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}