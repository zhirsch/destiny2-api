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




// This is a reference to, and summary data for, a specific item that you can get as a result of Using or Acquiring some other Item (For example, this could be summary information for an Emote that you can get by opening an an Eververse Box) See DestinyDerivedItemCategoryDefinition for more information.
type DestinyDefinitionsItemsDestinyDerivedItemDefinition struct {
	// The hash for the DestinyInventoryItemDefinition of this derived item, if there is one. Sometimes we are given this information as a manual override, in which case there won't be an actual DestinyInventoryItemDefinition for what we display, but you can still show the strings from this object itself.
	ItemHash int32 `json:"itemHash,omitempty"`
	// The name of the derived item.
	ItemName string `json:"itemName,omitempty"`
	// Additional details about the derived item, in addition to the description.
	ItemDetail string `json:"itemDetail,omitempty"`
	// A brief description of the item.
	ItemDescription string `json:"itemDescription,omitempty"`
	// An icon for the item.
	IconPath string `json:"iconPath,omitempty"`
	// If the item was derived from a \"Preview Vendor\", this will be an index into the DestinyVendorDefinition's itemList property. Otherwise, -1.
	VendorItemIndex int32 `json:"vendorItemIndex,omitempty"`
}
