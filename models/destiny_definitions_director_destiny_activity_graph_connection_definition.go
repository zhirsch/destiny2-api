// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition Nodes on a graph can be visually connected: this appears to be the information about which nodes to link. It appears to lack more detailed information, such as the path for that linking.
// swagger:model Destiny.Definitions.Director.DestinyActivityGraphConnectionDefinition

type DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition struct {

	// dest node hash
	DestNodeHash uint32 `json:"destNodeHash,omitempty"`

	// source node hash
	SourceNodeHash uint32 `json:"sourceNodeHash,omitempty"`
}

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphConnectionDefinition destNodeHash false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphConnectionDefinition sourceNodeHash false */

// Validate validates this destiny definitions director destiny activity graph connection definition
func (m *DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDirectorDestinyActivityGraphConnectionDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
