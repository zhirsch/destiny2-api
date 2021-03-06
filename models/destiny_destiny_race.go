// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DestinyDestinyRace destiny destiny race
// swagger:model Destiny.DestinyRace

type DestinyDestinyRace int32

// for schema
var destinyDestinyRaceEnum []interface{}

func init() {
	var res []DestinyDestinyRace
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyRaceEnum = append(destinyDestinyRaceEnum, v)
	}
}

func (m DestinyDestinyRace) validateDestinyDestinyRaceEnum(path, location string, value DestinyDestinyRace) error {
	if err := validate.Enum(path, location, value, destinyDestinyRaceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny race
func (m DestinyDestinyRace) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyRaceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
