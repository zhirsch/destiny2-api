# \UserApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGetAvailableThemes**](UserApi.md#UserGetAvailableThemes) | **Get** /User/GetAvailableThemes/ | 
[**UserGetBungieNetUserById**](UserApi.md#UserGetBungieNetUserById) | **Get** /User/GetBungieNetUserById/{id}/ | 
[**UserGetMembershipDataById**](UserApi.md#UserGetMembershipDataById) | **Get** /User/GetMembershipsById/{membershipId}/{membershipType}/ | 
[**UserGetMembershipDataForCurrentUser**](UserApi.md#UserGetMembershipDataForCurrentUser) | **Get** /User/GetMembershipsForCurrentUser/ | 
[**UserGetPartnerships**](UserApi.md#UserGetPartnerships) | **Get** /User/{membershipId}/Partnerships/ | 
[**UserGetUserAliases**](UserApi.md#UserGetUserAliases) | **Get** /User/GetUserAliases/{id}/ | 
[**UserSearchUsers**](UserApi.md#UserSearchUsers) | **Get** /User/SearchUsers/ | 


# **UserGetAvailableThemes**
> InlineResponse2003 UserGetAvailableThemes()


Returns a list of all available user themes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetBungieNetUserById**
> InlineResponse200 UserGetBungieNetUserById(id)


Loads a bungienet user by membership id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int64**| The requested Bungie.net membership id. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataById**
> InlineResponse2004 UserGetMembershipDataById(membershipId, membershipType)


Returns a list of accounts associated with the supplied membership ID and membership type. This will include all linked accounts (even when hidden) if supplied credentials permit it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **membershipId** | **int64**| The membership ID of the target user. | 
  **membershipType** | **int32**| Type of the supplied membership ID. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataForCurrentUser**
> InlineResponse2004 UserGetMembershipDataForCurrentUser(ctx, )


Returns a list of accounts associated with signed in user. This is useful for OAuth implementations that do not give you access to the token response.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetPartnerships**
> InlineResponse2005 UserGetPartnerships(membershipId)


Returns a user's linked Partnerships.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **membershipId** | **int64**| The ID of the member for whom partnerships should be returned. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetUserAliases**
> InlineResponse2001 UserGetUserAliases(id)


Loads aliases of a bungienet membership id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int64**| The requested Bungie.net membership id. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSearchUsers**
> InlineResponse2002 UserSearchUsers(optional)


Returns a list of possible users based on the search string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**| The search string. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

