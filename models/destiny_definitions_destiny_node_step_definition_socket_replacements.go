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

// DestinyDefinitionsDestinyNodeStepDefinitionSocketReplacements If this step is activated, this will be a list of information used to replace socket items with new Plugs. See DestinyInventoryItemDefinition for more information about sockets and plugs.
// swagger:model destinyDefinitionsDestinyNodeStepDefinitionSocketReplacements

type DestinyDefinitionsDestinyNodeStepDefinitionSocketReplacements []*DestinyDefinitionsDestinyNodeSocketReplaceResponse

// Validate validates this destiny definitions destiny node step definition socket replacements
func (m DestinyDefinitionsDestinyNodeStepDefinitionSocketReplacements) Validate(formats strfmt.Registry) error {
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
