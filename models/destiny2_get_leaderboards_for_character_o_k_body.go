// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Destiny2GetLeaderboardsForCharacterOKBody destiny2 get leaderboards for character o k body
// swagger:model destiny2GetLeaderboardsForCharacterOKBody

type Destiny2GetLeaderboardsForCharacterOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response Destiny2GetLeaderboardsForCharacterOKBodyResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph destiny2GetLeaderboardsForCharacterOKBody ErrorCode false */

/* polymorph destiny2GetLeaderboardsForCharacterOKBody ErrorStatus false */

/* polymorph destiny2GetLeaderboardsForCharacterOKBody Message false */

/* polymorph destiny2GetLeaderboardsForCharacterOKBody MessageData false */

/* polymorph destiny2GetLeaderboardsForCharacterOKBody Response false */

/* polymorph destiny2GetLeaderboardsForCharacterOKBody ThrottleSeconds false */

// Validate validates this destiny2 get leaderboards for character o k body
func (m *Destiny2GetLeaderboardsForCharacterOKBody) Validate(formats strfmt.Registry) error {
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

func (m *Destiny2GetLeaderboardsForCharacterOKBody) validateErrorCode(formats strfmt.Registry) error {

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
func (m *Destiny2GetLeaderboardsForCharacterOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destiny2GetLeaderboardsForCharacterOKBody) UnmarshalBinary(b []byte) error {
	var res Destiny2GetLeaderboardsForCharacterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
