// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyHistoricalWeaponStats destiny historical stats destiny historical weapon stats
// swagger:model Destiny.HistoricalStats.DestinyHistoricalWeaponStats

type DestinyHistoricalStatsDestinyHistoricalWeaponStats struct {

	// The hash ID of the item definition that describes the weapon.
	ReferenceID uint32 `json:"referenceId,omitempty"`

	// values
	Values DestinyHistoricalStatsDestinyHistoricalWeaponStatsValues `json:"values,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyHistoricalWeaponStats referenceId false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalWeaponStats values false */

// Validate validates this destiny historical stats destiny historical weapon stats
func (m *DestinyHistoricalStatsDestinyHistoricalWeaponStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalWeaponStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalWeaponStats) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyHistoricalWeaponStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
