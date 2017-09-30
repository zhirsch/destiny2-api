// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserUserInfoCard This contract supplies basic information commonly used to display a minimal amount of information about a user. Take care to not add more properties here unless the property applies in all (or at least the majority) of the situations where UserInfoCard is used. Avoid adding game specific or platform specific details here. In cases where UserInfoCard is a subset of the data needed in a contract, use UserInfoCard as a property of other contracts.
// swagger:model User.UserInfoCard

type UserUserInfoCard struct {

	// Display Name the player has chosen for themselves. The display name is optional when the data type is used as input to a platform API.
	DisplayName string `json:"displayName,omitempty"`

	// URL the Icon if available.
	IconPath string `json:"iconPath,omitempty"`

	// Membership ID as they user is known in the Accounts service
	MembershipID int64 `json:"membershipId,omitempty"`

	// Type of the membership.
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`

	// A platform specific additional display name - ex: psn Real Name, bnet Unique Name, etc.
	SupplementalDisplayName string `json:"supplementalDisplayName,omitempty"`
}

/* polymorph User.UserInfoCard displayName false */

/* polymorph User.UserInfoCard iconPath false */

/* polymorph User.UserInfoCard membershipId false */

/* polymorph User.UserInfoCard membershipType false */

/* polymorph User.UserInfoCard supplementalDisplayName false */

// Validate validates this user user info card
func (m *UserUserInfoCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembershipType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUserInfoCard) validateMembershipType(formats strfmt.Registry) error {

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
func (m *UserUserInfoCard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUserInfoCard) UnmarshalBinary(b []byte) error {
	var res UserUserInfoCard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
