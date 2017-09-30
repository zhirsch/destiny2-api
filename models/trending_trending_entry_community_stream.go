// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TrendingTrendingEntryCommunityStream trending trending entry community stream
// swagger:model Trending.TrendingEntryCommunityStream

type TrendingTrendingEntryCommunityStream struct {

	// image
	Image string `json:"image,omitempty"`

	// partnership identifier
	PartnershipIdentifier string `json:"partnershipIdentifier,omitempty"`

	// partnership type
	PartnershipType PartnershipsPartnershipType `json:"partnershipType,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

/* polymorph Trending.TrendingEntryCommunityStream image false */

/* polymorph Trending.TrendingEntryCommunityStream partnershipIdentifier false */

/* polymorph Trending.TrendingEntryCommunityStream partnershipType false */

/* polymorph Trending.TrendingEntryCommunityStream title false */

// Validate validates this trending trending entry community stream
func (m *TrendingTrendingEntryCommunityStream) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartnershipType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrendingTrendingEntryCommunityStream) validatePartnershipType(formats strfmt.Registry) error {

	if swag.IsZero(m.PartnershipType) { // not required
		return nil
	}

	if err := m.PartnershipType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("partnershipType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrendingTrendingEntryCommunityStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrendingTrendingEntryCommunityStream) UnmarshalBinary(b []byte) error {
	var res TrendingTrendingEntryCommunityStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
