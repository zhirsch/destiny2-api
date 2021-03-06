// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyVendorCategoryEntryDefinition This is the definition for a single Vendor Category, into which Sale Items are grouped.
// swagger:model Destiny.Definitions.DestinyVendorCategoryEntryDefinition

type DestinyDefinitionsDestinyVendorCategoryEntryDefinition struct {

	// The localized string for making purchases from this category, if it is different from the vendor's string for purchasing.
	BuyStringOverride string `json:"buyStringOverride,omitempty"`

	// The hashed identifier for the category. (note that this is NOT pointing to a DestinyVendorCategoryDefinition, it's confusing but this is a sale item category in a vendor, not a categorization of vendors themselves)
	CategoryHash uint32 `json:"categoryHash,omitempty"`

	// The string identifier of the category.
	CategoryID string `json:"categoryId,omitempty"`

	// The index of the category in the original category definitions for the vendor.
	CategoryIndex int32 `json:"categoryIndex,omitempty"`

	// If the category is disabled, this is the localized description to show.
	DisabledDescription string `json:"disabledDescription,omitempty"`

	// The localized title of the category.
	DisplayTitle string `json:"displayTitle,omitempty"`

	// True if this category doesn't allow purchases.
	HideFromRegularPurchase bool `json:"hideFromRegularPurchase,omitempty"`

	// If you don't have the currency required to buy items from this category, should the items be hidden?
	HideIfNoCurrency bool `json:"hideIfNoCurrency,omitempty"`

	// If this category has an overlay prompt that should appear, this contains the details of that prompt.
	Overlay *DestinyDefinitionsDestinyVendorCategoryOverlayDefinition `json:"overlay,omitempty"`

	// The amount of items that will be available when this category is shown.
	QuantityAvailable int32 `json:"quantityAvailable,omitempty"`

	// If items aren't up for sale in this category, should we still show them (greyed out)?
	ShowUnavailableItems bool `json:"showUnavailableItems,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition buyStringOverride false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition categoryHash false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition categoryId false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition categoryIndex false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition disabledDescription false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition displayTitle false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition hideFromRegularPurchase false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition hideIfNoCurrency false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition overlay false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition quantityAvailable false */

/* polymorph Destiny.Definitions.DestinyVendorCategoryEntryDefinition showUnavailableItems false */

// Validate validates this destiny definitions destiny vendor category entry definition
func (m *DestinyDefinitionsDestinyVendorCategoryEntryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverlay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyVendorCategoryEntryDefinition) validateOverlay(formats strfmt.Registry) error {

	if swag.IsZero(m.Overlay) { // not required
		return nil
	}

	if m.Overlay != nil {

		if err := m.Overlay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overlay")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorCategoryEntryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyVendorCategoryEntryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyVendorCategoryEntryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
