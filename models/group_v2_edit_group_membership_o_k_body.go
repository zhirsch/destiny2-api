// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupV2EditGroupMembershipOKBody group v2 edit group membership o k body
// swagger:model groupV2EditGroupMembershipOKBody

type GroupV2EditGroupMembershipOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response int32 `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph groupV2EditGroupMembershipOKBody ErrorCode false */

/* polymorph groupV2EditGroupMembershipOKBody ErrorStatus false */

/* polymorph groupV2EditGroupMembershipOKBody Message false */

/* polymorph groupV2EditGroupMembershipOKBody MessageData false */

/* polymorph groupV2EditGroupMembershipOKBody Response false */

/* polymorph groupV2EditGroupMembershipOKBody ThrottleSeconds false */

// Validate validates this group v2 edit group membership o k body
func (m *GroupV2EditGroupMembershipOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupV2EditGroupMembershipOKBody) validateErrorCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorCode) { // not required
		return nil
	}

	if err := m.ErrorCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ErrorCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupV2EditGroupMembershipOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupV2EditGroupMembershipOKBody) UnmarshalBinary(b []byte) error {
	var res GroupV2EditGroupMembershipOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
