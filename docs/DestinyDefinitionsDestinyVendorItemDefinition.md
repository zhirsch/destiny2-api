# DestinyDefinitionsDestinyVendorItemDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorItemIndex** | **int32** | The index into the DestinyVendorDefinition.saleList. This is what we use to refer to items being sold throughout live and definition data. | [optional] [default to null]
**ItemHash** | **int32** | The hash identifier of the item being sold (DestinyInventoryItemDefinition).  Note that a vendor can sell the same item in multiple ways, so don&#39;t assume that itemHash is a unique identifier for this entity. | [optional] [default to null]
**Quantity** | **int32** | The amount you will recieve of the item described in itemHash if you make the purchase. | [optional] [default to null]
**FailureIndexes** | **[]int32** | An list of indexes into the DestinyVendorDefinition.failureStrings array, indicating the possible failure strings that can be relevant for this item. | [optional] [default to null]
**Currencies** | [**[]DestinyDestinyItemQuantity**](Destiny.DestinyItemQuantity.md) | This is a pre-compiled aggregation of item value and priceOverrideList, so that we have one place to check for what the purchaser must pay for the item. Use this instead of trying to piece together the price separately. | [optional] [default to null]
**RefundPolicy** | [***interface{}**](interface{}.md) | If this item can be refunded, this is the policy for what will be refundd, how, and in what time period. | [optional] [default to null]
**RefundTimeLimit** | **int32** | The amount of time before refundability of the newly purchased item will expire. | [optional] [default to null]
**CreationLevels** | [**[]DestinyDefinitionsDestinyItemCreationEntryLevelDefinition**](Destiny.Definitions.DestinyItemCreationEntryLevelDefinition.md) | The Default level at which the item will spawn. Almost always driven by an adjusto these days. Ideally should be singular. It&#39;s a long story how this ended up as a list, but there is always either going to be 0:1 of these entities. | [optional] [default to null]
**DisplayCategoryIndex** | **int32** | This is an index specifically into the display category, as opposed to the server-side Categories (which do not need to match or pair with each other in any way: server side categories are really just structures for common validation. Display Category will let us more easily categorize items visually) | [optional] [default to null]
**CategoryIndex** | **int32** | The index into the DestinyVendorDefinition.categories array, so you can find the category associated with this item. | [optional] [default to null]
**OriginalCategoryIndex** | **int32** | Same as above, but for the original category indexes. | [optional] [default to null]
**MinimumLevel** | **int32** | The minimum character level at which this item is available for sale. | [optional] [default to null]
**MaximumLevel** | **int32** | The maximum character level at which this item is available for sale. | [optional] [default to null]
**Action** | [***interface{}**](interface{}.md) | The action to be performed when purchasing the item, if it&#39;s not just \&quot;buy\&quot;. | [optional] [default to null]
**DisplayCategory** | **string** | The string identifier for the category selling this item. | [optional] [default to null]
**InventoryBucketHash** | **int32** | The inventory bucket into which this item will be placed upon purchase. | [optional] [default to null]
**VisibilityScope** | [***interface{}**](interface{}.md) | The most restrictive scope that determines whether the item is available in the Vendor&#39;s inventory. See DestinyGatingScope&#39;s documentation for more information.  This can be determined by Unlock gating, or by whether or not the item has purchase level requirements (minimumLevel and maximumLevel properties). | [optional] [default to null]
**PurchasableScope** | [***interface{}**](interface{}.md) | Similar to visibilityScope, it represents the most restrictive scope that determines whether the item can be purchased. It will at least be as restrictive as visibilityScope, but could be more restrictive if the item has additional purchase requirements beyond whether it is merely visible or not.  See DestinyGatingScope&#39;s documentation for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


