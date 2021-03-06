// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyVendorsDestinyVendorReceipt If a character purchased an item that is refundable, a Vendor Receipt will be created on the user's Destiny Profile. These expire after a configurable period of time, but until then can be used to get refunds on items. BNet does not provide the ability to refund a purchase *yet*, but you know.
// swagger:model Destiny.Vendors.DestinyVendorReceipt

type DestinyVendorsDestinyVendorReceipt struct {

	// currency paid
	CurrencyPaid DestinyVendorsDestinyVendorReceiptCurrencyPaid `json:"currencyPaid"`

	// The date at which this receipt is rendered invalid.
	ExpiresOn strfmt.DateTime `json:"expiresOn,omitempty"`

	// The item that was received, and its quantity.
	ItemReceived *DestinyDestinyItemQuantity `json:"itemReceived,omitempty"`

	// The unlock flag used to determine whether you still have the purchased item.
	LicenseUnlockHash uint32 `json:"licenseUnlockHash,omitempty"`

	// The ID of the character who made the purchase.
	PurchasedByCharacterID int64 `json:"purchasedByCharacterId,omitempty"`

	// Whether you can get a refund, and what happens in order for the refund to be received. See the DestinyVendorItemRefundPolicy enum for details.
	RefundPolicy DestinyDestinyVendorItemRefundPolicy `json:"refundPolicy,omitempty"`

	// The identifier of this receipt.
	SequenceNumber int32 `json:"sequenceNumber,omitempty"`

	// The seconds since epoch at which this receipt is rendered invalid.
	TimeToExpiration int64 `json:"timeToExpiration,omitempty"`
}

/* polymorph Destiny.Vendors.DestinyVendorReceipt currencyPaid false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt expiresOn false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt itemReceived false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt licenseUnlockHash false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt purchasedByCharacterId false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt refundPolicy false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt sequenceNumber false */

/* polymorph Destiny.Vendors.DestinyVendorReceipt timeToExpiration false */

// Validate validates this destiny vendors destiny vendor receipt
func (m *DestinyVendorsDestinyVendorReceipt) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemReceived(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundPolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyVendorsDestinyVendorReceipt) validateItemReceived(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemReceived) { // not required
		return nil
	}

	if m.ItemReceived != nil {

		if err := m.ItemReceived.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemReceived")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyVendorsDestinyVendorReceipt) validateRefundPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RefundPolicy) { // not required
		return nil
	}

	if err := m.RefundPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("refundPolicy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyVendorsDestinyVendorReceipt) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyVendorsDestinyVendorReceipt) UnmarshalBinary(b []byte) error {
	var res DestinyVendorsDestinyVendorReceipt
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
