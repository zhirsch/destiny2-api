# DestinyDefinitionsDestinySandboxPerkDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***interface{}**](interface{}.md) | These display properties are by no means guaranteed to be populated. Usually when it is, it&#39;s only because we back-filled them with the displayProperties of some Talent Node or Plug item that happened to be uniquely providing that perk. | [optional] [default to null]
**PerkIdentifier** | **string** | The string identifier for the perk. | [optional] [default to null]
**IsDisplayable** | **bool** | If true, you can actually show the perk in the UI. Otherwise, it doesn&#39;t have useful player-facing information. | [optional] [default to null]
**DamageType** | [***interface{}**](interface{}.md) | If this perk grants a damage type to a weapon, the damage type will be defined here.  Unless you have a compelling reason to use this enum value, use the damageTypeHash instead to look up the actual DestinyDamageTypeDefinition. | [optional] [default to null]
**DamageTypeHash** | **int32** | The hash identifier for looking up the DestinyDamageTypeDefinition, if this perk has a damage type.  This is preferred over using the damageType enumeration value, which has been left purely because it is occasionally convenient. | [optional] [default to null]
**PerkGroups** | [***interface{}**](interface{}.md) | An old holdover from the original Armory, this was an attempt to group perks by functionality.  It is as yet unpopulated, and there will be quite a bit of work needed to restore it to its former working order. | [optional] [default to null]
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] [default to null]
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] [default to null]
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


