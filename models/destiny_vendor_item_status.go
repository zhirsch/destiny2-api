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

// DestinyVendorItemStatus destiny vendor item status
// swagger:model Destiny.VendorItemStatus

type DestinyVendorItemStatus int32

// for schema
var destinyVendorItemStatusEnum []interface{}

func init() {
	var res []DestinyVendorItemStatus
	if err := json.Unmarshal([]byte(`[0,1,2,4,8,16,32,64,128,256,512,1024,2048,4096]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyVendorItemStatusEnum = append(destinyVendorItemStatusEnum, v)
	}
}

func (m DestinyVendorItemStatus) validateDestinyVendorItemStatusEnum(path, location string, value DestinyVendorItemStatus) error {
	if err := validate.Enum(path, location, value, destinyVendorItemStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny vendor item status
func (m DestinyVendorItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyVendorItemStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}