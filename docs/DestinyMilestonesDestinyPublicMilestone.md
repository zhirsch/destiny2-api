# DestinyMilestonesDestinyPublicMilestone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MilestoneHash** | **int32** | The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone. | [optional] [default to null]
**AvailableQuests** | [**[]DestinyMilestonesDestinyPublicMilestoneQuest**](Destiny.Milestones.DestinyPublicMilestoneQuest.md) | A milestone not need have even a single quest, but if there are active quests they will be returned here. | [optional] [default to null]
**VendorHashes** | **[]int32** | Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant. | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone started/became active. | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone will expire/recycle/end. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


