// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyHistoricalStatsDestinyPostGameCarnageReportExtendedDataValues Collection of stats for the player in this activity.
// swagger:model destinyHistoricalStatsDestinyPostGameCarnageReportExtendedDataValues

type DestinyHistoricalStatsDestinyPostGameCarnageReportExtendedDataValues map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue

// Validate validates this destiny historical stats destiny post game carnage report extended data values
func (m DestinyHistoricalStatsDestinyPostGameCarnageReportExtendedDataValues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", DestinyHistoricalStatsDestinyPostGameCarnageReportExtendedDataValues(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
