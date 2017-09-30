// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesVendorsDestinyVendorComponent This component contains essential/summary information about the vendor.
// swagger:model Destiny.Entities.Vendors.DestinyVendorComponent

type DestinyEntitiesVendorsDestinyVendorComponent struct {

	// Long ago, we thought it would be a good idea to have special UI that showed whether or not you've seen a Vendor's inventory after cycling.
	// For now, we don't have that UI anymore. This property still exists for historical purposes. Don't worry about it.
	AckState *UserAckState `json:"ackState,omitempty"`

	// If True, you can purchase from the Vendor.
	// Theoretically, Vendors can be restricted from selling items. In practice, none do that (yet?).
	CanPurchase bool `json:"canPurchase,omitempty"`

	// If True, the Vendor is currently accessible.
	// If False, they may not actually be visible in the world at the moment.
	Enabled bool `json:"enabled,omitempty"`

	// The date when this vendor's inventory will next rotate/refresh.
	NextRefreshDate strfmt.DateTime `json:"nextRefreshDate,omitempty"`

	// If the Vendor has a related Reputation, this is the Progression data that represents the character's Reputation level with this Vendor.
	Progression *DestinyDestinyProgression `json:"progression,omitempty"`

	// The unique identifier for the vendor. Use it to look up their DestinyVendorDefinition.
	VendorHash uint32 `json:"vendorHash,omitempty"`
}

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent ackState false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent canPurchase false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent enabled false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent nextRefreshDate false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent progression false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorComponent vendorHash false */

// Validate validates this destiny entities vendors destiny vendor component
func (m *DestinyEntitiesVendorsDestinyVendorComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAckState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProgression(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyEntitiesVendorsDestinyVendorComponent) validateAckState(formats strfmt.Registry) error {

	if swag.IsZero(m.AckState) { // not required
		return nil
	}

	if m.AckState != nil {

		if err := m.AckState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ackState")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyEntitiesVendorsDestinyVendorComponent) validateProgression(formats strfmt.Registry) error {

	if swag.IsZero(m.Progression) { // not required
		return nil
	}

	if m.Progression != nil {

		if err := m.Progression.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progression")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesVendorsDestinyVendorComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesVendorsDestinyVendorComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesVendorsDestinyVendorComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}