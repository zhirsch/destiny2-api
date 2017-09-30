// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TrendingTrendingEntryNews trending trending entry news
// swagger:model Trending.TrendingEntryNews

type TrendingTrendingEntryNews struct {

	// article
	Article *ContentContentItemPublicContract `json:"article,omitempty"`
}

/* polymorph Trending.TrendingEntryNews article false */

// Validate validates this trending trending entry news
func (m *TrendingTrendingEntryNews) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArticle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrendingTrendingEntryNews) validateArticle(formats strfmt.Registry) error {

	if swag.IsZero(m.Article) { // not required
		return nil
	}

	if m.Article != nil {

		if err := m.Article.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("article")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrendingTrendingEntryNews) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrendingTrendingEntryNews) UnmarshalBinary(b []byte) error {
	var res TrendingTrendingEntryNews
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
