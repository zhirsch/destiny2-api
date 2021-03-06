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

// DestinyDefinitionsDestinyTalentNodeStepLightAbilities destiny definitions destiny talent node step light abilities
// swagger:model Destiny.Definitions.DestinyTalentNodeStepLightAbilities

type DestinyDefinitionsDestinyTalentNodeStepLightAbilities int32

// for schema
var destinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum []interface{}

func init() {
	var res []DestinyDefinitionsDestinyTalentNodeStepLightAbilities
	if err := json.Unmarshal([]byte(`[0,1,2,4,8,16,32,63]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum = append(destinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum, v)
	}
}

func (m DestinyDefinitionsDestinyTalentNodeStepLightAbilities) validateDestinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum(path, location string, value DestinyDefinitionsDestinyTalentNodeStepLightAbilities) error {
	if err := validate.Enum(path, location, value, destinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny definitions destiny talent node step light abilities
func (m DestinyDefinitionsDestinyTalentNodeStepLightAbilities) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyDefinitionsDestinyTalentNodeStepLightAbilitiesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
