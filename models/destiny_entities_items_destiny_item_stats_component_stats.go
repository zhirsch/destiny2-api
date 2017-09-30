// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyEntitiesItemsDestinyItemStatsComponentStats If the item has stats that it provides (damage, defense, etc...), it will be given here.
// swagger:model destinyEntitiesItemsDestinyItemStatsComponentStats

type DestinyEntitiesItemsDestinyItemStatsComponentStats map[string]DestinyDestinyStat

// Validate validates this destiny entities items destiny item stats component stats
func (m DestinyEntitiesItemsDestinyItemStatsComponentStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyEntitiesItemsDestinyItemStatsComponentStats(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
