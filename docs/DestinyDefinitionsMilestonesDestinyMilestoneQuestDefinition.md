# DestinyDefinitionsMilestonesDestinyMilestoneQuestDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestItemHash** | **int32** | The item representing this Milestone quest. Use this hash to look up the DestinyInventoryItemDefinition for the quest to find its steps and human readable data. | [optional] [default to null]
**DisplayProperties** | [***interface{}**](interface{}.md) | The individual quests may have different definitions from the overall milestone: if there&#39;s a specific active quest, use these displayProperties instead of that of the overall DestinyMilestoneDefinition. | [optional] [default to null]
**OverrideImage** | **string** | If populated, this image can be shown instead of the generic milestone&#39;s image when this quest is live, or it can be used to show a background image for the quest itself that differs from that of the Activity or the Milestone. | [optional] [default to null]
**QuestRewards** | [***interface{}**](interface{}.md) | The rewards you will get for completing this quest, as best as we could extract them from our data. Sometimes, it&#39;ll be a decent amount of data. Sometimes, it&#39;s going to be sucky. Sorry. | [optional] [default to null]
**Activities** | [**map[string]DestinyDefinitionsMilestonesDestinyMilestoneActivityDefinition**](Destiny.Definitions.Milestones.DestinyMilestoneActivityDefinition.md) | The full set of all possible \&quot;conceptual activities\&quot; that are related to this Milestone. Tiers or alternative modes of play within these conceptual activities will be defined as sub-entities. Keyed by the Conceptual Activity Hash. Use the key to look up DestinyActivityDefinition. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


