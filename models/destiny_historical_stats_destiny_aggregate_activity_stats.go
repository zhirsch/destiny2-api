// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyAggregateActivityStats destiny historical stats destiny aggregate activity stats
// swagger:model Destiny.HistoricalStats.DestinyAggregateActivityStats

type DestinyHistoricalStatsDestinyAggregateActivityStats struct {

	// Hash ID that can be looked up in the DestinyActivityTable.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// values
	Values DestinyHistoricalStatsDestinyAggregateActivityStatsValues `json:"values,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyAggregateActivityStats activityHash false */

/* polymorph Destiny.HistoricalStats.DestinyAggregateActivityStats values false */

// Validate validates this destiny historical stats destiny aggregate activity stats
func (m *DestinyHistoricalStatsDestinyAggregateActivityStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyAggregateActivityStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyAggregateActivityStats) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyAggregateActivityStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
