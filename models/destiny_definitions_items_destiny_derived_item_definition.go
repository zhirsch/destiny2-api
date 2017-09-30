// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsItemsDestinyDerivedItemDefinition This is a reference to, and summary data for, a specific item that you can get as a result of Using or Acquiring some other Item (For example, this could be summary information for an Emote that you can get by opening an an Eververse Box) See DestinyDerivedItemCategoryDefinition for more information.
// swagger:model Destiny.Definitions.Items.DestinyDerivedItemDefinition

type DestinyDefinitionsItemsDestinyDerivedItemDefinition struct {

	// An icon for the item.
	IconPath string `json:"iconPath,omitempty"`

	// A brief description of the item.
	ItemDescription string `json:"itemDescription,omitempty"`

	// Additional details about the derived item, in addition to the description.
	ItemDetail string `json:"itemDetail,omitempty"`

	// The hash for the DestinyInventoryItemDefinition of this derived item, if there is one. Sometimes we are given this information as a manual override, in which case there won't be an actual DestinyInventoryItemDefinition for what we display, but you can still show the strings from this object itself.
	ItemHash uint32 `json:"itemHash,omitempty"`

	// The name of the derived item.
	ItemName string `json:"itemName,omitempty"`

	// If the item was derived from a "Preview Vendor", this will be an index into the DestinyVendorDefinition's itemList property. Otherwise, -1.
	VendorItemIndex int32 `json:"vendorItemIndex,omitempty"`
}

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition iconPath false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition itemDescription false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition itemDetail false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition itemHash false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition itemName false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemDefinition vendorItemIndex false */

// Validate validates this destiny definitions items destiny derived item definition
func (m *DestinyDefinitionsItemsDestinyDerivedItemDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsItemsDestinyDerivedItemDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsItemsDestinyDerivedItemDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsItemsDestinyDerivedItemDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
