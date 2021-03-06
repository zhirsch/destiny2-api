// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyNodeStepDefinition This defines the properties of a "Talent Node Step". When you see a talent node in game, the actual visible properties that you see (its icon, description, the perks and stats it provides) are not provided by the Node itself, but rather by the currently active Step on the node.
// When a Talent Node is activated, the currently active step's benefits are conferred upon the item and character.
// The currently active step on talent nodes are determined when an item is first instantiated. Sometimes it is random, sometimes it is more deterministic (particularly when a node has only a single step).
// Note that, when dealing with Talent Node Steps, you must ensure that you have the latest version of content. stepIndex and nodeStepHash - two ways of identifying the step within a node - are both content version dependent, and thus are subject to change between content updates.
// swagger:model Destiny.Definitions.DestinyNodeStepDefinition

type DestinyDefinitionsDestinyNodeStepDefinition struct {

	// If the step has requirements for activation (they almost always do, if nothing else than for the Talent Grid's Progression to have reached a certain level), they will be defined here.
	ActivationRequirement *DestinyDefinitionsDestinyNodeActivationRequirement `json:"activationRequirement,omitempty"`

	// If true, this step can affect the level of the item. See DestinyInventoryItemDefintion for more information about item levels and their effect on stats.
	AffectsLevel bool `json:"affectsLevel,omitempty"`

	// If this is true, the step affects the item's Quality in some way. See DestinyInventoryItemDefinition for more information about the meaning of Quality. I already made a joke about Zen and the Art of Motorcycle Maintenance elsewhere in the documentation, so I will avoid doing it again. Oops too late
	AffectsQuality bool `json:"affectsQuality,omitempty"`

	// There was a time when talent nodes could be activated multiple times, and the effects of subsequent Steps would be compounded on each other, essentially "upgrading" the node. We have moved away from this, but theoretically the capability still exists.
	// I continue to return this in case it is used in the future: if true and this step is the current step in the node, you are allowed to activate the node a second time to receive the benefits of the next step in the node, which will then become the active step.
	CanActivateNextStep bool `json:"canActivateNextStep,omitempty"`

	// An enum representing a damage type granted by activating this step, if any.
	DamageType DestinyDamageType `json:"damageType,omitempty"`

	// If the step provides a damage type, this will be the hash identifier used to look up the damage type's DestinyDamageTypeDefinition.
	DamageTypeHash uint32 `json:"damageTypeHash,omitempty"`

	// These are the display properties actually used to render the Talent Node. The currently active step's displayProperties are shown.
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// If you can interact with this node in some way, this is the localized description of that interaction.
	InteractionDescription string `json:"interactionDescription,omitempty"`

	// If true, the next step to be chosen is random, and if you're allowed to activate the next step. (if canActivateNextStep = true)
	IsNextStepRandom bool `json:"isNextStepRandom,omitempty"`

	// The stepIndex of the next step in the talent node, or -1 if this is the last step or if the next step to be chosen is random.
	// This doesn't really matter anymore unless canActivateNextStep begins to be used again.
	NextStepIndex int32 `json:"nextStepIndex,omitempty"`

	// The hash of this node step. Unfortunately, while it can be used to uniquely identify the step within a node, it is also content version dependent and should not be relied on without ensuring you have the latest vesion of content.
	NodeStepHash uint32 `json:"nodeStepHash,omitempty"`

	// The list of hash identifiers for Perks (DestinySandboxPerkDefinition) that are applied when this step is active. Perks provide a variety of benefits and modifications - examine DestinySandboxPerkDefinition to learn more.
	PerkHashes []uint32 `json:"perkHashes"`

	// socket replacements
	SocketReplacements DestinyDefinitionsDestinyNodeStepDefinitionSocketReplacements `json:"socketReplacements"`

	// When the Talent Grid's progression reaches this value, the circular "progress bar" that surrounds the talent node should be shown.
	// This also indicates the lower bound of said progress bar, with the upper bound being the progress required to reach activationRequirement.gridLevel. (at some point I should precalculate the upper bound and put it in the definition to save people time)
	StartProgressionBarAtProgress int32 `json:"startProgressionBarAtProgress,omitempty"`

	// When the step provides stat benefits on the item or character, this is the list of hash identifiers for stats (DestinyStatDefinition) that are provided.
	StatHashes []uint32 `json:"statHashes"`

	// In Destiny 1, the Armory's Perk Filtering was driven by a concept of TalentNodeStepGroups: categorizations of talent nodes based on their functionality. While the Armory isn't a BNet-facing thing for now, and the new Armory will need to account for Sockets rather than Talent Nodes, this categorization capability feels useful enough to still keep around.
	StepGroups *DestinyDefinitionsDestinyTalentNodeStepGroups `json:"stepGroups,omitempty"`

	// The index of this step in the list of Steps on the Talent Node.
	// Unfortunately, this is the closest thing we have to an identifier for the Step: steps are not provided a content version agnostic identifier. This means that, when you are dealing with talent nodes, you will need to first ensure that you have the latest version of content.
	StepIndex int32 `json:"stepIndex,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition activationRequirement false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition affectsLevel false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition affectsQuality false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition canActivateNextStep false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition damageType false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition damageTypeHash false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition interactionDescription false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition isNextStepRandom false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition nextStepIndex false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition nodeStepHash false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition perkHashes false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition socketReplacements false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition startProgressionBarAtProgress false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition statHashes false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition stepGroups false */

/* polymorph Destiny.Definitions.DestinyNodeStepDefinition stepIndex false */

// Validate validates this destiny definitions destiny node step definition
func (m *DestinyDefinitionsDestinyNodeStepDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationRequirement(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDamageType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePerkHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStepGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validateActivationRequirement(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivationRequirement) { // not required
		return nil
	}

	if m.ActivationRequirement != nil {

		if err := m.ActivationRequirement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activationRequirement")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validateDamageType(formats strfmt.Registry) error {

	if swag.IsZero(m.DamageType) { // not required
		return nil
	}

	if err := m.DamageType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("damageType")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validateDisplayProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayProperties) { // not required
		return nil
	}

	if m.DisplayProperties != nil {

		if err := m.DisplayProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("displayProperties")
			}
			return err
		}
	}

	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validatePerkHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.PerkHashes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validateStatHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.StatHashes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyNodeStepDefinition) validateStepGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.StepGroups) { // not required
		return nil
	}

	if m.StepGroups != nil {

		if err := m.StepGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stepGroups")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyNodeStepDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyNodeStepDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyNodeStepDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
