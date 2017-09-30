// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyHistoricalStatsValuePair destiny historical stats destiny historical stats value pair
// swagger:model Destiny.HistoricalStats.DestinyHistoricalStatsValuePair

type DestinyHistoricalStatsDestinyHistoricalStatsValuePair struct {

	// Localized formated version of the value.
	DisplayValue string `json:"displayValue,omitempty"`

	// Raw value of the statistic
	Value float64 `json:"value,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsValuePair displayValue false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsValuePair value false */

// Validate validates this destiny historical stats destiny historical stats value pair
func (m *DestinyHistoricalStatsDestinyHistoricalStatsValuePair) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsValuePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsValuePair) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyHistoricalStatsValuePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
