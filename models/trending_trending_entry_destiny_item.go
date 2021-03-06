// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TrendingTrendingEntryDestinyItem trending trending entry destiny item
// swagger:model Trending.TrendingEntryDestinyItem

type TrendingTrendingEntryDestinyItem struct {

	// item hash
	ItemHash uint32 `json:"itemHash,omitempty"`
}

/* polymorph Trending.TrendingEntryDestinyItem itemHash false */

// Validate validates this trending trending entry destiny item
func (m *TrendingTrendingEntryDestinyItem) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TrendingTrendingEntryDestinyItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrendingTrendingEntryDestinyItem) UnmarshalBinary(b []byte) error {
	var res TrendingTrendingEntryDestinyItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
