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

// DestinyHistoricalStatsDefinitionsDestinyActivityModeType destiny historical stats definitions destiny activity mode type
// swagger:model Destiny.HistoricalStats.Definitions.DestinyActivityModeType

type DestinyHistoricalStatsDefinitionsDestinyActivityModeType int32

// for schema
var destinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum []interface{}

func init() {
	var res []DestinyHistoricalStatsDefinitionsDestinyActivityModeType
	if err := json.Unmarshal([]byte(`[0,2,3,4,5,6,7,9,10,11,12,13,15,16,17,18,19,20,21,22,24,25,26,27,28,29,30,31,32,37,38,39,40]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		destinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum = append(destinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum, v)
	}
}

func (m DestinyHistoricalStatsDefinitionsDestinyActivityModeType) validateDestinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum(path, location string, value DestinyHistoricalStatsDefinitionsDestinyActivityModeType) error {
	if err := validate.Enum(path, location, value, destinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this destiny historical stats definitions destiny activity mode type
func (m DestinyHistoricalStatsDefinitionsDestinyActivityModeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDestinyHistoricalStatsDefinitionsDestinyActivityModeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
