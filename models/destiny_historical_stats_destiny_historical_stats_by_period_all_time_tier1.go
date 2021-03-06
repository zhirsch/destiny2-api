// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyHistoricalStatsDestinyHistoricalStatsByPeriodAllTimeTier1 destiny historical stats destiny historical stats by period all time tier1
// swagger:model destinyHistoricalStatsDestinyHistoricalStatsByPeriodAllTimeTier1

type DestinyHistoricalStatsDestinyHistoricalStatsByPeriodAllTimeTier1 map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue

// Validate validates this destiny historical stats destiny historical stats by period all time tier1
func (m DestinyHistoricalStatsDestinyHistoricalStatsByPeriodAllTimeTier1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyHistoricalStatsDestinyHistoricalStatsByPeriodAllTimeTier1(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
