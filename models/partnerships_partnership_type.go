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

// PartnershipsPartnershipType Representing external partners to which BNet users can link accounts, but that are not Account System credentials: partnerships that BNet uses exclusively for data.
// swagger:model Partnerships.PartnershipType

type PartnershipsPartnershipType int32

// for schema
var partnershipsPartnershipTypeEnum []interface{}

func init() {
	var res []PartnershipsPartnershipType
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partnershipsPartnershipTypeEnum = append(partnershipsPartnershipTypeEnum, v)
	}
}

func (m PartnershipsPartnershipType) validatePartnershipsPartnershipTypeEnum(path, location string, value PartnershipsPartnershipType) error {
	if err := validate.Enum(path, location, value, partnershipsPartnershipTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this partnerships partnership type
func (m PartnershipsPartnershipType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePartnershipsPartnershipTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}