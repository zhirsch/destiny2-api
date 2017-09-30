// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupMemberApplication groups v2 group member application
// swagger:model GroupsV2.GroupMemberApplication

type GroupsV2GroupMemberApplication struct {

	// bungie net user info
	BungieNetUserInfo *UserUserInfoCard `json:"bungieNetUserInfo,omitempty"`

	// creation date
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// destiny user info
	DestinyUserInfo *UserUserInfoCard `json:"destinyUserInfo,omitempty"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`

	// request message
	RequestMessage string `json:"requestMessage,omitempty"`

	// resolve date
	ResolveDate strfmt.DateTime `json:"resolveDate,omitempty"`

	// resolve message
	ResolveMessage string `json:"resolveMessage,omitempty"`

	// resolve state
	ResolveState GroupsV2GroupApplicationResolveState `json:"resolveState,omitempty"`

	// resolved by membership Id
	ResolvedByMembershipID int64 `json:"resolvedByMembershipId,omitempty"`
}

/* polymorph GroupsV2.GroupMemberApplication bungieNetUserInfo false */

/* polymorph GroupsV2.GroupMemberApplication creationDate false */

/* polymorph GroupsV2.GroupMemberApplication destinyUserInfo false */

/* polymorph GroupsV2.GroupMemberApplication groupId false */

/* polymorph GroupsV2.GroupMemberApplication requestMessage false */

/* polymorph GroupsV2.GroupMemberApplication resolveDate false */

/* polymorph GroupsV2.GroupMemberApplication resolveMessage false */

/* polymorph GroupsV2.GroupMemberApplication resolveState false */

/* polymorph GroupsV2.GroupMemberApplication resolvedByMembershipId false */

// Validate validates this groups v2 group member application
func (m *GroupsV2GroupMemberApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBungieNetUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDestinyUserInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResolveState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupMemberApplication) validateBungieNetUserInfo(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupMemberApplication) validateDestinyUserInfo(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupMemberApplication) validateResolveState(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolveState) { // not required
		return nil
	}

	if err := m.ResolveState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resolveState")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupMemberApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupMemberApplication) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupMemberApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
