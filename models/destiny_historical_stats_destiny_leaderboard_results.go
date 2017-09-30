// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyHistoricalStatsDestinyLeaderboardResults destiny historical stats destiny leaderboard results
// swagger:model Destiny.HistoricalStats.DestinyLeaderboardResults

type DestinyHistoricalStatsDestinyLeaderboardResults struct {

	// Indicate the character ID of the character that is the focal point of the provided leaderboards. May be null, in which case any character from the focus membership can appear in the provided leaderboards.
	FocusCharacterID int64 `json:"focusCharacterId,omitempty"`

	// Indicate the membership ID of the account that is the focal point of the provided leaderboards.
	FocusMembershipID int64 `json:"focusMembershipId,omitempty"`

	// destiny historical stats destiny leaderboard results
	DestinyHistoricalStatsDestinyLeaderboardResults map[string]DestinyHistoricalStatsDestinyLeaderboardResultsAdditionalProperties `json:"-"`
}

/* polymorph Destiny.HistoricalStats.DestinyLeaderboardResults focusCharacterId false */

/* polymorph Destiny.HistoricalStats.DestinyLeaderboardResults focusMembershipId false */

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *DestinyHistoricalStatsDestinyLeaderboardResults) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Indicate the character ID of the character that is the focal point of the provided leaderboards. May be null, in which case any character from the focus membership can appear in the provided leaderboards.
		FocusCharacterID int64 `json:"focusCharacterId,omitempty"`

		// Indicate the membership ID of the account that is the focal point of the provided leaderboards.
		FocusMembershipID int64 `json:"focusMembershipId,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DestinyHistoricalStatsDestinyLeaderboardResults

	rcv.FocusCharacterID = stage1.FocusCharacterID

	rcv.FocusMembershipID = stage1.FocusMembershipID

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "focusCharacterId")

	delete(stage2, "focusMembershipId")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]DestinyHistoricalStatsDestinyLeaderboardResultsAdditionalProperties)
		for k, v := range stage2 {
			var toadd DestinyHistoricalStatsDestinyLeaderboardResultsAdditionalProperties
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.DestinyHistoricalStatsDestinyLeaderboardResults = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m DestinyHistoricalStatsDestinyLeaderboardResults) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Indicate the character ID of the character that is the focal point of the provided leaderboards. May be null, in which case any character from the focus membership can appear in the provided leaderboards.
		FocusCharacterID int64 `json:"focusCharacterId,omitempty"`

		// Indicate the membership ID of the account that is the focal point of the provided leaderboards.
		FocusMembershipID int64 `json:"focusMembershipId,omitempty"`
	}

	stage1.FocusCharacterID = m.FocusCharacterID

	stage1.FocusMembershipID = m.FocusMembershipID

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.DestinyHistoricalStatsDestinyLeaderboardResults) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.DestinyHistoricalStatsDestinyLeaderboardResults)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this destiny historical stats destiny leaderboard results
func (m *DestinyHistoricalStatsDestinyLeaderboardResults) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyLeaderboardResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyHistoricalStatsDestinyLeaderboardResults) UnmarshalBinary(b []byte) error {
	var res DestinyHistoricalStatsDestinyLeaderboardResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
