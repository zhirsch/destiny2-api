# DestinyHistoricalStatsDestinyPlayer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinyUserInfo** | [***interface{}**](interface{}.md) | Details about the player as they are known in game (platform display name, Destiny emblem) | [optional] [default to null]
**CharacterClass** | **string** | Class of the character if applicable and available. | [optional] [default to null]
**CharacterLevel** | **int32** | Level of the character if available. Zero if it is not available. | [optional] [default to null]
**LightLevel** | **int32** | Light Level of the character if available. Zero if it is not available. | [optional] [default to null]
**BungieNetUserInfo** | [***interface{}**](interface{}.md) | Details about the player as they are known on BungieNet. This will be undefined if the player has marked their credential private, or does not have a BungieNet account. | [optional] [default to null]
**ClanName** | **string** | Current clan name for the player. This value may be null or an empty string if the user does not have a clan. | [optional] [default to null]
**ClanTag** | **string** | Current clan tag for the player. This value may be null or an empty string if the user does not have a clan. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


