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

// DestinyBucketCategory destiny bucket category
// swagger:model Destiny.BucketCategory

type DestinyBucketCategory int32

// for schema
var destinyBucketCategoryEnum []interface{}

func init() {
	var res []DestinyBucketCategory
	if err := json.Unmarshal([]byte(`[0,1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyBucketCategoryEnum = append(destinyBucketCategoryEnum, v)
	}
}

func (m DestinyBucketCategory) validateDestinyBucketCategoryEnum(path, location string, value DestinyBucketCategory) error {
	if err := validate.Enum(path, location, value, destinyBucketCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny bucket category
func (m DestinyBucketCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyBucketCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}