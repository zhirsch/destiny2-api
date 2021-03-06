// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserSearchUsersOKBody user search users o k body
// swagger:model userSearchUsersOKBody

type UserSearchUsersOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response UserSearchUsersOKBodyResponse `json:"Response"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph userSearchUsersOKBody ErrorCode false */

/* polymorph userSearchUsersOKBody ErrorStatus false */

/* polymorph userSearchUsersOKBody Message false */

/* polymorph userSearchUsersOKBody MessageData false */

/* polymorph userSearchUsersOKBody Response false */

/* polymorph userSearchUsersOKBody ThrottleSeconds false */

// Validate validates this user search users o k body
func (m *UserSearchUsersOKBody) Validate(formats strfmt.Registry) error {
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

func (m *UserSearchUsersOKBody) validateErrorCode(formats strfmt.Registry) error {

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
func (m *UserSearchUsersOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSearchUsersOKBody) UnmarshalBinary(b []byte) error {
	var res UserSearchUsersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
