/* 
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@bungie.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package api




// Items like Sacks or Boxes can have items that it shows in-game when you view details that represent the items you can obtain if you use or acquire the item.  This defines those categories, and gives some insights into that data's source.
type DestinyDefinitionsDestinyItemPreviewBlockDefinition struct {
	// If the preview data is derived from a fake \"Preview\" Vendor, this will be the hash identifier for the DestinyVendorDefinition of that fake vendor.
	PreviewVendorHash int32 `json:"previewVendorHash,omitempty"`
	// If the preview has an associated action (like \"Open\"), this will be the localized string for that action.
	PreviewActionString string `json:"previewActionString,omitempty"`
	// This is a list of the items being previewed, categorized in the same way as they are in the preview UI.
	DerivedItemCategories []DestinyDefinitionsItemsDestinyDerivedItemCategoryDefinition `json:"derivedItemCategories,omitempty"`
}

