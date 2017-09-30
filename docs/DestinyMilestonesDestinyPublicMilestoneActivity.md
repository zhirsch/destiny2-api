# DestinyMilestonesDestinyPublicMilestoneActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash identifier of the activity that&#39;s been chosen to be considered the canonical \&quot;conceptual\&quot; activity definition. This may have many variants, defined herein. | [optional] [default to null]
**ModifierHashes** | **[]int32** | The activity may have 0-to-many modifiers: if it does, this will contain the hashes to the DestinyActivityModifierDefinition that defines the modifier being applied. | [optional] [default to null]
**Variants** | [**[]DestinyMilestonesDestinyPublicMilestoneActivityVariant**](Destiny.Milestones.DestinyPublicMilestoneActivityVariant.md) | Every relevant variation of this conceptual activity, including the conceptual activity itself, have variants defined here. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


