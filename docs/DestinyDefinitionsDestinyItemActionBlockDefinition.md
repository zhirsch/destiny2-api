# DestinyDefinitionsDestinyItemActionBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerbName** | **string** | Localized text for the verb of the action being performed. | [optional] [default to null]
**VerbDescription** | **string** | Localized text describing the action being performed. | [optional] [default to null]
**IsPositive** | **bool** | The content has this property, however it&#39;s not entirely clear how it is used. | [optional] [default to null]
**OverlayScreenName** | **string** | If the action has an overlay screen associated with it, this is the name of that screen. Unfortunately, we cannot return the screen&#39;s data itself. | [optional] [default to null]
**OverlayIcon** | **string** | The icon associated with the overlay screen for the action, if any. | [optional] [default to null]
**RequiredCooldownSeconds** | **int32** | The number of seconds to delay before allowing this action to be performed again. | [optional] [default to null]
**RequiredItems** | [**[]DestinyDefinitionsDestinyItemActionRequiredItemDefinition**](Destiny.Definitions.DestinyItemActionRequiredItemDefinition.md) | If the action requires other items to exist or be destroyed, this is the list of those items and requirements. | [optional] [default to null]
**ProgressionRewards** | [**[]DestinyDefinitionsDestinyProgressionRewardDefinition**](Destiny.Definitions.DestinyProgressionRewardDefinition.md) | If performing this action earns you Progression, this is the list of progressions and values granted for those progressions by performing this action. | [optional] [default to null]
**ActionTypeLabel** | **string** | The internal identifier for the action. | [optional] [default to null]
**RequiredLocation** | **string** | Theoretically, an item could have a localized string for a hint about the location in which the action should be performed. In practice, no items yet have this property. | [optional] [default to null]
**RequiredCooldownHash** | **int32** | The identifier hash for the Cooldown associated with this action. We have not pulled this data yet for you to have more data to use for cooldowns. | [optional] [default to null]
**DeleteOnAction** | **bool** | If true, the item is deleted when the action completes. | [optional] [default to null]
**ConsumeEntireStack** | **bool** | If true, the entire stack is deleted when the action completes. | [optional] [default to null]
**UseOnAcquire** | **bool** | If true, this action will be performed as soon as you earn this item. Some rewards work this way, providing you a single item to pick up from a reward-granting vendor in-game and then immediately consuming itself to provide you multiple items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


