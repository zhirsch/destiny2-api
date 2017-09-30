// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupsV2GroupCreationResponse groups v2 group creation response
// swagger:model GroupsV2.GroupCreationResponse

type GroupsV2GroupCreationResponse struct {

	// group Id
	GroupID int64 `json:"groupId,omitempty"`
}

/* polymorph GroupsV2.GroupCreationResponse groupId false */

// Validate validates this groups v2 group creation response
func (m *GroupsV2GroupCreationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GroupsV2GroupCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupsV2GroupCreationResponse) UnmarshalBinary(b []byte) error {
	var res GroupsV2GroupCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
