# \TrendingApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrendingGetTrendingCategories**](TrendingApi.md#TrendingGetTrendingCategories) | **Get** /Trending/Categories/ | 
[**TrendingGetTrendingCategory**](TrendingApi.md#TrendingGetTrendingCategory) | **Get** /Trending/Categories/{categoryId}/{pageNumber}/ | 
[**TrendingGetTrendingEntryDetail**](TrendingApi.md#TrendingGetTrendingEntryDetail) | **Get** /Trending/Details/{trendingEntryType}/{identifier}/ | 


# **TrendingGetTrendingCategories**
> InlineResponse20052 TrendingGetTrendingCategories()


Returns trending items for Bungie.net, collapsed into the first page of items per category. For pagination within a category, call GetTrendingCategory.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrendingGetTrendingCategory**
> InlineResponse20053 TrendingGetTrendingCategory(categoryId, pageNumber)


Returns paginated lists of trending items for a category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **categoryId** | **string**| The ID of the category for whom you want additional results. | 
  **pageNumber** | **int32**| The page # of results to return. | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrendingGetTrendingEntryDetail**
> InlineResponse20054 TrendingGetTrendingEntryDetail(identifier, trendingEntryType)


Returns the detailed results for a specific trending entry. Note that trending entries are uniquely identified by a combination of *both* the TrendingEntryType *and* the identifier: the identifier alone is not guaranteed to be globally unique.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **identifier** | **string**| The identifier for the entity to be returned. | 
  **trendingEntryType** | **int32**| The type of entity to be returned. | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

