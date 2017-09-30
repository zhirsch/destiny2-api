// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupV2GetPotentialGroupsForMemberOKBody group v2 get potential groups for member o k body
// swagger:model groupV2GetPotentialGroupsForMemberOKBody

type GroupV2GetPotentialGroupsForMemberOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *GroupsV2GroupPotentialMembershipSearchResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph groupV2GetPotentialGroupsForMemberOKBody ErrorCode false */

/* polymorph groupV2GetPotentialGroupsForMemberOKBody ErrorStatus false */

/* polymorph groupV2GetPotentialGroupsForMemberOKBody Message false */

/* polymorph groupV2GetPotentialGroupsForMemberOKBody MessageData false */

/* polymorph groupV2GetPotentialGroupsForMemberOKBody Response false */

/* polymorph groupV2GetPotentialGroupsForMemberOKBody ThrottleSeconds false */

// Validate validates this group v2 get potential groups for member o k body
func (m *GroupV2GetPotentialGroupsForMemberOKBody) Validate(formats strfmt.Registry) error {
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

func (m *GroupV2GetPotentialGroupsForMemberOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *GroupV2GetPotentialGroupsForMemberOKBody) validateResponse(formats strfmt.Registry) error {

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
func (m *GroupV2GetPotentialGroupsForMemberOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupV2GetPotentialGroupsForMemberOKBody) UnmarshalBinary(b []byte) error {
	var res GroupV2GetPotentialGroupsForMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
