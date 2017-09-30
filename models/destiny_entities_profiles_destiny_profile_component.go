// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesProfilesDestinyProfileComponent The most essential summary information about a Profile (in Destiny 1, we called these "Accounts").
// swagger:model Destiny.Entities.Profiles.DestinyProfileComponent

type DestinyEntitiesProfilesDestinyProfileComponent struct {

	// A list of the character IDs, for further querying on your part.
	CharacterIds []int64 `json:"characterIds"`

	// The last time the user played with any character on this Profile.
	DateLastPlayed strfmt.DateTime `json:"dateLastPlayed,omitempty"`

	// If you need to render the Profile (their platform name, icon, etc...) somewhere, this property contains that information.
	UserInfo *UserUserInfoCard `json:"userInfo,omitempty"`

	// If you want to know what expansions they own, this will contain that data.
	VersionsOwned DestinyDestinyGameVersions `json:"versionsOwned,omitempty"`
}

/* polymorph Destiny.Entities.Profiles.DestinyProfileComponent characterIds false */

/* polymorph Destiny.Entities.Profiles.DestinyProfileComponent dateLastPlayed false */

/* polymorph Destiny.Entities.Profiles.DestinyProfileComponent userInfo false */

/* polymorph Destiny.Entities.Profiles.DestinyProfileComponent versionsOwned false */

// Validate validates this destiny entities profiles destiny profile component
func (m *DestinyEntitiesProfilesDestinyProfileComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersionsOwned(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyEntitiesProfilesDestinyProfileComponent) validateCharacterIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CharacterIds) { // not required
		return nil
	}

	return nil
}

func (m *DestinyEntitiesProfilesDestinyProfileComponent) validateUserInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {

		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyEntitiesProfilesDestinyProfileComponent) validateVersionsOwned(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionsOwned) { // not required
		return nil
	}

	if err := m.VersionsOwned.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("versionsOwned")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesProfilesDestinyProfileComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesProfilesDestinyProfileComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesProfilesDestinyProfileComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}