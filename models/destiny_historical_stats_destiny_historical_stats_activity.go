// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyHistoricalStatsActivity destiny historical stats destiny historical stats activity
// swagger:model Destiny.HistoricalStats.DestinyHistoricalStatsActivity

type DestinyHistoricalStatsDestinyHistoricalStatsActivity struct {

	// director activity hash
	DirectorActivityHash uint32 `json:"directorActivityHash,omitempty"`

	// This value can be used to get additional data about this activity such as who else was playing.
	InstanceID int64 `json:"instanceId,omitempty"`

	// Whether or not the match was a private match.
	IsPrivate bool `json:"isPrivate,omitempty"`

	// Indicates the game mode of the activity.
	Mode DestinyHistoricalStatsDefinitionsDestinyActivityModeType `json:"mode,omitempty"`

	// modes
	Modes []DestinyHistoricalStatsDefinitionsDestinyActivityModeType `json:"modes"`

	// Hash ID that can be looked up in the DestinyActivityTable.
	ReferenceID uint32 `json:"referenceId,omitempty"`
}

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity directorActivityHash false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity instanceId false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity isPrivate false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity mode false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity modes false */

/* polymorph Destiny.HistoricalStats.DestinyHistoricalStatsActivity referenceId false */

// Validate validates this destiny historical stats destiny historical stats activity
func (m *DestinyHistoricalStatsDestinyHistoricalStatsActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyHistoricalStatsDestinyHistoricalStatsActivity) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *DestinyHistoricalStatsDestinyHistoricalStatsActivity) validateModes(formats strfmt.Registry) error {

	if swag.IsZero(m.Modes) { // not required
		return nil
	}

	for i := 0; i < len(m.Modes); i++ {

		if err := m.Modes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyHistoricalStatsActivity) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyHistoricalStatsActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
