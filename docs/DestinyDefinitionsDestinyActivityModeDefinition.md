# DestinyDefinitionsDestinyActivityModeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] [default to null]
**PgcrImage** | **string** |  | [optional] [default to null]
**ModeType** | [***DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](Destiny.HistoricalStats.Definitions.DestinyActivityModeType.md) |  | [optional] [default to null]
**ActivityModeCategory** | [***DestinyDestinyActivityModeCategory**](Destiny.DestinyActivityModeCategory.md) |  | [optional] [default to null]
**IsTeamBased** | **bool** |  | [optional] [default to null]
**IsAggregateMode** | **bool** |  | [optional] [default to null]
**ParentHashes** | **[]int32** |  | [optional] [default to null]
**FriendlyName** | **string** |  | [optional] [default to null]
**ActivityModeMappings** | [**map[string]DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](Destiny.HistoricalStats.Definitions.DestinyActivityModeType.md) |  | [optional] [default to null]
**Display** | **bool** | If FALSE, we want to ignore this type when we&#39;re showing activity modes in BNet UI. It will still be returned in case 3rd parties want to use it for any purpose. | [optional] [default to null]
**Order** | **int32** | The relative ordering of activity modes. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


