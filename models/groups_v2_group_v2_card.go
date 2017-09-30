// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupV2Card A small infocard of group information, usually used for when a list of groups are returned
// swagger:model GroupsV2.GroupV2Card

type GroupsV2GroupV2Card struct {

	// about
	About string `json:"about,omitempty"`

	// avatar path
	AvatarPath string `json:"avatarPath,omitempty"`

	// capabilities
	Capabilities GroupsV2Capabilities `json:"capabilities,omitempty"`

	// clan info
	ClanInfo *GroupsV2GroupV2ClanInfo `json:"clanInfo,omitempty"`

	// creation date
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`

	// group type
	GroupType GroupsV2GroupType `json:"groupType,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// member count
	MemberCount int32 `json:"memberCount,omitempty"`

	// membership option
	MembershipOption GroupsV2MembershipOption `json:"membershipOption,omitempty"`

	// motto
	Motto string `json:"motto,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// theme
	Theme string `json:"theme,omitempty"`
}

/* polymorph GroupsV2.GroupV2Card about false */

/* polymorph GroupsV2.GroupV2Card avatarPath false */

/* polymorph GroupsV2.GroupV2Card capabilities false */

/* polymorph GroupsV2.GroupV2Card clanInfo false */

/* polymorph GroupsV2.GroupV2Card creationDate false */

/* polymorph GroupsV2.GroupV2Card groupId false */

/* polymorph GroupsV2.GroupV2Card groupType false */

/* polymorph GroupsV2.GroupV2Card locale false */

/* polymorph GroupsV2.GroupV2Card memberCount false */

/* polymorph GroupsV2.GroupV2Card membershipOption false */

/* polymorph GroupsV2.GroupV2Card motto false */

/* polymorph GroupsV2.GroupV2Card name false */

/* polymorph GroupsV2.GroupV2Card theme false */

// Validate validates this groups v2 group v2 card
func (m *GroupsV2GroupV2Card) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClanInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembershipOption(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupV2Card) validateCapabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	if err := m.Capabilities.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

func (m *GroupsV2GroupV2Card) validateClanInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ClanInfo) { // not required
		return nil
	}

	if m.ClanInfo != nil {

		if err := m.ClanInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clanInfo")
			}
			return err
		}
	}

	return nil
}

func (m *GroupsV2GroupV2Card) validateGroupType(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupType) { // not required
		return nil
	}

	if err := m.GroupType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("groupType")
		}
		return err
	}

	return nil
}

func (m *GroupsV2GroupV2Card) validateMembershipOption(formats strfmt.Registry) error {

	if swag.IsZero(m.MembershipOption) { // not required
		return nil
	}

	if err := m.MembershipOption.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("membershipOption")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupV2Card) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupV2Card) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupV2Card
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}