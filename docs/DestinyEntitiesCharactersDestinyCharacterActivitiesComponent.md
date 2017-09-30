# DestinyEntitiesCharactersDestinyCharacterActivitiesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateActivityStarted** | [**time.Time**](time.Time.md) | The last date that the user started playing an activity. | [optional] [default to null]
**AvailableActivities** | [**[]DestinyDestinyActivity**](Destiny.DestinyActivity.md) | The list of activities that the user can play. | [optional] [default to null]
**CurrentActivityHash** | **int32** | If the user is in an activity, this will be the hash of the Activity being played. Note that you must combine this info with currentActivityModeHash to get a real picture of what the user is doing right now. For instance, PVP \&quot;Activities\&quot; are just maps: it&#39;s the ActivityMode that determines what type of PVP game they&#39;re playing. | [optional] [default to null]
**CurrentActivityModeHash** | **int32** | If the user is in an activity, this will be the hash of the activity mode being played. Combine with currentActivityHash to give a person a full picture of what they&#39;re doing right now. | [optional] [default to null]
**CurrentActivityModeType** | **int32** |  | [optional] [default to null]
**CurrentActivityModeHashes** | **[]int32** |  | [optional] [default to null]
**CurrentActivityModeTypes** | [**[]DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](Destiny.HistoricalStats.Definitions.DestinyActivityModeType.md) |  | [optional] [default to null]
**CurrentPlaylistActivityHash** | **int32** |  | [optional] [default to null]
**LastCompletedStoryHash** | **int32** | This will have the activity hash of the last completed story/campaign mission, in case you care about that. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


