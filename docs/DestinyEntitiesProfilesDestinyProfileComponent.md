# DestinyEntitiesProfilesDestinyProfileComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfo** | [***interface{}**](interface{}.md) | If you need to render the Profile (their platform name, icon, etc...) somewhere, this property contains that information. | [optional] [default to null]
**DateLastPlayed** | [**time.Time**](time.Time.md) | The last time the user played with any character on this Profile. | [optional] [default to null]
**VersionsOwned** | [***interface{}**](interface{}.md) | If you want to know what expansions they own, this will contain that data. | [optional] [default to null]
**CharacterIds** | **[]int64** | A list of the character IDs, for further querying on your part. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


