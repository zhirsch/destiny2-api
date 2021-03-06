// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesItemsDestinyItemPerksComponent Instanced items can have perks: benefits that the item bestows.
// These are related to DestinySandboxPerkDefinition, and sometimes - but not always - have human readable info. When they do, they are the icons and text that you see in an item's tooltip.
// Talent Grids, Sockets, and the item itself can apply Perks, which are then summarized here for your convenience.
// swagger:model Destiny.Entities.Items.DestinyItemPerksComponent

type DestinyEntitiesItemsDestinyItemPerksComponent struct {

	// perks
	Perks DestinyEntitiesItemsDestinyItemPerksComponentPerks `json:"perks"`
}

/* polymorph Destiny.Entities.Items.DestinyItemPerksComponent perks false */

// Validate validates this destiny entities items destiny item perks component
func (m *DestinyEntitiesItemsDestinyItemPerksComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemPerksComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemPerksComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesItemsDestinyItemPerksComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
