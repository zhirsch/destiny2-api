# DestinyDefinitionsDestinyVendorInventoryFlyoutDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockedDescription** | **string** | If the flyout is locked, this is the reason why. | [optional] [default to null]
**DisplayProperties** | [***interface{}**](interface{}.md) | The title and other common properties of the flyout. | [optional] [default to null]
**Buckets** | [**[]DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition**](Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition.md) | A list of inventory buckets and other metadata to show on the screen. | [optional] [default to null]
**FlyoutId** | **int32** | An identifier for the flyout, in case anything else needs to refer to them. | [optional] [default to null]
**SuppressNewness** | **bool** | If this is true, don&#39;t show any of the glistening \&quot;this is a new item\&quot; UI elements, like we show on the inventory items themselves in in-game UI. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


