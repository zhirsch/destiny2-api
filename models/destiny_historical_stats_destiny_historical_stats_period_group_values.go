// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues Collection of stats for the period.
// swagger:model destinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues

type DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue

// Validate validates this destiny historical stats destiny historical stats period group values
func (m DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
