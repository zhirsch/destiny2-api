# DestinyDefinitionsDestinyItemInventoryBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackUniqueLabel** | **string** | If this string is populated, you can&#39;t have more than one stack with this label in a given inventory. Note that this is different from the equipping block&#39;s unique label, which is used for equipping uniqueness. | [optional] [default to null]
**MaxStackSize** | **int32** | The maximum quantity of this item that can exist in a stack. | [optional] [default to null]
**BucketTypeHash** | **int32** | The hash identifier for the DestinyInventoryBucketDefinition to which this item belongs. I should have named this \&quot;bucketHash\&quot;, but too many things refer to it now. Sigh. | [optional] [default to null]
**RecoveryBucketTypeHash** | **int32** | If the item is picked up by the lost loot queue, this is the hash identifier for the DestinyInventoryBucketDefinition into which it will be placed. Again, I should have named this recoveryBucketHash instead. | [optional] [default to null]
**TierTypeHash** | **int32** | The hash identifier for the Tier Type of the item, use to look up its DestinyItemTierTypeDefinition if you need to show localized data for the item&#39;s tier. | [optional] [default to null]
**IsInstanceItem** | **bool** | If TRUE, this item is instanced. Otherwise, it is a generic item that merely has a quantity in a stack (like Glimmer). | [optional] [default to null]
**TierTypeName** | **string** | The localized name of the tier type, which is a useful shortcut so you don&#39;t have to look up the definition every time. However, it&#39;s mostly a holdover from days before we had a DestinyItemTierTypeDefinition to refer to. | [optional] [default to null]
**TierType** | [***interface{}**](interface{}.md) | The enumeration matching the tier type of the item to known values, again for convenience sake. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


