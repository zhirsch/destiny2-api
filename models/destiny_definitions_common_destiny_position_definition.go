// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsCommonDestinyPositionDefinition destiny definitions common destiny position definition
// swagger:model Destiny.Definitions.Common.DestinyPositionDefinition

type DestinyDefinitionsCommonDestinyPositionDefinition struct {

	// x
	X int32 `json:"x,omitempty"`

	// y
	Y int32 `json:"y,omitempty"`

	// z
	Z int32 `json:"z,omitempty"`
}

/* polymorph Destiny.Definitions.Common.DestinyPositionDefinition x false */

/* polymorph Destiny.Definitions.Common.DestinyPositionDefinition y false */

/* polymorph Destiny.Definitions.Common.DestinyPositionDefinition z false */

// Validate validates this destiny definitions common destiny position definition
func (m *DestinyDefinitionsCommonDestinyPositionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsCommonDestinyPositionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsCommonDestinyPositionDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsCommonDestinyPositionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}