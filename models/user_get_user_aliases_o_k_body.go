// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserGetUserAliasesOKBody user get user aliases o k body
// swagger:model userGetUserAliasesOKBody

type UserGetUserAliasesOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response UserGetUserAliasesOKBodyResponse `json:"Response"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph userGetUserAliasesOKBody ErrorCode false */

/* polymorph userGetUserAliasesOKBody ErrorStatus false */

/* polymorph userGetUserAliasesOKBody Message false */

/* polymorph userGetUserAliasesOKBody MessageData false */

/* polymorph userGetUserAliasesOKBody Response false */

/* polymorph userGetUserAliasesOKBody ThrottleSeconds false */

// Validate validates this user get user aliases o k body
func (m *UserGetUserAliasesOKBody) Validate(formats strfmt.Registry) error {
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

func (m *UserGetUserAliasesOKBody) validateErrorCode(formats strfmt.Registry) error {

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
func (m *UserGetUserAliasesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGetUserAliasesOKBody) UnmarshalBinary(b []byte) error {
	var res UserGetUserAliasesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}