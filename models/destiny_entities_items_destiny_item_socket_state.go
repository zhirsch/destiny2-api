// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesItemsDestinyItemSocketState The status of a given item's socket. (which plug is inserted, if any: whether it is enabled, what "reusable" plugs can be inserted, etc...)
// swagger:model Destiny.Entities.Items.DestinyItemSocketState

type DestinyEntitiesItemsDestinyItemSocketState struct {

	// If a plug is inserted but not enabled, this will be populated with indexes into the plug item definition's plug.enabledRules property, so that you can show the reasons why it is not enabled.
	EnableFailIndexes []int32 `json:"enableFailIndexes"`

	// Even if a plug is inserted, it doesn't mean it's enabled.
	// This flag indicates whether the plug is active and providing its benefits.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// The currently active plug, if any.
	// Note that, because all plugs are statically defined, its effect on stats and perks can be statically determined using the plug item's definition. The stats and perks can be taken at face value on the plug item as the stats and perks it will provide to the user/item.
	PlugHash uint32 `json:"plugHash,omitempty"`

	// If the item supports reusable plugs, this is the list of plug item hashes that are currently allowed to be used for this socket. (sometimes restrictions may cause reusable plugs defined on the item definition to not be valid, so you should trust the instanced reusablePlugHashes list rather than the definition's list)
	// A Reusable Plug is a plug that you can *always* insert into this socket, regardless of whether or not you have the plug in your inventory. In practice, a socket will *either* have reusable plugs *or* it will allow for plugs in your inventory to be inserted. See DestinyInventoryItemDefinition.socket for more info.
	ReusablePlugHashes []uint32 `json:"reusablePlugHashes"`
}

/* polymorph Destiny.Entities.Items.DestinyItemSocketState enableFailIndexes false */

/* polymorph Destiny.Entities.Items.DestinyItemSocketState isEnabled false */

/* polymorph Destiny.Entities.Items.DestinyItemSocketState plugHash false */

/* polymorph Destiny.Entities.Items.DestinyItemSocketState reusablePlugHashes false */

// Validate validates this destiny entities items destiny item socket state
func (m *DestinyEntitiesItemsDestinyItemSocketState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnableFailIndexes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReusablePlugHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyEntitiesItemsDestinyItemSocketState) validateEnableFailIndexes(formats strfmt.Registry) error {

	if swag.IsZero(m.EnableFailIndexes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyEntitiesItemsDestinyItemSocketState) validateReusablePlugHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ReusablePlugHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemSocketState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemSocketState) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesItemsDestinyItemSocketState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
