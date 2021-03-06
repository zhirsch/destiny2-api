// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry destiny historical stats destiny post game carnage report team entry
// swagger:model Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry

type DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry struct {

	// Score earned by the team
	Score *DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"score,omitempty"`

	// Team's standing relative to other teams.
	Standing *DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"standing,omitempty"`

	// Integer ID for the team.
	TeamID int32 `json:"teamId,omitempty"`

	// Alpha or Bravo
	TeamName string `json:"teamName,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry score false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry standing false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry teamId false */

/* polymorph Destiny.HistoricalStats.DestinyPostGameCarnageReportTeamEntry teamName false */

// Validate validates this destiny historical stats destiny post game carnage report team entry
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStanding(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry) validateScore(formats strfmt.Registry) error {

	if swag.IsZero(m.Score) { // not required
		return nil
	}

	if m.Score != nil {

		if err := m.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry) validateStanding(formats strfmt.Registry) error {

	if swag.IsZero(m.Standing) { // not required
		return nil
	}

	if m.Standing != nil {

		if err := m.Standing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
