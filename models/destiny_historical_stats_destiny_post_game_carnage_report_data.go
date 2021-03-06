// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyPostGameCarnageReportData destiny historical stats destiny post game carnage report data
// swagger:model Destiny.HistoricalStats.DestinyPostGameCarnageReportData

type DestinyHistoricalStatsDestinyPostGameCarnageReportData struct {

	// Details about the activity.
	ActivityDetails *DestinyHistoricalStatsDestinyHistoricalStatsActivity `json:"activityDetails,omitempty"`

	// entries
	Entries DestinyHistoricalStatsDestinyPostGameCarnageReportDataEntries `json:"entries"`

	// Date and time for the activity.
	Period strfmt.DateTime `json:"period,omitempty"`

	// teams
	Teams DestinyHistoricalStatsDestinyPostGameCarnageReportDataTeams `json:"teams"`
}

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportData activityDetails false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportData entries false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportData period false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportData teams false */

// Validate validates this destiny historical stats destiny post game carnage report data
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportData) Validate(formats strfmt.Registry) error {
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

func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportData) validateActivityDetails(formats strfmt.Registry) error {

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
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportData) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyPostGameCarnageReportData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
