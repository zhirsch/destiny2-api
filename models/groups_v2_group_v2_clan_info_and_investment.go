// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupV2ClanInfoAndInvestment The same as GroupV2ClanInfo, but includes any investment data.
// swagger:model GroupsV2.GroupV2ClanInfoAndInvestment

type GroupsV2GroupV2ClanInfoAndInvestment struct {

	// clan banner data
	ClanBannerData *GroupsV2ClanBanner `json:"clanBannerData,omitempty"`

	// clan callsign
	ClanCallsign string `json:"clanCallsign,omitempty"`

	// d2 clan progressions
	D2ClanProgressions GroupsV2GroupV2ClanInfoAndInvestmentD2ClanProgressions `json:"d2ClanProgressions,omitempty"`
}

/* polymorph GroupsV2.GroupV2ClanInfoAndInvestment clanBannerData false */

/* polymorph GroupsV2.GroupV2ClanInfoAndInvestment clanCallsign false */

/* polymorph GroupsV2.GroupV2ClanInfoAndInvestment d2ClanProgressions false */

// Validate validates this groups v2 group v2 clan info and investment
func (m *GroupsV2GroupV2ClanInfoAndInvestment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClanBannerData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupV2ClanInfoAndInvestment) validateClanBannerData(formats strfmt.Registry) error {

	if swag.IsZero(m.ClanBannerData) { // not required
		return nil
	}

	if m.ClanBannerData != nil {

		if err := m.ClanBannerData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clanBannerData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupV2ClanInfoAndInvestment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupV2ClanInfoAndInvestment) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupV2ClanInfoAndInvestment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
