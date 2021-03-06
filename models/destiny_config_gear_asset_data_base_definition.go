// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyConfigGearAssetDataBaseDefinition destiny config gear asset data base definition
// swagger:model Destiny.Config.GearAssetDataBaseDefinition

type DestinyConfigGearAssetDataBaseDefinition struct {

	// path
	Path string `json:"path,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

/* polymorph Destiny.Config.GearAssetDataBaseDefinition path false */

/* polymorph Destiny.Config.GearAssetDataBaseDefinition version false */

// Validate validates this destiny config gear asset data base definition
func (m *DestinyConfigGearAssetDataBaseDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyConfigGearAssetDataBaseDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyConfigGearAssetDataBaseDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyConfigGearAssetDataBaseDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
