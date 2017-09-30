// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Destiny2GetClanAggregateStatsOKBody destiny2 get clan aggregate stats o k body
// swagger:model destiny2GetClanAggregateStatsOKBody

type Destiny2GetClanAggregateStatsOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response Destiny2GetClanAggregateStatsOKBodyResponse `json:"Response"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph destiny2GetClanAggregateStatsOKBody ErrorCode false */

/* polymorph destiny2GetClanAggregateStatsOKBody ErrorStatus false */

/* polymorph destiny2GetClanAggregateStatsOKBody Message false */

/* polymorph destiny2GetClanAggregateStatsOKBody MessageData false */

/* polymorph destiny2GetClanAggregateStatsOKBody Response false */

/* polymorph destiny2GetClanAggregateStatsOKBody ThrottleSeconds false */

// Validate validates this destiny2 get clan aggregate stats o k body
func (m *Destiny2GetClanAggregateStatsOKBody) Validate(formats strfmt.Registry) error {
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

func (m *Destiny2GetClanAggregateStatsOKBody) validateErrorCode(formats strfmt.Registry) error {

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
func (m *Destiny2GetClanAggregateStatsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destiny2GetClanAggregateStatsOKBody) UnmarshalBinary(b []byte) error {
	var res Destiny2GetClanAggregateStatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}