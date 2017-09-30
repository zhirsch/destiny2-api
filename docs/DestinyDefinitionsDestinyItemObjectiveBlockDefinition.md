# DestinyDefinitionsDestinyItemObjectiveBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectiveHashes** | **[]int32** | The hashes to Objectives (DestinyObjectiveDefinition) that are part of this Quest Step, in the order that they should be rendered. | [optional] [default to null]
**DisplayActivityHashes** | **[]int32** | For every entry in objectiveHashes, there is a corresponding entry in this array at the same index. If the objective is meant to be associated with a specific DestinyActivityDefinition, there will be a valid hash at that index. Otherwise, it will be invalid (0). | [optional] [default to null]
**RequireFullObjectiveCompletion** | **bool** | If True, all objectives must be completed for the step to be completed. If False, any one objective can be completed for the step to be completed. | [optional] [default to null]
**QuestlineItemHash** | **int32** | The hash for the DestinyInventoryItemDefinition representing the Quest to which this Quest Step belongs. | [optional] [default to null]
**Narrative** | **string** | The localized string for narrative text related to this quest step, if any. | [optional] [default to null]
**ObjectiveVerbName** | **string** | The localized string describing an action to be performed associated with the objectives, if any. | [optional] [default to null]
**QuestTypeIdentifier** | **string** | The identifier for the type of quest being performed, if any. Not associated with any fixed definition, yet. | [optional] [default to null]
**QuestTypeHash** | **int32** | A hashed value for the questTypeIdentifier, because apparently I like to be redundant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


