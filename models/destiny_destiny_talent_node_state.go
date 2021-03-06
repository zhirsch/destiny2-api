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

// DestinyDestinyTalentNodeState destiny destiny talent node state
// swagger:model Destiny.DestinyTalentNodeState

type DestinyDestinyTalentNodeState int32

// for schema
var destinyDestinyTalentNodeStateEnum []interface{}

func init() {
	var res []DestinyDestinyTalentNodeState
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7,8,9,10,11,12,13]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyTalentNodeStateEnum = append(destinyDestinyTalentNodeStateEnum, v)
	}
}

func (m DestinyDestinyTalentNodeState) validateDestinyDestinyTalentNodeStateEnum(path, location string, value DestinyDestinyTalentNodeState) error {
	if err := validate.Enum(path, location, value, destinyDestinyTalentNodeStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny talent node state
func (m DestinyDestinyTalentNodeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyTalentNodeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
