// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserGeneralUser user general user
// swagger:model User.GeneralUser

type UserGeneralUser struct {

	// about
	About string `json:"about,omitempty"`

	// blizzard display name
	BlizzardDisplayName string `json:"blizzardDisplayName,omitempty"`

	// context
	Context *UserUserToUserContext `json:"context,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// fb display name
	FbDisplayName string `json:"fbDisplayName,omitempty"`

	// first access
	FirstAccess strfmt.DateTime `json:"firstAccess,omitempty"`

	// is deleted
	IsDeleted bool `json:"isDeleted,omitempty"`

	// last ban report Id
	LastBanReportID int64 `json:"lastBanReportId,omitempty"`

	// last update
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// legacy portal UID
	LegacyPortalUID int64 `json:"legacyPortalUID,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// locale inherit default
	LocaleInheritDefault bool `json:"localeInheritDefault,omitempty"`

	// membership Id
	MembershipID int64 `json:"membershipId,omitempty"`

	// normalized name
	NormalizedName string `json:"normalizedName,omitempty"`

	// profile ban expire
	ProfileBanExpire strfmt.DateTime `json:"profileBanExpire,omitempty"`

	// profile picture
	ProfilePicture int32 `json:"profilePicture,omitempty"`

	// profile picture path
	ProfilePicturePath string `json:"profilePicturePath,omitempty"`

	// profile picture wide path
	ProfilePictureWidePath string `json:"profilePictureWidePath,omitempty"`

	// profile theme
	ProfileTheme int32 `json:"profileTheme,omitempty"`

	// profile theme name
	ProfileThemeName string `json:"profileThemeName,omitempty"`

	// psn display name
	PsnDisplayName string `json:"psnDisplayName,omitempty"`

	// show activity
	ShowActivity bool `json:"showActivity,omitempty"`

	// show group messaging
	ShowGroupMessaging bool `json:"showGroupMessaging,omitempty"`

	// status date
	StatusDate strfmt.DateTime `json:"statusDate,omitempty"`

	// status text
	StatusText string `json:"statusText,omitempty"`

	// success message flags
	SuccessMessageFlags int64 `json:"successMessageFlags,omitempty"`

	// unique name
	UniqueName string `json:"uniqueName,omitempty"`

	// user title
	UserTitle int32 `json:"userTitle,omitempty"`

	// user title display
	UserTitleDisplay string `json:"userTitleDisplay,omitempty"`

	// xbox display name
	XboxDisplayName string `json:"xboxDisplayName,omitempty"`
}

/* polymorph User.GeneralUser about false */

/* polymorph User.GeneralUser blizzardDisplayName false */

/* polymorph User.GeneralUser context false */

/* polymorph User.GeneralUser displayName false */

/* polymorph User.GeneralUser fbDisplayName false */

/* polymorph User.GeneralUser firstAccess false */

/* polymorph User.GeneralUser isDeleted false */

/* polymorph User.GeneralUser lastBanReportId false */

/* polymorph User.GeneralUser lastUpdate false */

/* polymorph User.GeneralUser legacyPortalUID false */

/* polymorph User.GeneralUser locale false */

/* polymorph User.GeneralUser localeInheritDefault false */

/* polymorph User.GeneralUser membershipId false */

/* polymorph User.GeneralUser normalizedName false */

/* polymorph User.GeneralUser profileBanExpire false */

/* polymorph User.GeneralUser profilePicture false */

/* polymorph User.GeneralUser profilePicturePath false */

/* polymorph User.GeneralUser profilePictureWidePath false */

/* polymorph User.GeneralUser profileTheme false */

/* polymorph User.GeneralUser profileThemeName false */

/* polymorph User.GeneralUser psnDisplayName false */

/* polymorph User.GeneralUser showActivity false */

/* polymorph User.GeneralUser showGroupMessaging false */

/* polymorph User.GeneralUser statusDate false */

/* polymorph User.GeneralUser statusText false */

/* polymorph User.GeneralUser successMessageFlags false */

/* polymorph User.GeneralUser uniqueName false */

/* polymorph User.GeneralUser userTitle false */

/* polymorph User.GeneralUser userTitleDisplay false */

/* polymorph User.GeneralUser xboxDisplayName false */

// Validate validates this user general user
func (m *UserGeneralUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGeneralUser) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {

		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserGeneralUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGeneralUser) UnmarshalBinary(b []byte) error {
	var res UserGeneralUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}