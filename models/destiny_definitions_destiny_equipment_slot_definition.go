// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyEquipmentSlotDefinition Characters can not only have Inventory buckets (containers of items that are generally matched by their type or functionality), they can also have Equipment Slots.
// The Equipment Slot is an indicator that the related bucket can have instanced items equipped on the character. For instance, the Primary Weapon bucket has an Equipment Slot that determines whether you can equip primary weapons, and holds the association between its slot and the inventory bucket from which it can have items equipped.
// An Equipment Slot must have a related Inventory Bucket, but not all inventory buckets must have Equipment Slots.
// swagger:model Destiny.Definitions.DestinyEquipmentSlotDefinition

type DestinyDefinitionsDestinyEquipmentSlotDefinition struct {

	// If True, equipped items should have their custom art dyes applied when rendering the item. Otherwise, custom art dyes on an item should be ignored if the item is equipped in this slot.
	ApplyCustomArtDyes bool `json:"applyCustomArtDyes,omitempty"`

	// The inventory bucket that owns this equipment slot.
	BucketTypeHash uint32 `json:"bucketTypeHash,omitempty"`

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// These technically point to "Equipment Category Definitions". But don't get excited. There's nothing of significant value in those definitions, so I didn't bother to expose them. You can use the hash here to group equipment slots by common functionality, which serves the same purpose as if we had the Equipment Category definitions exposed.
	EquipmentCategoryHash uint32 `json:"equipmentCategoryHash,omitempty"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition applyCustomArtDyes false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition bucketTypeHash false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition equipmentCategoryHash false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition hash false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition index false */

/* polymorph Destiny.Definitions.DestinyEquipmentSlotDefinition redacted false */

// Validate validates this destiny definitions destiny equipment slot definition
func (m *DestinyDefinitionsDestinyEquipmentSlotDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyEquipmentSlotDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyEquipmentSlotDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyEquipmentSlotDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyEquipmentSlotDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
