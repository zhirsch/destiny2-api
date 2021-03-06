// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyResponsesDestinyCharacterResponse The response contract for GetDestinyCharacter, with components that can be returned for character and item-level data.
// swagger:model Destiny.Responses.DestinyCharacterResponse

type DestinyResponsesDestinyCharacterResponse struct {

	// Activity data - info about current activities available to the player.
	// COMPONENT TYPE: CharacterActivities
	Activities *SingleComponentResponseOfDestinyCharacterActivitiesComponent `json:"activities,omitempty"`

	// Base information about the character in question.
	// COMPONENT TYPE: Characters
	Character *SingleComponentResponseOfDestinyCharacterComponent `json:"character,omitempty"`

	// Equipped items on the character.
	// COMPONENT TYPE: CharacterEquipment
	Equipment *SingleComponentResponseOfDestinyInventoryComponent `json:"equipment,omitempty"`

	// The character-level non-equipped inventory items.
	// COMPONENT TYPE: CharacterInventories
	Inventory *SingleComponentResponseOfDestinyInventoryComponent `json:"inventory,omitempty"`

	// The set of components belonging to the player's instanced items.
	// COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
	ItemComponents *DestinyItemComponentSetOfint64 `json:"itemComponents,omitempty"`

	// Items available from Kiosks that are available to this specific character.
	// COMPONENT TYPE: Kiosks
	Kiosks *SingleComponentResponseOfDestinyKiosksComponent `json:"kiosks,omitempty"`

	// Character progression data, including Milestones.
	// COMPONENT TYPE: CharacterProgressions
	Progressions *SingleComponentResponseOfDestinyCharacterProgressionComponent `json:"progressions,omitempty"`

	// Character rendering data - a minimal set of information about equipment and dyes used for rendering.
	// COMPONENT TYPE: CharacterRenderData
	RenderData *SingleComponentResponseOfDestinyCharacterRenderComponent `json:"renderData,omitempty"`
}

/* polymorph Destiny.Responses.DestinyCharacterResponse activities false */

/* polymorph Destiny.Responses.DestinyCharacterResponse character false */

/* polymorph Destiny.Responses.DestinyCharacterResponse equipment false */

/* polymorph Destiny.Responses.DestinyCharacterResponse inventory false */

/* polymorph Destiny.Responses.DestinyCharacterResponse itemComponents false */

/* polymorph Destiny.Responses.DestinyCharacterResponse kiosks false */

/* polymorph Destiny.Responses.DestinyCharacterResponse progressions false */

/* polymorph Destiny.Responses.DestinyCharacterResponse renderData false */

// Validate validates this destiny responses destiny character response
func (m *DestinyResponsesDestinyCharacterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCharacter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEquipment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInventory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItemComponents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKiosks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProgressions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRenderData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateActivities(formats strfmt.Registry) error {

	if swag.IsZero(m.Activities) { // not required
		return nil
	}

	if m.Activities != nil {

		if err := m.Activities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activities")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateCharacter(formats strfmt.Registry) error {

	if swag.IsZero(m.Character) { // not required
		return nil
	}

	if m.Character != nil {

		if err := m.Character.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("character")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateEquipment(formats strfmt.Registry) error {

	if swag.IsZero(m.Equipment) { // not required
		return nil
	}

	if m.Equipment != nil {

		if err := m.Equipment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("equipment")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateInventory(formats strfmt.Registry) error {

	if swag.IsZero(m.Inventory) { // not required
		return nil
	}

	if m.Inventory != nil {

		if err := m.Inventory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateItemComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemComponents) { // not required
		return nil
	}

	if m.ItemComponents != nil {

		if err := m.ItemComponents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemComponents")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateKiosks(formats strfmt.Registry) error {

	if swag.IsZero(m.Kiosks) { // not required
		return nil
	}

	if m.Kiosks != nil {

		if err := m.Kiosks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kiosks")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateProgressions(formats strfmt.Registry) error {

	if swag.IsZero(m.Progressions) { // not required
		return nil
	}

	if m.Progressions != nil {

		if err := m.Progressions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressions")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyResponsesDestinyCharacterResponse) validateRenderData(formats strfmt.Registry) error {

	if swag.IsZero(m.RenderData) { // not required
		return nil
	}

	if m.RenderData != nil {

		if err := m.RenderData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renderData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyResponsesDestinyCharacterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyResponsesDestinyCharacterResponse) UnmarshalBinary(b []byte) error {
	var res DestinyResponsesDestinyCharacterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
