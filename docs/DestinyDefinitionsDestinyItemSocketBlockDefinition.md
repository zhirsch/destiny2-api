# DestinyDefinitionsDestinyItemSocketBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** | This was supposed to be a string that would give per-item details about sockets. In practice, it turns out that all this ever has is the localized word \&quot;details\&quot;. ... that&#39;s lame, but perhaps it will become something cool in the future. | [optional] [default to null]
**SocketEntries** | [**[]DestinyDefinitionsDestinyItemSocketEntryDefinition**](Destiny.Definitions.DestinyItemSocketEntryDefinition.md) | Each socket on an item is defined here. Check inside for more info. | [optional] [default to null]
**SocketCategories** | [**[]DestinyDefinitionsDestinyItemSocketCategoryDefinition**](Destiny.Definitions.DestinyItemSocketCategoryDefinition.md) | A convenience property, that refers to the sockets in the \&quot;sockets\&quot; property, pre-grouped by category and ordered in the manner that they should be grouped in the UI. You could form this yourself with the existing data, but why would you want to? Enjoy life man. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


