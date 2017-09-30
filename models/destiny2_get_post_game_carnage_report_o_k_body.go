// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Destiny2GetPostGameCarnageReportOKBody destiny2 get post game carnage report o k body
// swagger:model destiny2GetPostGameCarnageReportOKBody

type Destiny2GetPostGameCarnageReportOKBody struct {

	// error code
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`

	// error status
	ErrorStatus string `json:"ErrorStatus,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// message data
	MessageData map[string]string `json:"MessageData,omitempty"`

	// response
	Response *DestinyHistoricalStatsDestinyPostGameCarnageReportData `json:"Response,omitempty"`

	// throttle seconds
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
}

/* polymorph destiny2GetPostGameCarnageReportOKBody ErrorCode false */

/* polymorph destiny2GetPostGameCarnageReportOKBody ErrorStatus false */

/* polymorph destiny2GetPostGameCarnageReportOKBody Message false */

/* polymorph destiny2GetPostGameCarnageReportOKBody MessageData false */

/* polymorph destiny2GetPostGameCarnageReportOKBody Response false */

/* polymorph destiny2GetPostGameCarnageReportOKBody ThrottleSeconds false */

// Validate validates this destiny2 get post game carnage report o k body
func (m *Destiny2GetPostGameCarnageReportOKBody) Validate(formats strfmt.Registry) error {
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

func (m *Destiny2GetPostGameCarnageReportOKBody) validateErrorCode(formats strfmt.Registry) error {

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

func (m *Destiny2GetPostGameCarnageReportOKBody) validateResponse(formats strfmt.Registry) error {

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
func (m *Destiny2GetPostGameCarnageReportOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Destiny2GetPostGameCarnageReportOKBody) UnmarshalBinary(b []byte) error {
	var res Destiny2GetPostGameCarnageReportOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}