// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupMember groups v2 group member
// swagger:model GroupsV2.GroupMember

type GroupsV2GroupMember struct {

	// bungie net user info
	BungieNetUserInfo *UserUserInfoCard `json:"bungieNetUserInfo,omitempty"`

	// destiny user info
	DestinyUserInfo *UserUserInfoCard `json:"destinyUserInfo,omitempty"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`

	// is online
	IsOnline bool `json:"isOnline,omitempty"`

	// join date
	JoinDate strfmt.DateTime `json:"joinDate,omitempty"`

	// member type
	MemberType GroupsV2RuntimeGroupMemberType `json:"memberType,omitempty"`
}

/* polymorph GroupsV2.GroupMember bungieNetUserInfo false */

/* polymorph GroupsV2.GroupMember destinyUserInfo false */

/* polymorph GroupsV2.GroupMember groupId false */

/* polymorph GroupsV2.GroupMember isOnline false */

/* polymorph GroupsV2.GroupMember joinDate false */

/* polymorph GroupsV2.GroupMember memberType false */

// Validate validates this groups v2 group member
func (m *GroupsV2GroupMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBungieNetUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDestinyUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemberType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupMember) validateBungieNetUserInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.BungieNetUserInfo) { // not required
		return nil
	}

	if m.BungieNetUserInfo != nil {

		if err := m.BungieNetUserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bungieNetUserInfo")
			}
			return err
		}
	}

	return nil
}

func (m *GroupsV2GroupMember) validateDestinyUserInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinyUserInfo) { // not required
		return nil
	}

	if m.DestinyUserInfo != nil {

		if err := m.DestinyUserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinyUserInfo")
			}
			return err
		}
	}

	return nil
}

func (m *GroupsV2GroupMember) validateMemberType(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	if err := m.MemberType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("memberType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupMember) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
