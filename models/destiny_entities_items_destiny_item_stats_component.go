// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesItemsDestinyItemStatsComponent If you want the stats on an item's instanced data, get this component.
// These are stats like Attack, Defense etc... and *not* historical stats.
// Note that some stats have additional computation in-game at runtime - for instance, Magazine Size - and thus these stats might not be 100% accurate compared to what you see in-game for some stats. I know, it sucks. I hate it too.
// swagger:model Destiny.Entities.Items.DestinyItemStatsComponent

type DestinyEntitiesItemsDestinyItemStatsComponent struct {

	// stats
	Stats DestinyEntitiesItemsDestinyItemStatsComponentStats `json:"stats,omitempty"`
}

/* polymorph Destiny.Entities.Items.DestinyItemStatsComponent stats false */

// Validate validates this destiny entities items destiny item stats component
func (m *DestinyEntitiesItemsDestinyItemStatsComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemStatsComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemStatsComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesItemsDestinyItemStatsComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}