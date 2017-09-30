# DestinyVendorsDestinyVendorReceipt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyPaid** | [**[]DestinyDestinyItemQuantity**](Destiny.DestinyItemQuantity.md) | The amount paid for the item, in terms of items that were consumed in the purchase and their quantity. | [optional] [default to null]
**ItemReceived** | [***interface{}**](interface{}.md) | The item that was received, and its quantity. | [optional] [default to null]
**LicenseUnlockHash** | **int32** | The unlock flag used to determine whether you still have the purchased item. | [optional] [default to null]
**PurchasedByCharacterId** | **int64** | The ID of the character who made the purchase. | [optional] [default to null]
**RefundPolicy** | [***interface{}**](interface{}.md) | Whether you can get a refund, and what happens in order for the refund to be received. See the DestinyVendorItemRefundPolicy enum for details. | [optional] [default to null]
**SequenceNumber** | **int32** | The identifier of this receipt. | [optional] [default to null]
**TimeToExpiration** | **int64** | The seconds since epoch at which this receipt is rendered invalid. | [optional] [default to null]
**ExpiresOn** | [**time.Time**](time.Time.md) | The date at which this receipt is rendered invalid. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


