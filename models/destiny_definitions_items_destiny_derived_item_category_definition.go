// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition A shortcut for the fact that some items have a "Preview Vendor" - See DestinyInventoryItemDefinition.preview.previewVendorHash - that is intended to be used to show what items you can get as a result of acquiring or using this item.
// A common example of this in Destiny 1 was Eververse "Boxes," which could have many possible items. This "Preview Vendor" is not a vendor you can actually see in the game, but it defines categories and sale items for all of the possible items you could get from the Box so that the game can show them to you. We summarize that info here so that you don't have to do that Vendor lookup and aggregation manually.
// swagger:model Destiny.Definitions.Items.DestinyDerivedItemCategoryDefinition

type DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition struct {

	// The localized string for the category title. This will be something describing the items you can get as a group, or your likelihood/the quantity you'll get.
	CategoryDescription string `json:"categoryDescription,omitempty"`

	// items
	Items DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinitionItems `json:"items"`
}

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemCategoryDefinition categoryDescription false */

/* polymorph Destiny.Definitions.Items.DestinyDerivedItemCategoryDefinition items false */

// Validate validates this destiny definitions items destiny derived item category definition
func (m *DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
