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

// DestinyBucketScope destiny bucket scope
// swagger:model Destiny.BucketScope

type DestinyBucketScope int32

// for schema
var destinyBucketScopeEnum []interface{}

func init() {
	var res []DestinyBucketScope
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyBucketScopeEnum = append(destinyBucketScopeEnum, v)
	}
}

func (m DestinyBucketScope) validateDestinyBucketScopeEnum(path, location string, value DestinyBucketScope) error {
	if err := validate.Enum(path, location, value, destinyBucketScopeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny bucket scope
func (m DestinyBucketScope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyBucketScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
