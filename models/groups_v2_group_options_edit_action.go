// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GroupsV2GroupOptionsEditAction groups v2 group options edit action
// swagger:model GroupsV2.GroupOptionsEditAction

type GroupsV2GroupOptionsEditAction struct {

	// Minimum Member Level allowed to host guided games
	// Always Allowed: Founder, Acting Founder, Admin
	// Allowed Overrides: None, Member, Beginner
	// Default is Member for clans, None for groups, although this means nothing for groups.
	HostGUIDEDGamePermissionOverride int32 `json:"HostGuidedGamePermissionOverride,omitempty"`

	// Minimum Member Level allowed to invite new members to group
	// Always Allowed: Founder, Acting Founder
	// True means admins have this power, false means they don't
	// Default is false for clans, true for groups.
	InvitePermissionOverride bool `json:"InvitePermissionOverride,omitempty"`

	// Level to join a member at when accepting an invite, application, or joining an open clan
	// Default is Beginner.
	JoinLevel int32 `json:"JoinLevel,omitempty"`

	// Minimum Member Level allowed to update banner
	// Always Allowed: Founder, Acting Founder
	// True means admins have this power, false means they don't
	// Default is false for clans, true for groups.
	UpdateBannerPermissionOverride bool `json:"UpdateBannerPermissionOverride,omitempty"`

	// Minimum Member Level allowed to update group culture
	// Always Allowed: Founder, Acting Founder
	// True means admins have this power, false means they don't
	// Default is false for clans, true for groups.
	UpdateCulturePermissionOverride bool `json:"UpdateCulturePermissionOverride,omitempty"`
}

/* polymorph GroupsV2.GroupOptionsEditAction HostGuidedGamePermissionOverride false */

/* polymorph GroupsV2.GroupOptionsEditAction InvitePermissionOverride false */

/* polymorph GroupsV2.GroupOptionsEditAction JoinLevel false */

/* polymorph GroupsV2.GroupOptionsEditAction UpdateBannerPermissionOverride false */

/* polymorph GroupsV2.GroupOptionsEditAction UpdateCulturePermissionOverride false */

// Validate validates this groups v2 group options edit action
func (m *GroupsV2GroupOptionsEditAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostGUIDEDGamePermissionOverride(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJoinLevel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var groupsV2GroupOptionsEditActionTypeHostGUIDEDGamePermissionOverridePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupsV2GroupOptionsEditActionTypeHostGUIDEDGamePermissionOverridePropEnum = append(groupsV2GroupOptionsEditActionTypeHostGUIDEDGamePermissionOverridePropEnum, v)
	}
}

// prop value enum
func (m *GroupsV2GroupOptionsEditAction) validateHostGUIDEDGamePermissionOverrideEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, groupsV2GroupOptionsEditActionTypeHostGUIDEDGamePermissionOverridePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GroupsV2GroupOptionsEditAction) validateHostGUIDEDGamePermissionOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGUIDEDGamePermissionOverride) { // not required
		return nil
	}

	// value enum
	if err := m.validateHostGUIDEDGamePermissionOverrideEnum("HostGuidedGamePermissionOverride", "body", m.HostGUIDEDGamePermissionOverride); err != nil {
		return err
	}

	return nil
}

var groupsV2GroupOptionsEditActionTypeJoinLevelPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupsV2GroupOptionsEditActionTypeJoinLevelPropEnum = append(groupsV2GroupOptionsEditActionTypeJoinLevelPropEnum, v)
	}
}

// prop value enum
func (m *GroupsV2GroupOptionsEditAction) validateJoinLevelEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, groupsV2GroupOptionsEditActionTypeJoinLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GroupsV2GroupOptionsEditAction) validateJoinLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.JoinLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateJoinLevelEnum("JoinLevel", "body", m.JoinLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupOptionsEditAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupOptionsEditAction) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupOptionsEditAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
