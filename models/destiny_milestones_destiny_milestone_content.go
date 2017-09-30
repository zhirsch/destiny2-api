// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyMilestonesDestinyMilestoneContent Represents localized, extended content related to Milestones. This is intentionally returned by a separate endpoint and not with Character-level Milestone data because we do not put localized data into standard Destiny responses, both for brevity of response and for caching purposes. If you really need this data, hit the Milestone Content endpoint.
// swagger:model Destiny.Milestones.DestinyMilestoneContent

type DestinyMilestonesDestinyMilestoneContent struct {

	// The "About this Milestone" text from the Firehose.
	About string `json:"about,omitempty"`

	// item categories
	ItemCategories DestinyMilestonesDestinyMilestoneContentItemCategories `json:"itemCategories"`

	// The Current Status of the Milestone, as driven by the Firehose.
	Status string `json:"status,omitempty"`

	// A list of tips, provided by the Firehose.
	Tips []string `json:"tips"`
}

/* polymorph Destiny.Milestones.DestinyMilestoneContent about false */

/* polymorph Destiny.Milestones.DestinyMilestoneContent itemCategories false */

/* polymorph Destiny.Milestones.DestinyMilestoneContent status false */

/* polymorph Destiny.Milestones.DestinyMilestoneContent tips false */

// Validate validates this destiny milestones destiny milestone content
func (m *DestinyMilestonesDestinyMilestoneContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTips(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyMilestonesDestinyMilestoneContent) validateTips(formats strfmt.Registry) error {

	if swag.IsZero(m.Tips) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyMilestonesDestinyMilestoneContent) UnmarshalBinary(b []byte) error {
	var res DestinyMilestonesDestinyMilestoneContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}