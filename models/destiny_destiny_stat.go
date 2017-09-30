// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDestinyStat Represents a stat on an item *or* Character (NOT a Historical Stat, but a physical attribute stat like Attack, Defense etc...)
// swagger:model Destiny.DestinyStat

type DestinyDestinyStat struct {

	// The highest possible value for the stat, if we were able to compute it. (I wouldn't necessarily trust this value right now. I would like to improve its calculation in later iterations of the API. Consider this a placeholder for desired future functionality)
	MaximumValue int32 `json:"maximumValue,omitempty"`

	// The hash identifier for the Stat. Use it to look up the DestinyStatDefinition for static data about the stat.
	StatHash uint32 `json:"statHash,omitempty"`

	// The current value of the Stat.
	Value int32 `json:"value,omitempty"`
}

/* polymorph Destiny.DestinyStat maximumValue false */

/* polymorph Destiny.DestinyStat statHash false */

/* polymorph Destiny.DestinyStat value false */

// Validate validates this destiny destiny stat
func (m *DestinyDestinyStat) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDestinyStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDestinyStat) UnmarshalBinary(b []byte) error {
	var res DestinyDestinyStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
