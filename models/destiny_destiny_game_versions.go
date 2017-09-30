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

// DestinyDestinyGameVersions A flags enumeration indicating the versions of the game that a given user has purchased.
// swagger:model Destiny.DestinyGameVersions

type DestinyDestinyGameVersions int32

// for schema
var destinyDestinyGameVersionsEnum []interface{}

func init() {
	var res []DestinyDestinyGameVersions
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyGameVersionsEnum = append(destinyDestinyGameVersionsEnum, v)
	}
}

func (m DestinyDestinyGameVersions) validateDestinyDestinyGameVersionsEnum(path, location string, value DestinyDestinyGameVersions) error {
	if err := validate.Enum(path, location, value, destinyDestinyGameVersionsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny game versions
func (m DestinyDestinyGameVersions) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyGameVersionsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
