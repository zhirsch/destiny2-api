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

// DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod destiny historical stats definitions destiny stats merge method
// swagger:model Destiny.HistoricalStats.Definitions.DestinyStatsMergeMethod

type DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod int32

// for schema
var destinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum []interface{}

func init() {
	var res []DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum = append(destinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum, v)
	}
}

func (m DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod) validateDestinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum(path, location string, value DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod) error {
	if err := validate.Enum(path, location, value, destinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny historical stats definitions destiny stats merge method
func (m DestinyHistoricalStatsDefinitionsDestinyStatsMergeMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyHistoricalStatsDefinitionsDestinyStatsMergeMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}