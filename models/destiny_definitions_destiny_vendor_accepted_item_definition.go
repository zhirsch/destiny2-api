// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyVendorAcceptedItemDefinition If you ever wondered how the Vault works, here it is.
// The Vault is merely a set of inventory buckets that exist on your Profile/Account level. When you transfer items in the Vault, the game is using the Vault Vendor's DestinyVendorAcceptedItemDefinitions to see where the appropriate destination bucket is for the source bucket from whence your item is moving. If it finds such an entry, it transfers the item to the other bucket.
// The mechanics for Postmaster works similarly, which is also a vendor. All driven by Accepted Items.
// swagger:model Destiny.Definitions.DestinyVendorAcceptedItemDefinition

type DestinyDefinitionsDestinyVendorAcceptedItemDefinition struct {

	// The "source" bucket for a transfer. When a user wants to transfer an item, the appropriate DestinyVendorDefinition's acceptedItems property is evaluated, looking for an entry where acceptedInventoryBucketHash matches the bucket that the item being transferred is currently located. If it exists, the item will be transferred into whatever bucket is defined by destinationInventoryBucketHash.
	AcceptedInventoryBucketHash uint32 `json:"acceptedInventoryBucketHash,omitempty"`

	// This is the bucket where the item being transferred will be put, given that it was being transferred *from* the bucket defined in acceptedInventoryBucketHash.
	DestinationInventoryBucketHash uint32 `json:"destinationInventoryBucketHash,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyVendorAcceptedItemDefinition acceptedInventoryBucketHash false */

/* polymorph Destiny.Definitions.DestinyVendorAcceptedItemDefinition destinationInventoryBucketHash false */

// Validate validates this destiny definitions destiny vendor accepted item definition
func (m *DestinyDefinitionsDestinyVendorAcceptedItemDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorAcceptedItemDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorAcceptedItemDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyVendorAcceptedItemDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
