// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommunityContentGetStreamingStatusForMemberOKBody community content get streaming status for member o k body
// swagger:model communityContentGetStreamingStatusForMemberOKBody

type CommunityContentGetStreamingStatusForMemberOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *CommunityCommunityLiveStatus `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph communityContentGetStreamingStatusForMemberOKBody ErrorCode false */

/* polymorph communityContentGetStreamingStatusForMemberOKBody ErrorStatus false */

/* polymorph communityContentGetStreamingStatusForMemberOKBody Message false */

/* polymorph communityContentGetStreamingStatusForMemberOKBody MessageData false */

/* polymorph communityContentGetStreamingStatusForMemberOKBody Response false */

/* polymorph communityContentGetStreamingStatusForMemberOKBody ThrottleSeconds false */

// Validate validates this community content get streaming status for member o k body
func (m *CommunityContentGetStreamingStatusForMemberOKBody) Validate(formats strfmt.Registry) error {
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

func (m *CommunityContentGetStreamingStatusForMemberOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *CommunityContentGetStreamingStatusForMemberOKBody) validateResponse(formats strfmt.Registry) error {

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
func (m *CommunityContentGetStreamingStatusForMemberOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommunityContentGetStreamingStatusForMemberOKBody) UnmarshalBinary(b []byte) error {
	var res CommunityContentGetStreamingStatusForMemberOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}