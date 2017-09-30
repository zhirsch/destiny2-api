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

// GroupsV2GroupsForMemberFilter groups v2 groups for member filter
// swagger:model GroupsV2.GroupsForMemberFilter

type GroupsV2GroupsForMemberFilter int32

// for schema
var groupsV2GroupsForMemberFilterEnum []interface{}

func init() {
	var res []GroupsV2GroupsForMemberFilter
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupsV2GroupsForMemberFilterEnum = append(groupsV2GroupsForMemberFilterEnum, v)
	}
}

func (m GroupsV2GroupsForMemberFilter) validateGroupsV2GroupsForMemberFilterEnum(path, location string, value GroupsV2GroupsForMemberFilter) error {
	if err := validate.Enum(path, location, value, groupsV2GroupsForMemberFilterEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this groups v2 groups for member filter
func (m GroupsV2GroupsForMemberFilter) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGroupsV2GroupsForMemberFilterEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
