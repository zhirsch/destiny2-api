// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupV2GetGroupsForMemberOKBody group v2 get groups for member o k body
// swagger:model groupV2GetGroupsForMemberOKBody

type GroupV2GetGroupsForMemberOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *GroupsV2GroupMembershipSearchResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph groupV2GetGroupsForMemberOKBody ErrorCode false */

/* polymorph groupV2GetGroupsForMemberOKBody ErrorStatus false */

/* polymorph groupV2GetGroupsForMemberOKBody Message false */

/* polymorph groupV2GetGroupsForMemberOKBody MessageData false */

/* polymorph groupV2GetGroupsForMemberOKBody Response false */

/* polymorph groupV2GetGroupsForMemberOKBody ThrottleSeconds false */

// Validate validates this group v2 get groups for member o k body
func (m *GroupV2GetGroupsForMemberOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupV2GetGroupsForMemberOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *GroupV2GetGroupsForMemberOKBody) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {

		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupV2GetGroupsForMemberOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupV2GetGroupsForMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GroupV2GetGroupsForMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
