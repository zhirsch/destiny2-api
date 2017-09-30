// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyHistoricalStatsDestinyHistoricalStatsWithMergedResults destiny historical stats destiny historical stats with merged results
// swagger:model destinyHistoricalStatsDestinyHistoricalStatsWithMergedResults

type DestinyHistoricalStatsDestinyHistoricalStatsWithMergedResults map[string]DestinyHistoricalStatsDestinyHistoricalStatsByPeriod

// Validate validates this destiny historical stats destiny historical stats with merged results
func (m DestinyHistoricalStatsDestinyHistoricalStatsWithMergedResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyHistoricalStatsDestinyHistoricalStatsWithMergedResults(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}