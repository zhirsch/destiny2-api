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




// This represents an item being sold by the vendor.
type DestinyDefinitionsDestinyVendorItemDefinition struct {
	// The index into the DestinyVendorDefinition.saleList. This is what we use to refer to items being sold throughout live and definition data.
	VendorItemIndex int32 `json:"vendorItemIndex,omitempty"`
	// The hash identifier of the item being sold (DestinyInventoryItemDefinition).  Note that a vendor can sell the same item in multiple ways, so don't assume that itemHash is a unique identifier for this entity.
	ItemHash int32 `json:"itemHash,omitempty"`
	// The amount you will recieve of the item described in itemHash if you make the purchase.
	Quantity int32 `json:"quantity,omitempty"`
	// An list of indexes into the DestinyVendorDefinition.failureStrings array, indicating the possible failure strings that can be relevant for this item.
	FailureIndexes []int32 `json:"failureIndexes,omitempty"`
	// This is a pre-compiled aggregation of item value and priceOverrideList, so that we have one place to check for what the purchaser must pay for the item. Use this instead of trying to piece together the price separately.
	Currencies []DestinyDestinyItemQuantity `json:"currencies,omitempty"`
	// If this item can be refunded, this is the policy for what will be refundd, how, and in what time period.
	RefundPolicy *interface{} `json:"refundPolicy,omitempty"`
	// The amount of time before refundability of the newly purchased item will expire.
	RefundTimeLimit int32 `json:"refundTimeLimit,omitempty"`
	// The Default level at which the item will spawn. Almost always driven by an adjusto these days. Ideally should be singular. It's a long story how this ended up as a list, but there is always either going to be 0:1 of these entities.
	CreationLevels []DestinyDefinitionsDestinyItemCreationEntryLevelDefinition `json:"creationLevels,omitempty"`
	// This is an index specifically into the display category, as opposed to the server-side Categories (which do not need to match or pair with each other in any way: server side categories are really just structures for common validation. Display Category will let us more easily categorize items visually)
	DisplayCategoryIndex int32 `json:"displayCategoryIndex,omitempty"`
	// The index into the DestinyVendorDefinition.categories array, so you can find the category associated with this item.
	CategoryIndex int32 `json:"categoryIndex,omitempty"`
	// Same as above, but for the original category indexes.
	OriginalCategoryIndex int32 `json:"originalCategoryIndex,omitempty"`
	// The minimum character level at which this item is available for sale.
	MinimumLevel int32 `json:"minimumLevel,omitempty"`
	// The maximum character level at which this item is available for sale.
	MaximumLevel int32 `json:"maximumLevel,omitempty"`
	// The action to be performed when purchasing the item, if it's not just \"buy\".
	Action *interface{} `json:"action,omitempty"`
	// The string identifier for the category selling this item.
	DisplayCategory string `json:"displayCategory,omitempty"`
	// The inventory bucket into which this item will be placed upon purchase.
	InventoryBucketHash int32 `json:"inventoryBucketHash,omitempty"`
	// The most restrictive scope that determines whether the item is available in the Vendor's inventory. See DestinyGatingScope's documentation for more information.  This can be determined by Unlock gating, or by whether or not the item has purchase level requirements (minimumLevel and maximumLevel properties).
	VisibilityScope *interface{} `json:"visibilityScope,omitempty"`
	// Similar to visibilityScope, it represents the most restrictive scope that determines whether the item can be purchased. It will at least be as restrictive as visibilityScope, but could be more restrictive if the item has additional purchase requirements beyond whether it is merely visible or not.  See DestinyGatingScope's documentation for more information.
	PurchasableScope *interface{} `json:"purchasableScope,omitempty"`
}
