# SearchResultOfDestinyEntitySearchResultItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]DestinyDefinitionsDestinyEntitySearchResultItem**](Destiny.Definitions.DestinyEntitySearchResultItem.md) |  | [optional] [default to null]
**TotalResults** | **int32** |  | [optional] [default to null]
**HasMore** | **bool** |  | [optional] [default to null]
**Query** | [***QueriesPagedQuery**](Queries.PagedQuery.md) |  | [optional] [default to null]
**ReplacementContinuationToken** | **string** |  | [optional] [default to null]
**UseTotalResults** | **bool** | If useTotalResults is true, then totalResults represents an accurate count.  If False, it does not, and may be estimated/only the size of the current page.  Either way, you should probably always only trust hasMore.  This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


