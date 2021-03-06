// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyTalentNodeDefinition Talent Grids on items have Nodes. These nodes have positions in the talent grid's UI, and contain "Steps" (DestinyTalentNodeStepDefinition), one of whom will be the "Current" step.
// The Current Step determines the visual properties of the node, as well as what the node grants when it is activated.
// See DestinyTalentGridDefinition for a more complete overview of how Talent Grids work, and how they are used in Destiny 2 (and how they were used in Destiny 1).
// swagger:model Destiny.Definitions.DestinyTalentNodeDefinition

type DestinyDefinitionsDestinyTalentNodeDefinition struct {

	// If true, this node will automatically unlock when the Talent Grid's level reaches the required level of the current step of this node.
	AutoUnlocks bool `json:"autoUnlocks,omitempty"`

	// At one point, Talent Nodes supported the idea of "Binary Pairs": nodes that overlapped each other visually, and where activating one deactivated the other. They ended up not being used, mostly because Exclusive Sets are *almost* a superset of this concept, but the potential for it to be used still exists in theory.
	// If this is ever used, this will be the index into the DestinyTalentGridDefinition.nodes property for the node that is the binary pair match to this node. Activating one deactivates the other.
	BinaryPairNodeIndex int32 `json:"binaryPairNodeIndex,omitempty"`

	// The visual "column" where the node should be shown in the UI. If negative, the node is hidden.
	Column int32 `json:"column,omitempty"`

	// The nodeHash values for nodes that are in an Exclusive Set with this node.
	// See DestinyTalentGridDefinition.exclusiveSets for more info about exclusive sets.
	// Again, note that these are nodeHashes and *not* nodeIndexes.
	ExclusiveWithNodeHashes []uint32 `json:"exclusiveWithNodeHashes"`

	// As of Destiny 2, nodes can exist as part of "Exclusive Groups". These differ from exclusive sets in that, within the group, many nodes can be activated. But the act of activating any node in the group will cause "opposing" nodes (nodes in groups that are not allowed to be activated at the same time as this group) to deactivate.
	// See DestinyTalentExclusiveGroup for more information on the details. This is an identifier for this node's group, if it is part of one.
	GroupHash uint32 `json:"groupHash,omitempty"`

	// Comes from the talent grid node style: if true, then this node should be ignored for determining whether the grid is complete.
	IgnoreForCompletion bool `json:"ignoreForCompletion,omitempty"`

	// If this is true, the node's step is determined randomly rather than the first step being chosen.
	IsRandom bool `json:"isRandom,omitempty"`

	// If this is true, the node can be "re-rolled" to acquire a different random current step. This is not used, but still exists for a theoretical future of talent grids.
	IsRandomRepurchasable bool `json:"isRandomRepurchasable,omitempty"`

	// At one point, Nodes were going to be able to be activated multiple times, changing the current step and potentially piling on multiple effects from the previously activated steps. This property would indicate if the last step could be activated multiple times.
	// This is not currently used, but it isn't out of the question that this could end up being used again in a theoretical future.
	LastStepRepeats bool `json:"lastStepRepeats,omitempty"`

	// A string identifier for a custom visual layout to apply to this talent node. Unfortunately, we do not have any data for rendering these custom layouts. It will be up to you to interpret these strings and change your UI if you want to have custom UI matching these layouts.
	LayoutIdentifier string `json:"layoutIdentifier,omitempty"`

	// Talent nodes can be associated with a piece of Lore, generally rendered in a tooltip. This is the hash identifier of the lore element to show, if there is one to be show.
	LoreHash uint32 `json:"loreHash,omitempty"`

	// The hash identifier for the node, which unfortunately is also content version dependent but can be (and ideally, should be) used instead of the nodeIndex to uniquely identify the node.
	// The two exist side-by-side for backcompat reasons due to the Great Talent Node Restructuring of Destiny 1, and I ran out of time to remove one of them and standardize on the other. Sorry!
	NodeHash uint32 `json:"nodeHash,omitempty"`

	// The index into the DestinyTalentGridDefinition's "nodes" property where this node is located. Used to uniquely identify the node within the Talent Grid. Note that this is content version dependent: make sure you have the latest version of content before trying to use these properties.
	NodeIndex int32 `json:"nodeIndex,omitempty"`

	// Comes from the talent grid node style: this identifier should be used to determine how to render the node in the UI.
	NodeStyleIdentifier string `json:"nodeStyleIdentifier,omitempty"`

	// Indexes into the DestinyTalentGridDefinition.nodes property for any nodes that must be activated before this one is allowed to be activated.
	// I would have liked to change this to hashes for Destiny 2, but we have run out of time.
	PrerequisiteNodeIndexes []int32 `json:"prerequisiteNodeIndexes"`

	// At one point, you were going to be able to repurchase talent nodes that had random steps, to "re-roll" the current step of the node (and thus change the properties of your item). This was to be the activation requirement for performing that re-roll.
	// The system still exists to do this, as far as I know, so it may yet come back around!
	RandomActivationRequirement *DestinyDefinitionsDestinyNodeActivationRequirement `json:"randomActivationRequirement,omitempty"`

	// If the node's step is randomly selected, this is the amount of the Talent Grid's progression experience at which the progression bar for the node should be shown.
	RandomStartProgressionBarAtProgression int32 `json:"randomStartProgressionBarAtProgression,omitempty"`

	// The visual "row" where the node should be shown in the UI. If negative, then the node is hidden.
	Row int32 `json:"row,omitempty"`

	// steps
	Steps DestinyDefinitionsDestinyTalentNodeDefinitionSteps `json:"steps"`
}

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition autoUnlocks false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition binaryPairNodeIndex false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition column false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition exclusiveWithNodeHashes false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition groupHash false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition ignoreForCompletion false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition isRandom false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition isRandomRepurchasable false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition lastStepRepeats false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition layoutIdentifier false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition loreHash false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition nodeHash false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition nodeIndex false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition nodeStyleIdentifier false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition prerequisiteNodeIndexes false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition randomActivationRequirement false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition randomStartProgressionBarAtProgression false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition row false */

/* polymorph Destiny.Definitions.DestinyTalentNodeDefinition steps false */

// Validate validates this destiny definitions destiny talent node definition
func (m *DestinyDefinitionsDestinyTalentNodeDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusiveWithNodeHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrerequisiteNodeIndexes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRandomActivationRequirement(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeDefinition) validateExclusiveWithNodeHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.ExclusiveWithNodeHashes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeDefinition) validatePrerequisiteNodeIndexes(formats strfmt.Registry) error {

	if swag.IsZero(m.PrerequisiteNodeIndexes) { // not required
		return nil
	}

	return nil
}

func (m *DestinyDefinitionsDestinyTalentNodeDefinition) validateRandomActivationRequirement(formats strfmt.Registry) error {

	if swag.IsZero(m.RandomActivationRequirement) { // not required
		return nil
	}

	if m.RandomActivationRequirement != nil {

		if err := m.RandomActivationRequirement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("randomActivationRequirement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyTalentNodeDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyTalentNodeDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyTalentNodeDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
