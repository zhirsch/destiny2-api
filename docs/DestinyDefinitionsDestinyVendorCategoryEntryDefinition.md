# DestinyDefinitionsDestinyVendorCategoryEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIndex** | **int32** | The index of the category in the original category definitions for the vendor. | [optional] [default to null]
**CategoryId** | **string** | The string identifier of the category. | [optional] [default to null]
**CategoryHash** | **int32** | The hashed identifier for the category. (note that this is NOT pointing to a DestinyVendorCategoryDefinition, it&#39;s confusing but this is a sale item category in a vendor, not a categorization of vendors themselves) | [optional] [default to null]
**QuantityAvailable** | **int32** | The amount of items that will be available when this category is shown. | [optional] [default to null]
**ShowUnavailableItems** | **bool** | If items aren&#39;t up for sale in this category, should we still show them (greyed out)? | [optional] [default to null]
**HideIfNoCurrency** | **bool** | If you don&#39;t have the currency required to buy items from this category, should the items be hidden? | [optional] [default to null]
**HideFromRegularPurchase** | **bool** | True if this category doesn&#39;t allow purchases. | [optional] [default to null]
**BuyStringOverride** | **string** | The localized string for making purchases from this category, if it is different from the vendor&#39;s string for purchasing. | [optional] [default to null]
**DisabledDescription** | **string** | If the category is disabled, this is the localized description to show. | [optional] [default to null]
**DisplayTitle** | **string** | The localized title of the category. | [optional] [default to null]
**Overlay** | [***interface{}**](interface{}.md) | If this category has an overlay prompt that should appear, this contains the details of that prompt. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


