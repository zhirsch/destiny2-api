// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyTalentNodeStepGroups These properties are an attempt to categorize talent node steps by certain common properties. See the related enumerations for the type of properties being categorized.
// swagger:model Destiny.Definitions.DestinyTalentNodeStepGroups

type DestinyDefinitionsDestinyTalentNodeStepGroups struct {

	// damage types
	DamageTypes DestinyDefinitionsDestinyTalentNodeStepDamageTypes `json:"damageTypes,omitempty"`

	// guardian attributes
	GuardianAttributes DestinyDefinitionsDestinyTalentNodeStepGuardianAttributes `json:"guardianAttributes,omitempty"`

	// impact effects
	ImpactEffects DestinyDefinitionsDestinyTalentNodeStepImpactEffects `json:"impactEffects,omitempty"`

	// light abilities
	LightAbilities DestinyDefinitionsDestinyTalentNodeStepLightAbilities `json:"lightAbilities,omitempty"`

	// weapon performance
	WeaponPerformance DestinyDefinitionsDestinyTalentNodeStepWeaponPerformances `json:"weaponPerformance,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyTalentNodeStepGroups damageTypes false */

/* polymorph Destiny.Definitions.DestinyTalentNodeStepGroups guardianAttributes false */

/* polymorph Destiny.Definitions.DestinyTalentNodeStepGroups impactEffects false */

/* polymorph Destiny.Definitions.DestinyTalentNodeStepGroups lightAbilities false */

/* polymorph Destiny.Definitions.DestinyTalentNodeStepGroups weaponPerformance false */

// Validate validates this destiny definitions destiny talent node step groups
func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDamageTypes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGuardianAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImpactEffects(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLightAbilities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWeaponPerformance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) validateDamageTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.DamageTypes) { // not required
		return nil
	}

	if err := m.DamageTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("damageTypes")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) validateGuardianAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.GuardianAttributes) { // not required
		return nil
	}

	if err := m.GuardianAttributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("guardianAttributes")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) validateImpactEffects(formats strfmt.Registry) error {

	if swag.IsZero(m.ImpactEffects) { // not required
		return nil
	}

	if err := m.ImpactEffects.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("impactEffects")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) validateLightAbilities(formats strfmt.Registry) error {

	if swag.IsZero(m.LightAbilities) { // not required
		return nil
	}

	if err := m.LightAbilities.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lightAbilities")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) validateWeaponPerformance(formats strfmt.Registry) error {

	if swag.IsZero(m.WeaponPerformance) { // not required
		return nil
	}

	if err := m.WeaponPerformance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("weaponPerformance")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyTalentNodeStepGroups) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyTalentNodeStepGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}