// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyHistoricalStatsWithMerged destiny historical stats destiny historical stats with merged
// swagger:model Destiny.HistoricalStats.DestinyHistoricalStatsWithMerged

type DestinyHistoricalStatsDestinyHistoricalStatsWithMerged struct {

	// merged
	Merged *DestinyHistoricalStatsDestinyHistoricalStatsByPeriod `json:"merged,omitempty"`

	// results
	Results DestinyHistoricalStatsDestinyHistoricalStatsWithMergedResults `json:"results,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsWithMerged merged false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsWithMerged results false */

// Validate validates this destiny historical stats destiny historical stats with merged
func (m *DestinyHistoricalStatsDestinyHistoricalStatsWithMerged) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMerged(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyHistoricalStatsDestinyHistoricalStatsWithMerged) validateMerged(formats strfmt.Registry) error {

	if swag.IsZero(m.Merged) { // not required
		return nil
	}

	if m.Merged != nil {

		if err := m.Merged.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("merged")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsWithMerged) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsWithMerged) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyHistoricalStatsWithMerged
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
