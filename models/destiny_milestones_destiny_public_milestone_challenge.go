// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyPublicMilestoneChallenge A Milestone can have many Challenges. Challenges are just extra Objectives that provide a fun way to mix-up play and provide extra rewards.
// swagger:model Destiny.Milestones.DestinyPublicMilestoneChallenge

type DestinyMilestonesDestinyPublicMilestoneChallenge struct {

	// IF the Objective is related to a specific Activity, this will be that activity's hash. Use it to look up the DestinyActivityDefinition for additional data to show.
	ActivityHash uint32 `json:"activityHash,omitempty"`

	// The objective for the Challenge, which should have human-readable data about what needs to be done to accomplish the objective. Use this hash to look up the DestinyObjectiveDefinition.
	ObjectiveHash uint32 `json:"objectiveHash,omitempty"`
}

/* polymorph Destiny.Milestones.DestinyPublicMilestoneChallenge activityHash false */

/* polymorph Destiny.Milestones.DestinyPublicMilestoneChallenge objectiveHash false */

// Validate validates this destiny milestones destiny public milestone challenge
func (m *DestinyMilestonesDestinyPublicMilestoneChallenge) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestoneChallenge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyPublicMilestoneChallenge) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyPublicMilestoneChallenge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
