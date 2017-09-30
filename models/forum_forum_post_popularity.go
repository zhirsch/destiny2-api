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

// ForumForumPostPopularity forum forum post popularity
// swagger:model Forum.ForumPostPopularity

type ForumForumPostPopularity int32

// for schema
var forumForumPostPopularityEnum []interface{}

func init() {
	var res []ForumForumPostPopularity
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		forumForumPostPopularityEnum = append(forumForumPostPopularityEnum, v)
	}
}

func (m ForumForumPostPopularity) validateForumForumPostPopularityEnum(path, location string, value ForumForumPostPopularity) error {
	if err := validate.Enum(path, location, value, forumForumPostPopularityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this forum forum post popularity
func (m ForumForumPostPopularity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateForumForumPostPopularityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
