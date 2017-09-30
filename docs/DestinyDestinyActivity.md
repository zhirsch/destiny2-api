# DestinyDestinyActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash identifier of the Activity. Use this to look up the DestinyActivityDefinition of the activity. | [optional] [default to null]
**IsNew** | **bool** | If true, then the activity should have a \&quot;new\&quot; indicator in the Director UI. | [optional] [default to null]
**CanLead** | **bool** | If true, the user is allowed to lead a Fireteam into this activity. | [optional] [default to null]
**CanJoin** | **bool** | If true, the user is allowed to join with another Fireteam in this activity. | [optional] [default to null]
**IsCompleted** | **bool** | If true, we both have the ability to know that the user has completed this activity and they have completed it. Unfortunately, we can&#39;t necessarily know this for all activities. As such, this should probably only be used if you already know in advance which specific activities you wish to check. | [optional] [default to null]
**IsVisible** | **bool** | If true, the user should be able to see this activity. | [optional] [default to null]
**DisplayLevel** | **int32** | The difficulty level of the activity, if applicable. | [optional] [default to null]
**RecommendedLight** | **int32** | The recommended light level for the activity, if applicable. | [optional] [default to null]
**DifficultyTier** | [***interface{}**](interface{}.md) | A DestinyActivityDifficultyTier enum value indicating the difficulty of the activity. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


