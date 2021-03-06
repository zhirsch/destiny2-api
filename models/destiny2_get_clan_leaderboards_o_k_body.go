// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Destiny2GetClanLeaderboardsOKBody destiny2 get clan leaderboards o k body
// swagger:model destiny2GetClanLeaderboardsOKBody

type Destiny2GetClanLeaderboardsOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response Destiny2GetClanLeaderboardsOKBodyResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph destiny2GetClanLeaderboardsOKBody ErrorCode false */

/* polymorph destiny2GetClanLeaderboardsOKBody ErrorStatus false */

/* polymorph destiny2GetClanLeaderboardsOKBody Message false */

/* polymorph destiny2GetClanLeaderboardsOKBody MessageData false */

/* polymorph destiny2GetClanLeaderboardsOKBody Response false */

/* polymorph destiny2GetClanLeaderboardsOKBody ThrottleSeconds false */

// Validate validates this destiny2 get clan leaderboards o k body
func (m *Destiny2GetClanLeaderboardsOKBody) Validate(formats strfmt.Registry) error {
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

func (m *Destiny2GetClanLeaderboardsOKBody) validateErrorCode(formats strfmt.Registry) error {

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
func (m *Destiny2GetClanLeaderboardsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destiny2GetClanLeaderboardsOKBody) UnmarshalBinary(b []byte) error {
	var res Destiny2GetClanLeaderboardsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
