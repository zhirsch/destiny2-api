// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDestinyUnlockStatus Indicates the status of an "Unlock Flag" on a Character or Profile.
// These are individual bits of state that can be either set or not set, and sometimes provide interesting human-readable information in their related DestinyUnlockDefinition.
// swagger:model Destiny.DestinyUnlockStatus

type DestinyDestinyUnlockStatus struct {

	// Whether the unlock flag is set.
	IsSet bool `json:"isSet,omitempty"`

	// The hash identifier for the Unlock Flag. Use to lookup DestinyUnlockDefinition for static data. Not all unlocks have human readable data - in fact, most don't. But when they do, it can be very useful to show. Even if they don't have human readable data, you might be able to infer the meaning of an unlock flag with a bit of experimentation...
	UnlockHash uint32 `json:"unlockHash,omitempty"`
}

/* polymorph Destiny.DestinyUnlockStatus isSet false */

/* polymorph Destiny.DestinyUnlockStatus unlockHash false */

// Validate validates this destiny destiny unlock status
func (m *DestinyDestinyUnlockStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDestinyUnlockStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDestinyUnlockStatus) UnmarshalBinary(b []byte) error {
	var res DestinyDestinyUnlockStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
