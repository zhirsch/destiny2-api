# DestinyResponsesDestinyItemResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharacterId** | **int64** | If the item is on a character, this will return the ID of the character that is holding the item. | [optional] [default to null]
**Item** | [***interface{}**](interface{}.md) | Common data for the item relevant to its non-instanced properties.  COMPONENT TYPE: ItemCommonData | [optional] [default to null]
**Instance** | [***interface{}**](interface{}.md) | Basic instance data for the item.  COMPONENT TYPE: ItemInstances | [optional] [default to null]
**Objectives** | [***interface{}**](interface{}.md) | Information specifically about the item&#39;s objectives.  COMPONENT TYPE: ItemObjectives | [optional] [default to null]
**Perks** | [***interface{}**](interface{}.md) | Information specifically about the perks currently active on the item.  COMPONENT TYPE: ItemPerks | [optional] [default to null]
**RenderData** | [***interface{}**](interface{}.md) | Information about how to render the item in 3D.  COMPONENT TYPE: ItemRenderData | [optional] [default to null]
**Stats** | [***interface{}**](interface{}.md) | Information about the computed stats of the item: power, defense, etc...  COMPONENT TYPE: ItemStats | [optional] [default to null]
**TalentGrid** | [***interface{}**](interface{}.md) | Information about the talent grid attached to the item. Talent nodes can provide a variety of benefits and abilities, and in Destiny 2 are used almost exclusively for the character&#39;s \&quot;Builds\&quot;.  COMPONENT TYPE: ItemTalentGrids | [optional] [default to null]
**Sockets** | [***interface{}**](interface{}.md) | Information about the sockets of the item: which are currently active, what potential sockets you could have and the stats/abilities/perks you can gain from them.  COMPONENT TYPE: ItemSockets | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


