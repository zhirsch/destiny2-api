// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition Information about a single inventory bucket in a vendor flyout UI and how it is shown.
// swagger:model Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition

type DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition struct {

	// If true, the inventory bucket should be able to be collapsed visually.
	Collapsible bool `json:"collapsible,omitempty"`

	// The inventory bucket whose contents should be shown.
	InventoryBucketHash uint32 `json:"inventoryBucketHash,omitempty"`

	// The methodology to use for sorting items from the flyout.
	SortItemsBy DestinyDestinyItemSortType `json:"sortItemsBy,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition collapsible false */

/* polymorph Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition inventoryBucketHash false */

/* polymorph Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition sortItemsBy false */

// Validate validates this destiny definitions destiny vendor inventory flyout bucket definition
func (m *DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortItemsBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition) validateSortItemsBy(formats strfmt.Registry) error {

	if swag.IsZero(m.SortItemsBy) { // not required
		return nil
	}

	if err := m.SortItemsBy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sortItemsBy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}