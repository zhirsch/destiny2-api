# DestinyDefinitionsItemsDestinyDerivedItemDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemHash** | **int32** | The hash for the DestinyInventoryItemDefinition of this derived item, if there is one. Sometimes we are given this information as a manual override, in which case there won&#39;t be an actual DestinyInventoryItemDefinition for what we display, but you can still show the strings from this object itself. | [optional] [default to null]
**ItemName** | **string** | The name of the derived item. | [optional] [default to null]
**ItemDetail** | **string** | Additional details about the derived item, in addition to the description. | [optional] [default to null]
**ItemDescription** | **string** | A brief description of the item. | [optional] [default to null]
**IconPath** | **string** | An icon for the item. | [optional] [default to null]
**VendorItemIndex** | **int32** | If the item was derived from a \&quot;Preview Vendor\&quot;, this will be an index into the DestinyVendorDefinition&#39;s itemList property. Otherwise, -1. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


