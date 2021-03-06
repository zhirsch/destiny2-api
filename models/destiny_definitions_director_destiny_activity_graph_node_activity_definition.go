// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition The actual activity to be redirected to when you click on the node. Note that a node can have many Activities attached to it: but only one will be active at any given time. The list of Node Activities will be traversed, and the first one found to be active will be displayed. This way, a node can layer multiple variants of an activity on top of each other. For instance, one node can control the weekly Crucible Playlist. There are multiple possible playlists, but only one is active for the week.
// swagger:model Destiny.Definitions.Director.DestinyActivityGraphNodeActivityDefinition

type DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition struct {

	// The activity that will be activated if the user clicks on this node. Controls all activity-related information displayed on the node if it is active (the text shown in the tooltip etc)
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// An identifier for this node activity. It is only guaranteed to be unique within the Activity Graph.
	NodeActivityID uint32 `json:"nodeActivityId,omitempty"`
}

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeActivityDefinition activityHash false */

/* polymorph Destiny.Definitions.Director.DestinyActivityGraphNodeActivityDefinition nodeActivityId false */

// Validate validates this destiny definitions director destiny activity graph node activity definition
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDirectorDestinyActivityGraphNodeActivityDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
