// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition Data related to what happens while a plug is being inserted, mostly for UI purposes.
// swagger:model Destiny.Definitions.Sockets.DestinyInsertPlugActionDefinition

type DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition struct {

	// How long it takes for the Plugging of the item to be completed once it is initiated, if you care.
	ActionExecuteSeconds int32 `json:"actionExecuteSeconds,omitempty"`
}

/* polymorph Destiny.Definitions.Sockets.DestinyInsertPlugActionDefinition actionExecuteSeconds false */

// Validate validates this destiny definitions sockets destiny insert plug action definition
func (m *DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsSocketsDestinyInsertPlugActionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
