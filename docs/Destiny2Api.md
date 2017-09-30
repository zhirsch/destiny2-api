# \Destiny2Api

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destiny2ActivateTalentNode**](Destiny2Api.md#Destiny2ActivateTalentNode) | **Post** /Destiny2/Actions/Items/ActivateTalentNode/ | 
[**Destiny2EquipItem**](Destiny2Api.md#Destiny2EquipItem) | **Post** /Destiny2/Actions/Items/EquipItem/ | 
[**Destiny2EquipItems**](Destiny2Api.md#Destiny2EquipItems) | **Post** /Destiny2/Actions/Items/EquipItems/ | 
[**Destiny2GetActivityHistory**](Destiny2Api.md#Destiny2GetActivityHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/Activities/ | 
[**Destiny2GetCharacter**](Destiny2Api.md#Destiny2GetCharacter) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/ | 
[**Destiny2GetClanAggregateStats**](Destiny2Api.md#Destiny2GetClanAggregateStats) | **Get** /Destiny2/Stats/AggregateClanStats/{groupId}/ | 
[**Destiny2GetClanLeaderboards**](Destiny2Api.md#Destiny2GetClanLeaderboards) | **Get** /Destiny2/Stats/Leaderboards/Clans/{groupId}/ | 
[**Destiny2GetClanWeeklyRewardState**](Destiny2Api.md#Destiny2GetClanWeeklyRewardState) | **Get** /Destiny2/Clan/{groupId}/WeeklyRewardState/ | 
[**Destiny2GetDestinyAggregateActivityStats**](Destiny2Api.md#Destiny2GetDestinyAggregateActivityStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/AggregateActivityStats/ | 
[**Destiny2GetDestinyEntityDefinition**](Destiny2Api.md#Destiny2GetDestinyEntityDefinition) | **Get** /Destiny2/Manifest/{entityType}/{hashIdentifier}/ | 
[**Destiny2GetDestinyManifest**](Destiny2Api.md#Destiny2GetDestinyManifest) | **Get** /Destiny2/Manifest/ | 
[**Destiny2GetHistoricalStats**](Destiny2Api.md#Destiny2GetHistoricalStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/ | 
[**Destiny2GetHistoricalStatsDefinition**](Destiny2Api.md#Destiny2GetHistoricalStatsDefinition) | **Get** /Destiny2/Stats/Definition/ | 
[**Destiny2GetHistoricalStatsForAccount**](Destiny2Api.md#Destiny2GetHistoricalStatsForAccount) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/ | 
[**Destiny2GetItem**](Destiny2Api.md#Destiny2GetItem) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Item/{itemInstanceId}/ | 
[**Destiny2GetLeaderboards**](Destiny2Api.md#Destiny2GetLeaderboards) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/ | 
[**Destiny2GetLeaderboardsForCharacter**](Destiny2Api.md#Destiny2GetLeaderboardsForCharacter) | **Get** /Destiny2/Stats/Leaderboards/{membershipType}/{destinyMembershipId}/{characterId}/ | 
[**Destiny2GetPostGameCarnageReport**](Destiny2Api.md#Destiny2GetPostGameCarnageReport) | **Get** /Destiny2/Stats/PostGameCarnageReport/{activityId}/ | 
[**Destiny2GetProfile**](Destiny2Api.md#Destiny2GetProfile) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/ | 
[**Destiny2GetPublicMilestoneContent**](Destiny2Api.md#Destiny2GetPublicMilestoneContent) | **Get** /Destiny2/Milestones/{milestoneHash}/Content/ | 
[**Destiny2GetPublicMilestones**](Destiny2Api.md#Destiny2GetPublicMilestones) | **Get** /Destiny2/Milestones/ | 
[**Destiny2GetUniqueWeaponHistory**](Destiny2Api.md#Destiny2GetUniqueWeaponHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/UniqueWeapons/ | 
[**Destiny2GetVendor**](Destiny2Api.md#Destiny2GetVendor) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/{vendorHash}/ | 
[**Destiny2GetVendors**](Destiny2Api.md#Destiny2GetVendors) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/ | 
[**Destiny2InsertSocketPlug**](Destiny2Api.md#Destiny2InsertSocketPlug) | **Post** /Destiny2/Actions/Items/InsertSocketPlug/ | 
[**Destiny2SearchDestinyEntities**](Destiny2Api.md#Destiny2SearchDestinyEntities) | **Get** /Destiny2/Armory/Search/{type}/{searchTerm}/ | 
[**Destiny2SearchDestinyPlayer**](Destiny2Api.md#Destiny2SearchDestinyPlayer) | **Get** /Destiny2/SearchDestinyPlayer/{membershipType}/{displayName}/ | 
[**Destiny2SetItemLockState**](Destiny2Api.md#Destiny2SetItemLockState) | **Post** /Destiny2/Actions/Items/SetLockState/ | 
[**Destiny2TransferItem**](Destiny2Api.md#Destiny2TransferItem) | **Post** /Destiny2/Actions/Items/TransferItem/ | 


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

# **Destiny2EquipItem**
> InlineResponse20015 Destiny2EquipItem(ctx, )


Equip an item. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline.

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

# **Destiny2EquipItems**
> InlineResponse20037 Destiny2EquipItems(ctx, )


Equip a list of items by itemInstanceIds. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. Any items not found on your character will be ignored.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

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

# **Destiny2GetCharacter**
> InlineResponse20032 Destiny2GetCharacter(characterId, destinyMembershipId, membershipType, optional)


Returns character information for the supplied character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **characterId** | **int64**| ID of the character. | 
  **destinyMembershipId** | **int64**| Destiny membership ID. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **characterId** | **int64**| ID of the character. | 
 **destinyMembershipId** | **int64**| Destiny membership ID. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **components** | [**[]DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

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

# **Destiny2GetClanWeeklyRewardState**
> InlineResponse20033 Destiny2GetClanWeeklyRewardState(groupId)


Returns information on the weekly clan rewards and if the clan has earned them or not. Note that this will always report rewards as not redeemed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **groupId** | **int64**| A valid group id of clan. | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

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

# **Destiny2GetDestinyEntityDefinition**
> InlineResponse20029 Destiny2GetDestinyEntityDefinition(entityType, hashIdentifier)


Returns the static definition of an entity of the given Type and hash identifier. Examine the API Documentation for the Type Names of entities that have their own definitions. Note that the return type will always *inherit from* DestinyDefinition, but the specific type returned will be the requested entity type if it can be found. Please don't use this as a chatty alternative to the Manifest database if you require large sets of data, but for simple and one-off accesses this should be handy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **entityType** | **string**| The type of entity for whom you would like results. These correspond to the entity&#39;s definition contract name. For instance, if you are looking for items, this property should be &#39;DestinyInventoryItemDefinition&#39;. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is tentatively in final form, but there may be bugs that prevent desirable operation. | 
  **hashIdentifier** | **int32**| The hash identifier for the specific Entity you want returned. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetDestinyManifest**
> InlineResponse20028 Destiny2GetDestinyManifest()


Returns the current version of the manifest as a json object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

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

# **Destiny2GetHistoricalStatsDefinition**
> InlineResponse20039 Destiny2GetHistoricalStatsDefinition()


Gets historical stats definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

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

# **Destiny2GetItem**
> InlineResponse20034 Destiny2GetItem(destinyMembershipId, itemInstanceId, membershipType, optional)


Retrieve the details of an instanced Destiny Item. An instanced Destiny item is one with an ItemInstanceId. Non-instanced items, such as materials, have no useful instance-specific details and thus are not queryable here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **destinyMembershipId** | **int64**| The membership ID of the destiny profile. | 
  **itemInstanceId** | **int64**| The Instance ID of the destiny item. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinyMembershipId** | **int64**| The membership ID of the destiny profile. | 
 **itemInstanceId** | **int64**| The Instance ID of the destiny item. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **components** | [**[]DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

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

# **Destiny2GetPostGameCarnageReport**
> InlineResponse20038 Destiny2GetPostGameCarnageReport(activityId)


Gets the available post game carnage report for the activity ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **activityId** | **int64**| The ID of the activity whose PGCR is requested. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetProfile**
> InlineResponse20031 Destiny2GetProfile(destinyMembershipId, membershipType, optional)


Returns Destiny Profile information for the supplied membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **destinyMembershipId** | **int64**| Destiny membership ID. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinyMembershipId** | **int64**| Destiny membership ID. | 
 **membershipType** | **int32**| A valid non-BungieNet membership type. | 
 **components** | [**[]DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPublicMilestoneContent**
> InlineResponse20048 Destiny2GetPublicMilestoneContent(milestoneHash)


Gets custom localized content for the milestone of the given hash, if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **milestoneHash** | **int32**| The identifier for the milestone to be returned. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPublicMilestones**
> InlineResponse20049 Destiny2GetPublicMilestones()


Gets public information about currently available Milestones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

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

# **Destiny2SearchDestinyPlayer**
> InlineResponse20030 Destiny2SearchDestinyPlayer(displayName, membershipType)


Returns a list of Destiny memberships given a full Gamertag or PSN ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **displayName** | **string**| The full gamertag or PSN id of the player. Spaces and case are ignored. | 
  **membershipType** | **int32**| A valid non-BungieNet membership type, or All. | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2SetItemLockState**
> InlineResponse20015 Destiny2SetItemLockState(ctx, )


Set the Lock State for an instanced item. You must have a valid Destiny Account.

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

# **Destiny2TransferItem**
> InlineResponse20015 Destiny2TransferItem(ctx, )


Transfer an item to/from your vault. You must have a valid Destiny account. You must also pass BOTH a reference AND an instance ID if it's an instanced item. itshappening.gif

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

