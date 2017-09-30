# DestinyEntitiesCharactersDestinyCharacterComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembershipId** | **int64** | Every Destiny Profile has a membershipId. This is provided on the character as well for convenience. | [optional] [default to null]
**MembershipType** | [***interface{}**](interface{}.md) | membershipType tells you the platform on which the character plays. Examine the BungieMembershipType enumeration for possible values. | [optional] [default to null]
**CharacterId** | **int64** | The unique identifier for the character. | [optional] [default to null]
**DateLastPlayed** | [**time.Time**](time.Time.md) | The last date that the user played Destiny. | [optional] [default to null]
**MinutesPlayedThisSession** | **int64** | If the user is currently playing, this is how long they&#39;ve been playing. | [optional] [default to null]
**MinutesPlayedTotal** | **int64** | If this value is 525,600, then they played Destiny for a year. Or they&#39;re a very dedicated Rent fan. Note that this includes idle time, not just time spent actually in activities shooting things. | [optional] [default to null]
**Light** | **int32** | The user&#39;s calculated \&quot;Light Level\&quot;. Light level is an indicator of your power that mostly matters in the end game, once you&#39;ve reached the maximum character level: it&#39;s a level that&#39;s dependent on the average Attack/Defense power of your items. | [optional] [default to null]
**Stats** | **map[string]int32** | Your character&#39;s stats, such as Agility, Resilience, etc... *not* historical stats.  You&#39;ll have to call a different endpoint for those. | [optional] [default to null]
**RaceHash** | **int32** | Use this hash to look up the character&#39;s DestinyRaceDefinition. | [optional] [default to null]
**GenderHash** | **int32** | Use this hash to look up the character&#39;s DestinyGenderDefinition. | [optional] [default to null]
**ClassHash** | **int32** | Use this hash to look up the character&#39;s DestinyClassDefinition. | [optional] [default to null]
**RaceType** | [***interface{}**](interface{}.md) | Mostly for historical purposes at this point, this is an enumeration for the character&#39;s race.  It&#39;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. | [optional] [default to null]
**ClassType** | [***interface{}**](interface{}.md) | Mostly for historical purposes at this point, this is an enumeration for the character&#39;s class.  It&#39;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. | [optional] [default to null]
**GenderType** | [***interface{}**](interface{}.md) | Mostly for historical purposes at this point, this is an enumeration for the character&#39;s Gender.  It&#39;ll be preferable in the general case to look up the related definition: but for some people this was too convenient to remove. And yeah, it&#39;s an enumeration and not a boolean. Fight me. | [optional] [default to null]
**EmblemPath** | **string** | A shortcut path to the user&#39;s currently equipped emblem image. If you&#39;re just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition. | [optional] [default to null]
**EmblemBackgroundPath** | **string** | A shortcut path to the user&#39;s currently equipped emblem background image. If you&#39;re just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition. | [optional] [default to null]
**EmblemHash** | **int32** | The hash of the currently equipped emblem for the user. Can be used to look up the DestinyInventoryItemDefinition. | [optional] [default to null]
**LevelProgression** | [***interface{}**](interface{}.md) | The progression that indicates your character&#39;s level. Not their light level, but their character level: you know, the thing you max out a couple hours in and then ignore for the sake of light level. | [optional] [default to null]
**BaseCharacterLevel** | **int32** | The \&quot;base\&quot; level of your character, not accounting for any light level. | [optional] [default to null]
**PercentToNextLevel** | **float32** | A number between 0 and 100, indicating the whole and fractional % remaining to get to the next character level. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


