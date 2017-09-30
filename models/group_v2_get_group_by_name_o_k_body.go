// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupV2GetGroupByNameOKBody group v2 get group by name o k body
// swagger:model groupV2GetGroupByNameOKBody

type GroupV2GetGroupByNameOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *GroupsV2GroupResponse `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph groupV2GetGroupByNameOKBody ErrorCode false */

/* polymorph groupV2GetGroupByNameOKBody ErrorStatus false */

/* polymorph groupV2GetGroupByNameOKBody Message false */

/* polymorph groupV2GetGroupByNameOKBody MessageData false */

/* polymorph groupV2GetGroupByNameOKBody Response false */

/* polymorph groupV2GetGroupByNameOKBody ThrottleSeconds false */

// Validate validates this group v2 get group by name o k body
func (m *GroupV2GetGroupByNameOKBody) Validate(formats strfmt.Registry) error {
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

func (m *GroupV2GetGroupByNameOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *GroupV2GetGroupByNameOKBody) validateResponse(formats strfmt.Registry) error {

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
func (m *GroupV2GetGroupByNameOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupV2GetGroupByNameOKBody) UnmarshalBinary(b []byte) error {
	var res GroupV2GetGroupByNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
