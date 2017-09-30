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

// DestinyTierType destiny tier type
// swagger:model Destiny.TierType

type DestinyTierType int32

// for schema
var destinyTierTypeEnum []interface{}

func init() {
	var res []DestinyTierType
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyTierTypeEnum = append(destinyTierTypeEnum, v)
	}
}

func (m DestinyTierType) validateDestinyTierTypeEnum(path, location string, value DestinyTierType) error {
	if err := validate.Enum(path, location, value, destinyTierTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny tier type
func (m DestinyTierType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyTierTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
