# \ForumApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForumApproveFireteamThread**](ForumApi.md#ForumApproveFireteamThread) | **Post** /Forum/Recruit/Approve/{topicId}/ | 
[**ForumGetCoreTopicsPaged**](ForumApi.md#ForumGetCoreTopicsPaged) | **Get** /Forum/GetCoreTopicsPaged/{page}/{sort}/{quickDate}/{categoryFilter}/ | 
[**ForumGetForumTagSuggestions**](ForumApi.md#ForumGetForumTagSuggestions) | **Get** /Forum/GetForumTagSuggestions/ | 
[**ForumGetPoll**](ForumApi.md#ForumGetPoll) | **Get** /Forum/Poll/{topicId}/ | 
[**ForumGetPostAndParent**](ForumApi.md#ForumGetPostAndParent) | **Get** /Forum/GetPostAndParent/{childPostId}/ | 
[**ForumGetPostAndParentAwaitingApproval**](ForumApi.md#ForumGetPostAndParentAwaitingApproval) | **Get** /Forum/GetPostAndParentAwaitingApproval/{childPostId}/ | 
[**ForumGetPostsThreadedPaged**](ForumApi.md#ForumGetPostsThreadedPaged) | **Get** /Forum/GetPostsThreadedPaged/{parentPostId}/{page}/{pageSize}/{replySize}/{getParentPost}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetPostsThreadedPagedFromChild**](ForumApi.md#ForumGetPostsThreadedPagedFromChild) | **Get** /Forum/GetPostsThreadedPagedFromChild/{childPostId}/{page}/{pageSize}/{replySize}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetRecruitmentThreadSummaries**](ForumApi.md#ForumGetRecruitmentThreadSummaries) | **Post** /Forum/Recruit/Summaries/ | 
[**ForumGetTopicForContent**](ForumApi.md#ForumGetTopicForContent) | **Get** /Forum/GetTopicForContent/{contentId}/ | 
[**ForumGetTopicsPaged**](ForumApi.md#ForumGetTopicsPaged) | **Get** /Forum/GetTopicsPaged/{page}/{pageSize}/{group}/{sort}/{quickDate}/{categoryFilter}/ | 
[**ForumJoinFireteamThread**](ForumApi.md#ForumJoinFireteamThread) | **Post** /Forum/Recruit/Join/{topicId}/ | 
[**ForumKickBanFireteamApplicant**](ForumApi.md#ForumKickBanFireteamApplicant) | **Post** /Forum/Recruit/KickBan/{topicId}/{targetMembershipId}/ | 
[**ForumLeaveFireteamThread**](ForumApi.md#ForumLeaveFireteamThread) | **Post** /Forum/Recruit/Leave/{topicId}/ | 


# **ForumApproveFireteamThread**
> InlineResponse20010 ForumApproveFireteamThread(ctx, topicId)


Allows the owner of a fireteam thread to approve all joined members and start a private message conversation with them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **topicId** | **int64**| The post id of the recruitment topic to approve. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetCoreTopicsPaged**
> InlineResponse2006 ForumGetCoreTopicsPaged(categoryFilter, page, quickDate, sort, optional)


Gets a listing of all topics marked as part of the core group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **categoryFilter** | **int32**| The category filter. | 
  **page** | **int32**| Zero base page | 
  **quickDate** | **int32**| The date filter. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryFilter** | **int32**| The category filter. | 
 **page** | **int32**| Zero base page | 
 **quickDate** | **int32**| The date filter. | 
 **sort** | **int32**| The sort mode. | 
 **locales** | **string**| Comma seperated list of locales posts must match to return in the result list. Default &#39;en&#39; | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetForumTagSuggestions**
> InlineResponse2008 ForumGetForumTagSuggestions(optional)


Gets tag suggestions based on partial text entry, matching them with other tags previously used in the forums.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partialtag** | **string**| The partial tag input to generate suggestions from. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPoll**
> InlineResponse2006 ForumGetPoll(topicId)


Gets the specified forum poll.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **topicId** | **int64**| The post id of the topic that has the poll. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParent**
> InlineResponse2006 ForumGetPostAndParent(childPostId, optional)


Returns the post specified and its immediate parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **childPostId** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childPostId** | **int32**|  | 
 **showbanned** | **string**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParentAwaitingApproval**
> InlineResponse2006 ForumGetPostAndParentAwaitingApproval(childPostId, optional)


Returns the post specified and its immediate parent of posts that are awaiting approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **childPostId** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childPostId** | **int32**|  | 
 **showbanned** | **string**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPaged**
> InlineResponse2006 ForumGetPostsThreadedPaged(getParentPost, page, pageSize, parentPostId, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts at the given parent, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **getParentPost** | **bool**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **parentPostId** | **int32**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getParentPost** | **bool**|  | 
 **page** | **int32**|  | 
 **pageSize** | **int32**|  | 
 **parentPostId** | **int32**|  | 
 **replySize** | **int32**|  | 
 **rootThreadMode** | **bool**|  | 
 **sortMode** | **int32**|  | 
 **showbanned** | **string**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPagedFromChild**
> InlineResponse2006 ForumGetPostsThreadedPagedFromChild(childPostId, page, pageSize, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts starting at the topicId of the input childPostId, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **childPostId** | **int32**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childPostId** | **int32**|  | 
 **page** | **int32**|  | 
 **pageSize** | **int32**|  | 
 **replySize** | **int32**|  | 
 **rootThreadMode** | **bool**|  | 
 **sortMode** | **int32**|  | 
 **showbanned** | **string**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetRecruitmentThreadSummaries**
> InlineResponse20011 ForumGetRecruitmentThreadSummaries()


Allows the caller to get a list of to 25 recruitment thread summary information objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicForContent**
> InlineResponse2007 ForumGetTopicForContent(contentId)


Gets the post Id for the given content item's comments, if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **contentId** | **int64**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicsPaged**
> InlineResponse2006 ForumGetTopicsPaged(categoryFilter, group, page, pageSize, quickDate, sort, optional)


Get topics from any forum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **categoryFilter** | **int32**| A category filter | 
  **group** | **int64**| The group, if any. | 
  **page** | **int32**| Zero paged page number | 
  **pageSize** | **int32**| Unused | 
  **quickDate** | **int32**| A date filter. | 
  **sort** | **int32**| The sort mode. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryFilter** | **int32**| A category filter | 
 **group** | **int64**| The group, if any. | 
 **page** | **int32**| Zero paged page number | 
 **pageSize** | **int32**| Unused | 
 **quickDate** | **int32**| A date filter. | 
 **sort** | **int32**| The sort mode. | 
 **locales** | **string**| Comma seperated list of locales posts must match to return in the result list. Default &#39;en&#39; | 
 **tagstring** | **string**| The tags to search, if any. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumJoinFireteamThread**
> InlineResponse2009 ForumJoinFireteamThread(ctx, topicId)


Allows a user to slot themselves into a recruitment thread fireteam slot. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **topicId** | **int64**| The post id of the recruitment topic you wish to join. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumKickBanFireteamApplicant**
> InlineResponse2009 ForumKickBanFireteamApplicant(ctx, targetMembershipId, topicId)


Allows a recruitment thread owner to kick a join user from the fireteam. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **targetMembershipId** | **int64**| The id of the user you wish to kick. | 
  **topicId** | **int64**| The post id of the recruitment topic you wish to join. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumLeaveFireteamThread**
> InlineResponse2009 ForumLeaveFireteamThread(ctx, topicId)


Allows a user to remove themselves from a recruitment thread fireteam slot. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **topicId** | **int64**| The post id of the recruitment topic you wish to leave. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

