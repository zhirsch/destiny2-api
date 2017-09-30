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

// DestinyDestinyUnlockValueUIStyle If you're showing an unlock value in the UI, this is the format in which it should be shown. You'll have to build your own algorithms on the client side to determine how best to render these options.
// swagger:model Destiny.DestinyUnlockValueUIStyle

type DestinyDestinyUnlockValueUIStyle int32

// for schema
var destinyDestinyUnlockValueUiStyleEnum []interface{}

func init() {
	var res []DestinyDestinyUnlockValueUIStyle
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7,8]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDestinyUnlockValueUiStyleEnum = append(destinyDestinyUnlockValueUiStyleEnum, v)
	}
}

func (m DestinyDestinyUnlockValueUIStyle) validateDestinyDestinyUnlockValueUIStyleEnum(path, location string, value DestinyDestinyUnlockValueUIStyle) error {
	if err := validate.Enum(path, location, value, destinyDestinyUnlockValueUiStyleEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny destiny unlock value UI style
func (m DestinyDestinyUnlockValueUIStyle) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDestinyUnlockValueUIStyleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
