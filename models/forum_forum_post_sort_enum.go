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

// ForumForumPostSortEnum forum forum post sort enum
// swagger:model Forum.ForumPostSortEnum

type ForumForumPostSortEnum int32

// for schema
var forumForumPostSortEnumEnum []interface{}

func init() {
	var res []ForumForumPostSortEnum
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		forumForumPostSortEnumEnum = append(forumForumPostSortEnumEnum, v)
	}
}

func (m ForumForumPostSortEnum) validateForumForumPostSortEnumEnum(path, location string, value ForumForumPostSortEnum) error {
	if err := validate.Enum(path, location, value, forumForumPostSortEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this forum forum post sort enum
func (m ForumForumPostSortEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateForumForumPostSortEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
