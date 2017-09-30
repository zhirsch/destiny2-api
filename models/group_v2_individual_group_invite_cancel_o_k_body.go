// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupV2IndividualGroupInviteCancelOKBody group v2 individual group invite cancel o k body
// swagger:model groupV2IndividualGroupInviteCancelOKBody

type GroupV2IndividualGroupInviteCancelOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *GroupsV2GroupApplicationResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph groupV2IndividualGroupInviteCancelOKBody ErrorCode false */

/* polymorph groupV2IndividualGroupInviteCancelOKBody ErrorStatus false */

/* polymorph groupV2IndividualGroupInviteCancelOKBody Message false */

/* polymorph groupV2IndividualGroupInviteCancelOKBody MessageData false */

/* polymorph groupV2IndividualGroupInviteCancelOKBody Response false */

/* polymorph groupV2IndividualGroupInviteCancelOKBody ThrottleSeconds false */

// Validate validates this group v2 individual group invite cancel o k body
func (m *GroupV2IndividualGroupInviteCancelOKBody) Validate(formats strfmt.Registry) error {
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

func (m *GroupV2IndividualGroupInviteCancelOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *GroupV2IndividualGroupInviteCancelOKBody) validateResponse(formats strfmt.Registry) error {

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
func (m *GroupV2IndividualGroupInviteCancelOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupV2IndividualGroupInviteCancelOKBody) UnmarshalBinary(b []byte) error {
	var res GroupV2IndividualGroupInviteCancelOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}