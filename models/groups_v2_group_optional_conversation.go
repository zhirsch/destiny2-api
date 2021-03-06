// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupOptionalConversation groups v2 group optional conversation
// swagger:model GroupsV2.GroupOptionalConversation

type GroupsV2GroupOptionalConversation struct {

	// chat enabled
	ChatEnabled bool `json:"chatEnabled,omitempty"`

	// chat name
	ChatName string `json:"chatName,omitempty"`

	// chat security
	ChatSecurity GroupsV2ChatSecuritySetting `json:"chatSecurity,omitempty"`

	// conversation Id
	ConversationID int64 `json:"conversationId,omitempty"`

	// group Id
	GroupID int64 `json:"groupId,omitempty"`
}

/* polymorph GroupsV2.GroupOptionalConversation chatEnabled false */

/* polymorph GroupsV2.GroupOptionalConversation chatName false */

/* polymorph GroupsV2.GroupOptionalConversation chatSecurity false */

/* polymorph GroupsV2.GroupOptionalConversation conversationId false */

/* polymorph GroupsV2.GroupOptionalConversation groupId false */

// Validate validates this groups v2 group optional conversation
func (m *GroupsV2GroupOptionalConversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChatSecurity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupsV2GroupOptionalConversation) validateChatSecurity(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *GroupsV2GroupOptionalConversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupOptionalConversation) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupOptionalConversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
