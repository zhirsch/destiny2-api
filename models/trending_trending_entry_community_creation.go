// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TrendingTrendingEntryCommunityCreation trending trending entry community creation
// swagger:model Trending.TrendingEntryCommunityCreation

type TrendingTrendingEntryCommunityCreation struct {

	// author
	Author string `json:"author,omitempty"`

	// author membership Id
	AuthorMembershipID int64 `json:"authorMembershipId,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// media
	Media string `json:"media,omitempty"`

	// post Id
	PostID uint64 `json:"postId,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// upvotes
	Upvotes int32 `json:"upvotes,omitempty"`
}

/* polymorph Trending.TrendingEntryCommunityCreation author false */

/* polymorph Trending.TrendingEntryCommunityCreation authorMembershipId false */

/* polymorph Trending.TrendingEntryCommunityCreation body false */

/* polymorph Trending.TrendingEntryCommunityCreation media false */

/* polymorph Trending.TrendingEntryCommunityCreation postId false */

/* polymorph Trending.TrendingEntryCommunityCreation title false */

/* polymorph Trending.TrendingEntryCommunityCreation upvotes false */

// Validate validates this trending trending entry community creation
func (m *TrendingTrendingEntryCommunityCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TrendingTrendingEntryCommunityCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrendingTrendingEntryCommunityCreation) UnmarshalBinary(b []byte) error {
	var res TrendingTrendingEntryCommunityCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}