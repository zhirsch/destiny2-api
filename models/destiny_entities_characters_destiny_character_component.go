// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesCharactersDestinyCharacterComponent This component contains base properties of the character. You'll probably want to always request this component, but hey you do you.
// swagger:model Destiny.Entities.Characters.DestinyCharacterComponent

type DestinyEntitiesCharactersDestinyCharacterComponent struct {

	// The "base" level of your character, not accounting for any light level.
	BaseCharacterLevel int32 `json:"baseCharacterLevel,omitempty"`

	// The unique identifier for the character.
	CharacterID int64 `json:"characterId,omitempty,string"`

	// Use this hash to look up the character's DestinyClassDefinition.
	ClassHash uint32 `json:"classHash,omitempty"`

	// Mostly for historical purposes at this point, this is an enumeration for the character's class.
	// It'll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove.
	ClassType DestinyDestinyClass `json:"classType,omitempty"`

	// The last date that the user played Destiny.
	DateLastPlayed strfmt.DateTime `json:"dateLastPlayed,omitempty"`

	// A shortcut path to the user's currently equipped emblem background image. If you're just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition.
	EmblemBackgroundPath string `json:"emblemBackgroundPath,omitempty"`

	// The hash of the currently equipped emblem for the user. Can be used to look up the DestinyInventoryItemDefinition.
	EmblemHash uint32 `json:"emblemHash,omitempty"`

	// A shortcut path to the user's currently equipped emblem image. If you're just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition.
	EmblemPath string `json:"emblemPath,omitempty"`

	// Use this hash to look up the character's DestinyGenderDefinition.
	GenderHash uint32 `json:"genderHash,omitempty"`

	// Mostly for historical purposes at this point, this is an enumeration for the character's Gender.
	// It'll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. And yeah, it's an enumeration and not a boolean. Fight me.
	GenderType DestinyDestinyGender `json:"genderType,omitempty"`

	// The progression that indicates your character's level. Not their light level, but their character level: you know, the thing you max out a couple hours in and then ignore for the sake of light level.
	LevelProgression *DestinyDestinyProgression `json:"levelProgression,omitempty"`

	// The user's calculated "Light Level". Light level is an indicator of your power that mostly matters in the end game, once you've reached the maximum character level: it's a level that's dependent on the average Attack/Defense power of your items.
	Light int32 `json:"light,omitempty"`

	// Every Destiny Profile has a membershipId. This is provided on the character as well for convenience.
	MembershipID int64 `json:"membershipId,omitempty,string"`

	// membershipType tells you the platform on which the character plays. Examine the BungieMembershipType enumeration for possible values.
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`

	// If the user is currently playing, this is how long they've been playing.
	MinutesPlayedThisSession int64 `json:"minutesPlayedThisSession,omitempty,string"`

	// If this value is 525,600, then they played Destiny for a year. Or they're a very dedicated Rent fan. Note that this includes idle time, not just time spent actually in activities shooting things.
	MinutesPlayedTotal int64 `json:"minutesPlayedTotal,omitempty,string"`

	// A number between 0 and 100, indicating the whole and fractional % remaining to get to the next character level.
	PercentToNextLevel float32 `json:"percentToNextLevel,omitempty"`

	// Use this hash to look up the character's DestinyRaceDefinition.
	RaceHash uint32 `json:"raceHash,omitempty"`

	// Mostly for historical purposes at this point, this is an enumeration for the character's race.
	// It'll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove.
	RaceType DestinyDestinyRace `json:"raceType,omitempty"`

	// Your character's stats, such as Agility, Resilience, etc... *not* historical stats.
	// You'll have to call a different endpoint for those.
	Stats map[string]int32 `json:"stats,omitempty"`
}

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent baseCharacterLevel false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent characterId false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent classHash false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent classType false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent dateLastPlayed false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent emblemBackgroundPath false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent emblemHash false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent emblemPath false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent genderHash false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent genderType false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent levelProgression false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent light false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent membershipId false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent membershipType false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent minutesPlayedThisSession false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent minutesPlayedTotal false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent percentToNextLevel false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent raceHash false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent raceType false */

/* polymorph Destiny.Entities.Characters.DestinyCharacterComponent stats false */

// Validate validates this destiny entities characters destiny character component
func (m *DestinyEntitiesCharactersDestinyCharacterComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGenderType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLevelProgression(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembershipType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRaceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyEntitiesCharactersDestinyCharacterComponent) validateClassType(formats strfmt.Registry) error {

	if swag.IsZero(m.ClassType) { // not required
		return nil
	}

	if err := m.ClassType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("classType")
		}
		return err
	}

	return nil
}

func (m *DestinyEntitiesCharactersDestinyCharacterComponent) validateGenderType(formats strfmt.Registry) error {

	if swag.IsZero(m.GenderType) { // not required
		return nil
	}

	if err := m.GenderType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("genderType")
		}
		return err
	}

	return nil
}

func (m *DestinyEntitiesCharactersDestinyCharacterComponent) validateLevelProgression(formats strfmt.Registry) error {

	if swag.IsZero(m.LevelProgression) { // not required
		return nil
	}

	if m.LevelProgression != nil {

		if err := m.LevelProgression.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("levelProgression")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyEntitiesCharactersDestinyCharacterComponent) validateMembershipType(formats strfmt.Registry) error {

	if swag.IsZero(m.MembershipType) { // not required
		return nil
	}

	if err := m.MembershipType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("membershipType")
		}
		return err
	}

	return nil
}

func (m *DestinyEntitiesCharactersDestinyCharacterComponent) validateRaceType(formats strfmt.Registry) error {

	if swag.IsZero(m.RaceType) { // not required
		return nil
	}

	if err := m.RaceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("raceType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesCharactersDestinyCharacterComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesCharactersDestinyCharacterComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesCharactersDestinyCharacterComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
