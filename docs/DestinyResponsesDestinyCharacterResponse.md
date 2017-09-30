# DestinyResponsesDestinyCharacterResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | [***interface{}**](interface{}.md) | The character-level non-equipped inventory items.  COMPONENT TYPE: CharacterInventories | [optional] [default to null]
**Character** | [***interface{}**](interface{}.md) | Base information about the character in question.  COMPONENT TYPE: Characters | [optional] [default to null]
**Progressions** | [***interface{}**](interface{}.md) | Character progression data, including Milestones.  COMPONENT TYPE: CharacterProgressions | [optional] [default to null]
**RenderData** | [***interface{}**](interface{}.md) | Character rendering data - a minimal set of information about equipment and dyes used for rendering.  COMPONENT TYPE: CharacterRenderData | [optional] [default to null]
**Activities** | [***interface{}**](interface{}.md) | Activity data - info about current activities available to the player.  COMPONENT TYPE: CharacterActivities | [optional] [default to null]
**Equipment** | [***interface{}**](interface{}.md) | Equipped items on the character.  COMPONENT TYPE: CharacterEquipment | [optional] [default to null]
**Kiosks** | [***interface{}**](interface{}.md) | Items available from Kiosks that are available to this specific character.   COMPONENT TYPE: Kiosks | [optional] [default to null]
**ItemComponents** | [***interface{}**](interface{}.md) | The set of components belonging to the player&#39;s instanced items.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


