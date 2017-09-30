// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupBan groups v2 group ban
// swagger:model GroupsV2.GroupBan

type GroupsV2GroupBan struct {

	// bungie net user info
	BungieNetUserInfo *UserUserInfoCard `json:"bungieNetUserInfo,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// created by
	CreatedBy *UserUserInfoCard `json:"createdBy,omitempty"`

	// date banned
	DateBanned strfmt.DateTime `json:"dateBanned,omitempty"`

	// date expires
	DateExpires strfmt.DateTime `json:"dateExpires,omitempty"`

	// destiny user info
	DestinyUserInfo *UserUserInfoCard `json:"destinyUserInfo,omitempty"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`

	// last modified by
	LastModifiedBy *UserUserInfoCard `json:"lastModifiedBy,omitempty"`
}

/* polymorph GroupsV2.GroupBan bungieNetUserInfo false */

/* polymorph GroupsV2.GroupBan comment false */

/* polymorph GroupsV2.GroupBan createdBy false */

/* polymorph GroupsV2.GroupBan dateBanned false */

/* polymorph GroupsV2.GroupBan dateExpires false */

/* polymorph GroupsV2.GroupBan destinyUserInfo false */

/* polymorph GroupsV2.GroupBan groupId false */

/* polymorph GroupsV2.GroupBan lastModifiedBy false */

// Validate validates this groups v2 group ban
func (m *GroupsV2GroupBan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBungieNetUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDestinyUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupBan) validateBungieNetUserInfo(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupBan) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {

		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *GroupsV2GroupBan) validateDestinyUserInfo(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupBan) validateLastModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedBy) { // not required
		return nil
	}

	if m.LastModifiedBy != nil {

		if err := m.LastModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastModifiedBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupBan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupBan) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupBan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}