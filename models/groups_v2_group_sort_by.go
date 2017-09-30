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

// GroupsV2GroupSortBy groups v2 group sort by
// swagger:model GroupsV2.GroupSortBy

type GroupsV2GroupSortBy int32

// for schema
var groupsV2GroupSortByEnum []interface{}

func init() {
	var res []GroupsV2GroupSortBy
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupsV2GroupSortByEnum = append(groupsV2GroupSortByEnum, v)
	}
}

func (m GroupsV2GroupSortBy) validateGroupsV2GroupSortByEnum(path, location string, value GroupsV2GroupSortBy) error {
	if err := validate.Enum(path, location, value, groupsV2GroupSortByEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this groups v2 group sort by
func (m GroupsV2GroupSortBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGroupsV2GroupSortByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}