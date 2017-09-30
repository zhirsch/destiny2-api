# \CommunityContentApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunityContentGetCommunityContent**](CommunityContentApi.md#CommunityContentGetCommunityContent) | **Get** /CommunityContent/Get/{sort}/{mediaFilter}/{page}/ | 
[**CommunityContentGetCommunityLiveStatuses**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatuses) | **Get** /CommunityContent/Live/All/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetCommunityLiveStatusesForClanmates**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatusesForClanmates) | **Get** /CommunityContent/Live/Clan/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetCommunityLiveStatusesForFriends**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatusesForFriends) | **Get** /CommunityContent/Live/Friends/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetFeaturedCommunityLiveStatuses**](CommunityContentApi.md#CommunityContentGetFeaturedCommunityLiveStatuses) | **Get** /CommunityContent/Live/Featured/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetStreamingStatusForMember**](CommunityContentApi.md#CommunityContentGetStreamingStatusForMember) | **Get** /CommunityContent/Live/Users/{partnershipType}/{membershipType}/{membershipId}/ | 


# **CommunityContentGetCommunityContent**
> InlineResponse2006 CommunityContentGetCommunityContent(mediaFilter, page, sort)


Returns community content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **mediaFilter** | **int32**| The type of media to get | 
  **page** | **int32**| Zero based page | 
  **sort** | **int32**| The sort mode. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatuses**
> InlineResponse20050 CommunityContentGetCommunityLiveStatuses(page, partnershipType, sort, optional)


Returns info about community members who are live streaming.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **page** | **int32**| Zero based page. | 
  **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Zero based page. | 
 **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
 **sort** | **int32**| The sort mode. | 
 **modeHash** | **int32**| The hash of the Activity Mode for which streams should be retrieved. Don&#39;t pass it in or pass 0 to not filter by mode. | 
 **streamLocale** | **string**| The locale for streams you&#39;d like to see. Don&#39;t pass this to fall back on your BNet locale. Pass &#39;ALL&#39; to not filter by locale. | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatusesForClanmates**
> InlineResponse20050 CommunityContentGetCommunityLiveStatusesForClanmates(page, partnershipType, sort)


Returns info about community members who are live streaming in your clans.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **page** | **int32**| Zero based page. | 
  **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
  **sort** | **int32**| The sort mode. | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatusesForFriends**
> InlineResponse20050 CommunityContentGetCommunityLiveStatusesForFriends(page, partnershipType, sort)


Returns info about community members who are live streaming among your friends.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **page** | **int32**| Zero based page. | 
  **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
  **sort** | **int32**| The sort mode. | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetFeaturedCommunityLiveStatuses**
> InlineResponse20050 CommunityContentGetFeaturedCommunityLiveStatuses(page, partnershipType, sort, optional)


Returns info about Featured live streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **page** | **int32**| Zero based page. | 
  **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Zero based page. | 
 **partnershipType** | **int32**| The type of partnership for which the status should be returned. | 
 **sort** | **int32**| The sort mode. | 
 **streamLocale** | **string**| The locale for streams you&#39;d like to see. Don&#39;t pass this to fall back on your BNet locale. Pass &#39;ALL&#39; to not filter by locale. | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetStreamingStatusForMember**
> InlineResponse20051 CommunityContentGetStreamingStatusForMember(membershipId, membershipType, partnershipType)


Gets the Live Streaming status of a particular Account and Membership Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **membershipId** | **int64**| The membershipId related to that type. | 
  **membershipType** | **int32**| The type of account for which info will be extracted. | 
  **partnershipType** | **int32**| The type of partnership for which info will be extracted. | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

