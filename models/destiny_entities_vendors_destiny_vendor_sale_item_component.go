// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesVendorsDestinyVendorSaleItemComponent Request this component if you want the details about an item being sold in relation to the character making the request: whether the character can buy it, whether they can afford it, and other data related to purchasing the item.
// Note that if you want instance, stats, etc... data for the item, you'll have to request additional components such as ItemInstances, ItemPerks etc... and acquire them from the DestinyVendorResponse's "items" property.
// swagger:model Destiny.Entities.Vendors.DestinyVendorSaleItemComponent

type DestinyEntitiesVendorsDestinyVendorSaleItemComponent struct {

	// costs
	Costs DestinyEntitiesVendorsDestinyVendorSaleItemComponentCosts `json:"costs"`

	// Indexes in to the "failureStrings" lookup table in DestinyVendorDefinition for the given Vendor. Gives some more reliable failure information for why you can't purchase an item.
	// It is preferred to use these over requiredUnlocks and unlockStatuses: the latter are provided mostly in case someone can do something interesting with it that I didn't anticipate.
	FailureIndexes []int32 `json:"failureIndexes"`

	// The hash of the item being sold, as a quick shortcut for looking up the DestinyInventoryItemDefinition of the sale item.
	ItemHash uint32 `json:"itemHash,omitempty"`

	// If you can't buy the item due to a complex character state, these will be hashes for DestinyUnlockDefinitions that you can check to see messages regarding the failure (if the unlocks have human readable information: it is not guaranteed that Unlocks will have human readable strings, and your application will have to handle that)
	// Prefer using failureIndexes instead. These are provided for informational purposes, but have largely been supplanted by failureIndexes.
	RequiredUnlocks []uint32 `json:"requiredUnlocks"`

	// A flag indicating whether the requesting character can buy the item, and if not the reasons why the character can't buy it.
	SaleStatus DestinyVendorItemStatus `json:"saleStatus,omitempty"`

	// unlock statuses
	UnlockStatuses DestinyEntitiesVendorsDestinyVendorSaleItemComponentUnlockStatuses `json:"unlockStatuses"`

	// The index into the DestinyVendorDefinition.itemList property. Note that this means Vendor data *is* Content Version dependent: make sure you have the latest content before you use Vendor data, or these indexes may mismatch.
	// Most systems avoid this problem, but Vendors is one area where we are unable to reasonably avoid content dependency at the moment.
	VendorItemIndex int32 `json:"vendorItemIndex,omitempty"`
}

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent costs false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent failureIndexes false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent itemHash false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent requiredUnlocks false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent saleStatus false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent unlockStatuses false */

/* polymorph Destiny.Entities.Vendors.DestinyVendorSaleItemComponent vendorItemIndex false */

// Validate validates this destiny entities vendors destiny vendor sale item component
func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailureIndexes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRequiredUnlocks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSaleStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) validateFailureIndexes(formats strfmt.Registry) error {

	if swag.IsZero(m.FailureIndexes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) validateRequiredUnlocks(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredUnlocks) { // not required
		return nil
	}

	return nil
}

func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) validateSaleStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SaleStatus) { // not required
		return nil
	}

	if err := m.SaleStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("saleStatus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesVendorsDestinyVendorSaleItemComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesVendorsDestinyVendorSaleItemComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}