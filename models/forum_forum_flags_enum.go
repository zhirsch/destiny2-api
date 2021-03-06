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

// ForumForumFlagsEnum forum forum flags enum
// swagger:model Forum.ForumFlagsEnum

type ForumForumFlagsEnum int32

// for schema
var forumForumFlagsEnumEnum []interface{}

func init() {
	var res []ForumForumFlagsEnum
	if err := json.Unmarshal([]byte(`[0,1,2,4,8,16,32,64,128]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		forumForumFlagsEnumEnum = append(forumForumFlagsEnumEnum, v)
	}
}

func (m ForumForumFlagsEnum) validateForumForumFlagsEnumEnum(path, location string, value ForumForumFlagsEnum) error {
	if err := validate.Enum(path, location, value, forumForumFlagsEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this forum forum flags enum
func (m ForumForumFlagsEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateForumForumFlagsEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
