// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyBubbleDefinition Basic identifying data about the bubble. Combine with DestinyDestinationBubbleSettingDefinition - see DestinyDestinationDefinition.bubbleSettings for more information.
// swagger:model Destiny.Definitions.DestinyBubbleDefinition

type DestinyDefinitionsDestinyBubbleDefinition struct {

	// The identifier for the bubble: only guaranteed to be unique within the Destination.
	Hash uint32 `json:"hash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyBubbleDefinition hash false */

// Validate validates this destiny definitions destiny bubble definition
func (m *DestinyDefinitionsDestinyBubbleDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyBubbleDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyBubbleDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyBubbleDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
