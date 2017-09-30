// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConfigGroupTheme config group theme
// swagger:model Config.GroupTheme

type ConfigGroupTheme struct {

	// description
	Description string `json:"description,omitempty"`

	// folder
	Folder string `json:"folder,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

/* polymorph Config.GroupTheme description false */

/* polymorph Config.GroupTheme folder false */

/* polymorph Config.GroupTheme name false */

// Validate validates this config group theme
func (m *ConfigGroupTheme) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigGroupTheme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigGroupTheme) UnmarshalBinary(b []byte) error {
	var res ConfigGroupTheme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
