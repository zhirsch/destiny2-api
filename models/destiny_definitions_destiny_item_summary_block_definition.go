// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemSummaryBlockDefinition This appears to be information used when rendering rewards. We don't currently use it on BNet.
// swagger:model Destiny.Definitions.DestinyItemSummaryBlockDefinition

type DestinyDefinitionsDestinyItemSummaryBlockDefinition struct {

	// Apparently when rendering an item in a reward, this should be used as a sort priority. We're not doing it presently.
	SortPriority int32 `json:"sortPriority,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemSummaryBlockDefinition sortPriority false */

// Validate validates this destiny definitions destiny item summary block definition
func (m *DestinyDefinitionsDestinyItemSummaryBlockDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSummaryBlockDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemSummaryBlockDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemSummaryBlockDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
