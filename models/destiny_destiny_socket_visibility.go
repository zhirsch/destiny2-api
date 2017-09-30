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

// DestinyDestinySocketVisibility destiny destiny socket visibility
// swagger:model Destiny.DestinySocketVisibility

type DestinyDestinySocketVisibility int32

// for schema
var destinyDestinySocketVisibilityEnum []interface{}

func init() {
	var res []DestinyDestinySocketVisibility
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinySocketVisibilityEnum = append(destinyDestinySocketVisibilityEnum, v)
	}
}

func (m DestinyDestinySocketVisibility) validateDestinyDestinySocketVisibilityEnum(path, location string, value DestinyDestinySocketVisibility) error {
	if err := validate.Enum(path, location, value, destinyDestinySocketVisibilityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny socket visibility
func (m DestinyDestinySocketVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinySocketVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
