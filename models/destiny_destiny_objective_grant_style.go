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

// DestinyDestinyObjectiveGrantStyle Some Objectives provide perks, generally as part of providing some kind of interesting modifier for a Challenge or Quest. This indicates when the Perk is granted.
// swagger:model Destiny.DestinyObjectiveGrantStyle

type DestinyDestinyObjectiveGrantStyle int32

// for schema
var destinyDestinyObjectiveGrantStyleEnum []interface{}

func init() {
	var res []DestinyDestinyObjectiveGrantStyle
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyObjectiveGrantStyleEnum = append(destinyDestinyObjectiveGrantStyleEnum, v)
	}
}

func (m DestinyDestinyObjectiveGrantStyle) validateDestinyDestinyObjectiveGrantStyleEnum(path, location string, value DestinyDestinyObjectiveGrantStyle) error {
	if err := validate.Enum(path, location, value, destinyDestinyObjectiveGrantStyleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny objective grant style
func (m DestinyDestinyObjectiveGrantStyle) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyObjectiveGrantStyleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}