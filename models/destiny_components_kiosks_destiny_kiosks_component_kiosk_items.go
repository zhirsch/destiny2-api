// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyComponentsKiosksDestinyKiosksComponentKioskItems A dictionary keyed by the Kiosk Vendor's hash identifier (use it to look up the DestinyVendorDefinition for the relevant kiosk vendor), and whose value is a list of all the items that the user can "see" in the Kiosk, and any other interesting metadata.
// swagger:model destinyComponentsKiosksDestinyKiosksComponentKioskItems

type DestinyComponentsKiosksDestinyKiosksComponentKioskItems map[string]DestinyComponentsKiosksDestinyKiosksComponentKioskItemsAdditionalProperties

// Validate validates this destiny components kiosks destiny kiosks component kiosk items
func (m DestinyComponentsKiosksDestinyKiosksComponentKioskItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyComponentsKiosksDestinyKiosksComponentKioskItems(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
