// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyVendorActionDefinition If a vendor can ever end up performing actions, these are the properties that will be related to those actions. I'm not going to bother documenting this yet, as it is unused and unclear if it will ever be used... but in case it is ever populated and someone finds it useful, it is defined here.
// swagger:model Destiny.Definitions.DestinyVendorActionDefinition

type DestinyDefinitionsDestinyVendorActionDefinition struct {

	// action hash
	ActionHash uint32 `json:"actionHash,omitempty"`

	// action Id
	ActionID string `json:"actionId,omitempty"`

	// auto perform action
	AutoPerformAction bool `json:"autoPerformAction,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// execute seconds
	ExecuteSeconds int32 `json:"executeSeconds,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// is positive
	IsPositive bool `json:"isPositive,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// verb
	Verb string `json:"verb,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition actionHash false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition actionId false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition autoPerformAction false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition description false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition executeSeconds false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition icon false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition isPositive false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition name false */

/* polymorph Destiny.Definitions.DestinyVendorActionDefinition verb false */

// Validate validates this destiny definitions destiny vendor action definition
func (m *DestinyDefinitionsDestinyVendorActionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorActionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorActionDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyVendorActionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
