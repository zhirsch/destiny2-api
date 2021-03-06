// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsAnimationsDestinyAnimationReference destiny definitions animations destiny animation reference
// swagger:model Destiny.Definitions.Animations.DestinyAnimationReference

type DestinyDefinitionsAnimationsDestinyAnimationReference struct {

	// anim identifier
	AnimIdentifier string `json:"animIdentifier,omitempty"`

	// anim name
	AnimName string `json:"animName,omitempty"`

	// path
	Path string `json:"path,omitempty"`
}

/* polymorph Destiny.Definitions.Animations.DestinyAnimationReference animIdentifier false */

/* polymorph Destiny.Definitions.Animations.DestinyAnimationReference animName false */

/* polymorph Destiny.Definitions.Animations.DestinyAnimationReference path false */

// Validate validates this destiny definitions animations destiny animation reference
func (m *DestinyDefinitionsAnimationsDestinyAnimationReference) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsAnimationsDestinyAnimationReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsAnimationsDestinyAnimationReference) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsAnimationsDestinyAnimationReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
