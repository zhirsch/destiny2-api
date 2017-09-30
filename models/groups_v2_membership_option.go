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

// GroupsV2MembershipOption groups v2 membership option
// swagger:model GroupsV2.MembershipOption

type GroupsV2MembershipOption int32

// for schema
var groupsV2MembershipOptionEnum []interface{}

func init() {
	var res []GroupsV2MembershipOption
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupsV2MembershipOptionEnum = append(groupsV2MembershipOptionEnum, v)
	}
}

func (m GroupsV2MembershipOption) validateGroupsV2MembershipOptionEnum(path, location string, value GroupsV2MembershipOption) error {
	if err := validate.Enum(path, location, value, groupsV2MembershipOptionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this groups v2 membership option
func (m GroupsV2MembershipOption) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGroupsV2MembershipOptionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
