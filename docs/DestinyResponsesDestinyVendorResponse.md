# DestinyResponsesDestinyVendorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | [***interface{}**](interface{}.md) | The base properties of the vendor.  COMPONENT TYPE: Vendors | [optional] [default to null]
**Categories** | [***interface{}**](interface{}.md) | Categories that the vendor has available, and references to the sales therein.  COMPONENT TYPE: VendorCategories | [optional] [default to null]
**Sales** | [***interface{}**](interface{}.md) | Sales, keyed by the vendorItemIndex of the item being sold.  COMPONENT TYPE: VendorSales | [optional] [default to null]
**Items** | [***interface{}**](interface{}.md) | Item components, keyed by the vendorItemIndex of the active sale items.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


