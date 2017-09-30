// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup destiny historical stats destiny historical stats period group
// swagger:model Destiny.HistoricalStats.DestinyHistoricalStatsPeriodGroup

type DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup struct {

	// If the period group is for a specific activity, this property will be set.
	ActivityDetails *DestinyHistoricalStatsDestinyHistoricalStatsActivity `json:"activityDetails,omitempty"`

	// Period for the group. If the stat periodType is day, then this will have a specific day. If the type is monthly, then this value will be the first day of the applicable month. This value is not set when the periodType is 'all time'.
	Period strfmt.DateTime `json:"period,omitempty"`

	// values
	Values DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroupValues `json:"values,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsPeriodGroup activityDetails false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsPeriodGroup period false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsPeriodGroup values false */

// Validate validates this destiny historical stats destiny historical stats period group
func (m *DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup) validateActivityDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivityDetails) { // not required
		return nil
	}

	if m.ActivityDetails != nil {

		if err := m.ActivityDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activityDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}