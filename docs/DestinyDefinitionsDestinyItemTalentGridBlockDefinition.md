# DestinyDefinitionsDestinyItemTalentGridBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TalentGridHash** | **int32** | The hash identifier of the DestinyTalentGridDefinition attached to this item. | [optional] [default to null]
**ItemDetailString** | **string** | This is meant to be a subtitle for looking at the talent grid. In practice, somewhat frustratingly, this always merely says the localized word for \&quot;Details\&quot;. Great. Maybe it&#39;ll have more if talent grids ever get used for more than builds and subclasses again. | [optional] [default to null]
**BuildName** | **string** | A shortcut string identifier for the \&quot;build\&quot; in question, if this talent grid has an associated build. Doesn&#39;t map to anything we can expose at the moment. | [optional] [default to null]
**HudDamageType** | [***interface{}**](interface{}.md) | If the talent grid implies a damage type, this is the enum value for that damage type. | [optional] [default to null]
**HudIcon** | **string** | If the talent grid has a special icon that&#39;s shown in the game UI (like builds, funny that), this is the identifier for that icon. Sadly, we don&#39;t actually get that icon right now. I&#39;ll be looking to replace this with a path to the actual icon itself. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


