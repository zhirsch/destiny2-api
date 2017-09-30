# \PreviewApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destiny2ActivateTalentNode**](PreviewApi.md#Destiny2ActivateTalentNode) | **Post** /Destiny2/Actions/Items/ActivateTalentNode/ | 
[**Destiny2GetActivityHistory**](PreviewApi.md#Destiny2GetActivityHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/Activities/ | 
[**Destiny2GetClanAggregateStats**](PreviewApi.md#Destiny2GetClanAggregateStats) | **Get** /Destiny2/Stats/AggregateClanStats/{groupId}/ | 
[**Destiny2GetClanLeaderboards**](PreviewApi.md#Destiny2GetClanLeaderboards) | **Get** /Destiny2/Stats/Leaderboards/Clans/{groupId}/ | 
[**Destiny2GetDestinyAggregateActivityStats**](PreviewApi.md#Destiny2GetDestinyAggregateActivityStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/AggregateActivityStats/ | 
[**Destiny2GetHistoricalStats**](PreviewApi.md#Destiny2GetHistoricalStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/ | 
[**Destiny2GetHistoricalStatsForAccount**](PreviewApi.md#Destiny2GetHistoricalStatsForAccount) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/ | 
[**Destiny2GetLeaderboards**](PreviewApi.md#Destiny2GetLeaderboards) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/ | 
[**Destiny2GetLeaderboardsForCharacter**](PreviewApi.md#Destiny2GetLeaderboardsForCharacter) | **Get** /Destiny2/Stats/Leaderboards/{membershipType}/{destinyMembershipId}/{characterId}/ | 
[**Destiny2GetUniqueWeaponHistory**](PreviewApi.md#Destiny2GetUniqueWeaponHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/UniqueWeapons/ | 
[**Destiny2GetVendor**](PreviewApi.md#Destiny2GetVendor) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/{vendorHash}/ | 
[**Destiny2GetVendors**](PreviewApi.md#Destiny2GetVendors) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/ | 
[**Destiny2InsertSocketPlug**](PreviewApi.md#Destiny2InsertSocketPlug) | **Post** /Destiny2/Actions/Items/InsertSocketPlug/ | 
[**Destiny2SearchDestinyEntities**](PreviewApi.md#Destiny2SearchDestinyEntities) | **Get** /Destiny2/Armory/Search/{type}/{searchTerm}/ | 


# **Destiny2ActivateTalentNode**
> InlineResponse20015 Destiny2ActivateTalentNode(ctx, )


Activate a Talent Node. Chill out, everyone: we haven't decided yet whether this will be able to activate nodes with costs, but if we do it will require special scope permission for an application attempting to do so. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. PREVIEW: This service is not actually implemented yet, but we are returning the planned schema of the endpoint for review, comment, and preparation for its eventual implementation.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetActivityHistory**
> InlineResponse20045 Destiny2GetActivityHistory(characterId, destinyMembershipId, membershipType, optional)


Gets activity history stats for indicated character. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The id of the character to retrieve. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The id of the character to retrieve. | 
 **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **count** | **int32**| Number of rows to return | 
 **mode** | **int32**| A filter for the activity mode to be returned. None returns all activities. See the documentation for DestinyActivityModeType for valid values, and pass in string representation. | 
 **page** | **int32**| Page number to return, starting with 0. | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanAggregateStats**
> InlineResponse20041 Destiny2GetClanAggregateStats(groupId, optional)


Gets aggregated stats for a clan using the same categories as the clan leaderboards. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **modes** | **string**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanLeaderboards**
> InlineResponse20040 Destiny2GetClanLeaderboards(groupId, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **maxtop** | **int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **string**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **string**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetDestinyAggregateActivityStats**
> InlineResponse20047 Destiny2GetDestinyAggregateActivityStats(characterId, destinyMembershipId, membershipType)


Gets all activities the character has participated in together with aggregate statistics for those activities. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The specific character whose activities should be returned. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetHistoricalStats**
> InlineResponse20043 Destiny2GetHistoricalStats(characterId, destinyMembershipId, membershipType, optional)


Gets historical stats for indicated character. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The id of the character to retrieve. You can omit this character ID or set it to 0 to get aggregate stats across all characters. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The id of the character to retrieve. You can omit this character ID or set it to 0 to get aggregate stats across all characters. | 
 **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **dayend** | **time.Time**| Last day to return when daily stats are requested. Use the format YYYY-MM-DD. | 
 **daystart** | **time.Time**| First day to return when daily stats are requested. Use the format YYYY-MM-DD | 
 **groups** | [**[]DestinyHistoricalStatsDefinitionsDestinyStatsGroupType**](DestinyHistoricalStatsDefinitionsDestinyStatsGroupType.md)| Group of stats to include, otherwise only general stats are returned. Comma separated list is allowed. Values: General, Weapons, Medals | 
 **modes** | [**[]DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](DestinyHistoricalStatsDefinitionsDestinyActivityModeType.md)| Game modes to return. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **periodType** | **int32**| Indicates a specific period type to return. Optional. May be: Daily, AllTime, or Activity | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetHistoricalStatsForAccount**
> InlineResponse20044 Destiny2GetHistoricalStatsForAccount(destinyMembershipId, membershipType, optional)


Gets aggregate historical stats organized around each character for a given account. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **groups** | [**[]DestinyHistoricalStatsDefinitionsDestinyStatsGroupType**](DestinyHistoricalStatsDefinitionsDestinyStatsGroupType.md)| Groups of stats to include, otherwise only general stats are returned. Comma separated list is allowed. Values: General, Weapons, Medals. | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboards**
> InlineResponse20040 Destiny2GetLeaderboards(destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint has not yet been implemented. It is being returned for a preview of future functionality, and for public comment/suggestion/preparation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **maxtop** | **int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **string**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **string**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboardsForCharacter**
> InlineResponse20040 Destiny2GetLeaderboardsForCharacter(characterId, destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The specific character to build the leaderboard around for the provided Destiny Membership. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The specific character to build the leaderboard around for the provided Destiny Membership. | 
 **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **maxtop** | **int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **string**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **string**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetUniqueWeaponHistory**
> InlineResponse20046 Destiny2GetUniqueWeaponHistory(characterId, destinyMembershipId, membershipType)


Gets details about unique weapon usage, including all exotic weapons. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The id of the character to retrieve. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetVendor**
> InlineResponse20036 Destiny2GetVendor(characterId, destinyMembershipId, membershipType, vendorHash, optional)


Get the details of a specific Vendor. PREVIEW: This service is not yet active, but we are returning the planned schema of the endpoint for review, comment, and preparation for its eventual implementation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
  **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
  **vendorHash** | **int32**| The Hash identifier of the Vendor to be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
 **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **vendorHash** | **int32**| The Hash identifier of the Vendor to be returned. | 
 **components** | [**[]DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetVendors**
> InlineResponse20035 Destiny2GetVendors(characterId, destinyMembershipId, membershipType, optional)


Get currently available vendors. PREVIEW: This service is not yet active, but we are returning the planned schema of the endpoint for review, comment, and preparation for its eventual implementation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
  **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
 **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **components** | [**[]DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2InsertSocketPlug**
> InlineResponse20015 Destiny2InsertSocketPlug(ctx, )


Insert a plug into a socketed item. I know how it sounds, but I assure you it's much more G-rated than you might be guessing. We haven't decided yet whether this will be able to insert plugs that have side effects, but if we do it will require special scope permission for an application attempting to do so. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. PREVIEW: This service is not yet active, but we are returning the planned schema of the endpoint for review, comment, and preparation for its eventual implementation.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2SearchDestinyEntities**
> InlineResponse20042 Destiny2SearchDestinyEntities(searchTerm, type_, optional)


Gets a page list of Destiny items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **searchTerm** | **string**| The string to use when searching for Destiny entities. | 
  **type_** | **string**| The type of entity for whom you would like results. These correspond to the entity&#39;s definition contract name. For instance, if you are looking for items, this property should be &#39;DestinyInventoryItemDefinition&#39;. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is tentatively in final form, but there may be bugs that prevent desirable operation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string**| The string to use when searching for Destiny entities. | 
 **type_** | **string**| The type of entity for whom you would like results. These correspond to the entity&#39;s definition contract name. For instance, if you are looking for items, this property should be &#39;DestinyInventoryItemDefinition&#39;. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is tentatively in final form, but there may be bugs that prevent desirable operation. | 
 **page** | **int32**| Page number to return, starting with 0. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

