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

// DestinyDestinyActivityNavPointType destiny destiny activity nav point type
// swagger:model Destiny.DestinyActivityNavPointType

type DestinyDestinyActivityNavPointType int32

// for schema
var destinyDestinyActivityNavPointTypeEnum []interface{}

func init() {
	var res []DestinyDestinyActivityNavPointType
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyActivityNavPointTypeEnum = append(destinyDestinyActivityNavPointTypeEnum, v)
	}
}

func (m DestinyDestinyActivityNavPointType) validateDestinyDestinyActivityNavPointTypeEnum(path, location string, value DestinyDestinyActivityNavPointType) error {
	if err := validate.Enum(path, location, value, destinyDestinyActivityNavPointTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny activity nav point type
func (m DestinyDestinyActivityNavPointType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyActivityNavPointTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}