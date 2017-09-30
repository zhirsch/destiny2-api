// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SearchResultOfGroupMember search result of group member
// swagger:model SearchResultOfGroupMember

type SearchResultOfGroupMember struct {

	// has more
	HasMore bool `json:"hasMore,omitempty"`

	// query
	Query *QueriesPagedQuery `json:"query,omitempty"`

	// replacement continuation token
	ReplacementContinuationToken string `json:"replacementContinuationToken,omitempty"`

	// results
	Results SearchResultOfGroupMemberResults `json:"results"`

	// total results
	TotalResults int32 `json:"totalResults,omitempty"`

	// If useTotalResults is true, then totalResults represents an accurate count.
	// If False, it does not, and may be estimated/only the size of the current page.
	// Either way, you should probably always only trust hasMore.
	// This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one.
	UseTotalResults bool `json:"useTotalResults,omitempty"`
}

/* polymorph SearchResultOfGroupMember hasMore false */

/* polymorph SearchResultOfGroupMember query false */

/* polymorph SearchResultOfGroupMember replacementContinuationToken false */

/* polymorph SearchResultOfGroupMember results false */

/* polymorph SearchResultOfGroupMember totalResults false */

/* polymorph SearchResultOfGroupMember useTotalResults false */

// Validate validates this search result of group member
func (m *SearchResultOfGroupMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResultOfGroupMember) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {

		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultOfGroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultOfGroupMember) UnmarshalBinary(b []byte) error {
	var res SearchResultOfGroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
