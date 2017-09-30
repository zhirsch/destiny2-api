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

// DestinyDefinitionsDestinyRewardSourceCategory BNet's custom categorization of reward sources. We took a look at the existing ways that items could be spawned, and tried to make high-level categorizations of them. This needs to be re-evaluated for Destiny 2.
// swagger:model Destiny.Definitions.DestinyRewardSourceCategory

type DestinyDefinitionsDestinyRewardSourceCategory int32

// for schema
var destinyDefinitionsDestinyRewardSourceCategoryEnum []interface{}

func init() {
	var res []DestinyDefinitionsDestinyRewardSourceCategory
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDefinitionsDestinyRewardSourceCategoryEnum = append(destinyDefinitionsDestinyRewardSourceCategoryEnum, v)
	}
}

func (m DestinyDefinitionsDestinyRewardSourceCategory) validateDestinyDefinitionsDestinyRewardSourceCategoryEnum(path, location string, value DestinyDefinitionsDestinyRewardSourceCategory) error {
	if err := validate.Enum(path, location, value, destinyDefinitionsDestinyRewardSourceCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny definitions destiny reward source category
func (m DestinyDefinitionsDestinyRewardSourceCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDefinitionsDestinyRewardSourceCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
