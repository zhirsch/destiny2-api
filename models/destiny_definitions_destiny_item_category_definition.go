// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyDefinitionsDestinyItemCategoryDefinition In an attempt to categorize items by type, usage, and other interesting properties, we created DestinyItemCategoryDefinition: information about types that is assembled using a set of heuristics that examine the properties of an item such as what inventory bucket it's in, its item type name, and whether it has or is missing certain blocks of data.
// This heuristic is imperfect, however. If you find an item miscategorized, let us know on the Bungie API forums!
// We then populate all of the categories that we think an item belongs to in its DestinyInventoryItemDefinition.itemCategoryHashes property. You can use that to provide your own custom item filtering, sorting, aggregating... go nuts on it! And let us know if you see more categories that you wish would be added!
// swagger:model Destiny.Definitions.DestinyItemCategoryDefinition

type DestinyDefinitionsDestinyItemCategoryDefinition struct {

	// display properties
	DisplayProperties *DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`

	// If an item belongs to this category, it will also get this class restriction enum value.
	// See the other "grant"-prefixed properties on this definition for my color commentary.
	GrantDestinyClass DestinyDestinyClass `json:"grantDestinyClass,omitempty"`

	// If an item belongs to this category, it will also receive this item type. This is now how DestinyItemType is populated for items: it used to be an even jankier process, but that's a story that requires more alcohol.
	GrantDestinyItemType DestinyDestinyItemType `json:"grantDestinyItemType,omitempty"`

	// If an item belongs to this category, it will also receive this subtype enum value.
	// I know what you're thinking - what if it belongs to multiple categories that provide sub-types?
	// The last one processed wins, as is the case with all of these "grant" enums. Now you can see one reason why we moved away from these enums... but they're so convenient when they work, aren't they?
	GrantDestinySubType DestinyDestinyItemSubType `json:"grantDestinySubType,omitempty"`

	// If this category is a "parent" category of other categories, those children will have their hashes listed in rendering order here, and can be looked up using these hashes against DestinyItemCategoryDefinition.
	// In this way, you can build up a visual hierarchy of item categories. That's what we did, and you can do it too. I believe in you. Yes, you, Carl.
	// (I hope someone named Carl reads this someday)
	GroupedCategoryHashes []uint32 `json:"groupedCategoryHashes"`

	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
	// When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash uint32 `json:"hash,omitempty"`

	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`

	// The janky regular expression we used against the item type to try and discern whether the item belongs to this category.
	ItemTypeRegex string `json:"itemTypeRegex,omitempty"`

	// If the item type matches this janky regex, it does *not* belong to this category.
	ItemTypeRegexNot string `json:"itemTypeRegexNot,omitempty"`

	// If the item belongs to this bucket, it does belong to this category.
	OriginBucketIdentifier string `json:"originBucketIdentifier,omitempty"`

	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`

	// A shortened version of the title. The reason why we have this is because the Armory in German had titles that were too long to display in our UI, so these were localized abbreviated versions of those categories. The property still exists today, even though the Armory doesn't exist for D2... yet.
	ShortTitle string `json:"shortTitle,omitempty"`

	// If True, this category should be visible in UI. Sometimes we make categories that we don't think are interesting externally. It's up to you if you want to skip on showing them.
	Visible bool `json:"visible,omitempty"`
}

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition displayProperties false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition grantDestinyClass false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition grantDestinyItemType false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition grantDestinySubType false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition groupedCategoryHashes false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition hash false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition index false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition itemTypeRegex false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition itemTypeRegexNot false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition originBucketIdentifier false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition redacted false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition shortTitle false */

/* polymorph Destiny.Definitions.DestinyItemCategoryDefinition visible false */

// Validate validates this destiny definitions destiny item category definition
func (m *DestinyDefinitionsDestinyItemCategoryDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGrantDestinyClass(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGrantDestinyItemType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGrantDestinySubType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupedCategoryHashes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DestinyDefinitionsDestinyItemCategoryDefinition) validateDisplayProperties(formats strfmt.Registry) error {

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

func (m *DestinyDefinitionsDestinyItemCategoryDefinition) validateGrantDestinyClass(formats strfmt.Registry) error {

	if swag.IsZero(m.GrantDestinyClass) { // not required
		return nil
	}

	if err := m.GrantDestinyClass.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grantDestinyClass")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyItemCategoryDefinition) validateGrantDestinyItemType(formats strfmt.Registry) error {

	if swag.IsZero(m.GrantDestinyItemType) { // not required
		return nil
	}

	if err := m.GrantDestinyItemType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grantDestinyItemType")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyItemCategoryDefinition) validateGrantDestinySubType(formats strfmt.Registry) error {

	if swag.IsZero(m.GrantDestinySubType) { // not required
		return nil
	}

	if err := m.GrantDestinySubType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grantDestinySubType")
		}
		return err
	}

	return nil
}

func (m *DestinyDefinitionsDestinyItemCategoryDefinition) validateGroupedCategoryHashes(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupedCategoryHashes) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemCategoryDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyDefinitionsDestinyItemCategoryDefinition) UnmarshalBinary(b []byte) error {
	var res DestinyDefinitionsDestinyItemCategoryDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
