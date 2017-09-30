// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyMilestoneActivityCompletionStatus Represents this player's personal completion status for the Activity under a Milestone, if the activity has trackable completion and progress information. (most activities won't, or the concept won't apply. For instance, it makes sense to talk about a tier of a raid as being Completed or having progress, but it doesn't make sense to talk about a Crucible Playlist in those terms.
// swagger:model Destiny.Milestones.DestinyMilestoneActivityCompletionStatus

type DestinyMilestonesDestinyMilestoneActivityCompletionStatus struct {

	// If the activity has been "completed", that information will be returned here.
	Completed bool `json:"completed,omitempty"`

	// phases
	Phases DestinyMilestonesDestinyMilestoneActivityCompletionStatusPhases `json:"phases"`
}

/* polymorph Destiny.Milestones.DestinyMilestoneActivityCompletionStatus completed false */

/* polymorph Destiny.Milestones.DestinyMilestoneActivityCompletionStatus phases false */

// Validate validates this destiny milestones destiny milestone activity completion status
func (m *DestinyMilestonesDestinyMilestoneActivityCompletionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneActivityCompletionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneActivityCompletionStatus) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyMilestoneActivityCompletionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}