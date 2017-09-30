// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyRequestsActionsDestinyItemSetActionRequest destiny requests actions destiny item set action request
// swagger:model Destiny.Requests.Actions.DestinyItemSetActionRequest

type DestinyRequestsActionsDestinyItemSetActionRequest struct {

	// character Id
	CharacterID int64 `json:"characterId,omitempty"`

	// item ids
	ItemIds []int64 `json:"itemIds"`

	// membership type
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`
}

/* polymorph Destiny.Requests.Actions.DestinyItemSetActionRequest characterId false */

/* polymorph Destiny.Requests.Actions.DestinyItemSetActionRequest itemIds false */

/* polymorph Destiny.Requests.Actions.DestinyItemSetActionRequest membershipType false */

// Validate validates this destiny requests actions destiny item set action request
func (m *DestinyRequestsActionsDestinyItemSetActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembershipType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyRequestsActionsDestinyItemSetActionRequest) validateItemIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemIds) { // not required
		return nil
	}

	return nil
}

func (m *DestinyRequestsActionsDestinyItemSetActionRequest) validateMembershipType(formats strfmt.Registry) error {

	if swag.IsZero(m.MembershipType) { // not required
		return nil
	}

	if err := m.MembershipType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("membershipType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyRequestsActionsDestinyItemSetActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyRequestsActionsDestinyItemSetActionRequest) UnmarshalBinary(b []byte) error {
	var res DestinyRequestsActionsDestinyItemSetActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
