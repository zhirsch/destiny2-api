// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupAction groups v2 group action
// swagger:model GroupsV2.GroupAction

type GroupsV2GroupAction struct {

	// about
	About string `json:"about,omitempty"`

	// allow chat
	AllowChat bool `json:"allowChat,omitempty"`

	// avatar image index
	AvatarImageIndex int32 `json:"avatarImageIndex,omitempty"`

	// callsign
	Callsign string `json:"callsign,omitempty"`

	// chat security
	ChatSecurity GroupsV2ChatSecuritySetting `json:"chatSecurity,omitempty"`

	// Type of group, either Bungie.net hosted group, or a game services hosted clan.
	GroupType GroupsV2GroupType `json:"groupType,omitempty"`

	// homepage
	Homepage GroupsV2GroupHomepage `json:"homepage,omitempty"`

	// is default post alliance
	IsDefaultPostAlliance bool `json:"isDefaultPostAlliance,omitempty"`

	// is default post public
	IsDefaultPostPublic bool `json:"isDefaultPostPublic,omitempty"`

	// is public
	IsPublic bool `json:"isPublic,omitempty"`

	// is public topic admin only
	IsPublicTopicAdminOnly bool `json:"isPublicTopicAdminOnly,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// membership option
	MembershipOption GroupsV2MembershipOption `json:"membershipOption,omitempty"`

	// motto
	Motto string `json:"motto,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// When operation needs a platform specific account ID for the present user, use this property. In particular, groupType of Clan requires this value to be set.
	PlatformMembershipType BungieMembershipType `json:"platformMembershipType,omitempty"`

	// tags
	Tags string `json:"tags,omitempty"`

	// theme
	Theme string `json:"theme,omitempty"`
}

/* polymorph GroupsV2.GroupAction about false */

/* polymorph GroupsV2.GroupAction allowChat false */

/* polymorph GroupsV2.GroupAction avatarImageIndex false */

/* polymorph GroupsV2.GroupAction callsign false */

/* polymorph GroupsV2.GroupAction chatSecurity false */

/* polymorph GroupsV2.GroupAction groupType false */

/* polymorph GroupsV2.GroupAction homepage false */

/* polymorph GroupsV2.GroupAction isDefaultPostAlliance false */

/* polymorph GroupsV2.GroupAction isDefaultPostPublic false */

/* polymorph GroupsV2.GroupAction isPublic false */

/* polymorph GroupsV2.GroupAction isPublicTopicAdminOnly false */

/* polymorph GroupsV2.GroupAction locale false */

/* polymorph GroupsV2.GroupAction membershipOption false */

/* polymorph GroupsV2.GroupAction motto false */

/* polymorph GroupsV2.GroupAction name false */

/* polymorph GroupsV2.GroupAction platformMembershipType false */

/* polymorph GroupsV2.GroupAction tags false */

/* polymorph GroupsV2.GroupAction theme false */

// Validate validates this groups v2 group action
func (m *GroupsV2GroupAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatSecurity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHomepage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembershipOption(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlatformMembershipType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupAction) validateChatSecurity(formats strfmt.Registry) error {

	if swag.IsZero(m.ChatSecurity) { // not required
		return nil
	}

	if err := m.ChatSecurity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("chatSecurity")
		}
		return err
	}

	return nil
}

func (m *GroupsV2GroupAction) validateGroupType(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupAction) validateHomepage(formats strfmt.Registry) error {

	if swag.IsZero(m.Homepage) { // not required
		return nil
	}

	if err := m.Homepage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("homepage")
		}
		return err
	}

	return nil
}

func (m *GroupsV2GroupAction) validateMembershipOption(formats strfmt.Registry) error {

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

func (m *GroupsV2GroupAction) validatePlatformMembershipType(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformMembershipType) { // not required
		return nil
	}

	if err := m.PlatformMembershipType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("platformMembershipType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupAction) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
