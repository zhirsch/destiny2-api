// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DestinyEntitiesItemsDestinyItemRenderComponent Many items can be rendered in 3D. When you request this block, you will obtain the custom data needed to render this specific instance of the item.
// swagger:model Destiny.Entities.Items.DestinyItemRenderComponent

type DestinyEntitiesItemsDestinyItemRenderComponent struct {

	// A dictionary for rendering gear components, with:
	// key = Art Arrangement Region Index
	// value = The chosen Arrangement Index for the Region, based on the value of a stat on the item used for making the choice.
	ArtRegions map[string]int32 `json:"artRegions,omitempty"`

	// If you should use custom dyes on this item, it will be indicated here.
	UseCustomDyes bool `json:"useCustomDyes,omitempty"`
}

/* polymorph Destiny.Entities.Items.DestinyItemRenderComponent artRegions false */

/* polymorph Destiny.Entities.Items.DestinyItemRenderComponent useCustomDyes false */

// Validate validates this destiny entities items destiny item render component
func (m *DestinyEntitiesItemsDestinyItemRenderComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemRenderComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DestinyEntitiesItemsDestinyItemRenderComponent) UnmarshalBinary(b []byte) error {
	var res DestinyEntitiesItemsDestinyItemRenderComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}