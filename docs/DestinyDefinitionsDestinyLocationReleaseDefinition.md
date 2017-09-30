# DestinyDefinitionsDestinyLocationReleaseDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [***interface{}**](interface{}.md) | Sadly, these don&#39;t appear to be populated anymore (ever?) | [optional] [default to null]
**SpawnPoint** | **int32** | If we had map information, this spawnPoint would be interesting. But sadly, we don&#39;t have that info. | [optional] [default to null]
**DestinationHash** | **int32** | The Destination being pointed to by this location. | [optional] [default to null]
**ActivityHash** | **int32** | The Activity being pointed to by this location. | [optional] [default to null]
**ActivityGraphHash** | **int32** | The Activity Graph being pointed to by this location. | [optional] [default to null]
**ActivityGraphNodeHash** | **int32** | The Activity Graph Node being pointed to by this location. (Remember that Activity Graph Node hashes are only unique within an Activity Graph: so use the combination to find the node being spoken of) | [optional] [default to null]
**ActivityBubbleName** | **int32** | The Activity Bubble within the Destination. Look this up in the DestinyDestinationDefinition&#39;s bubbles and bubbleSettings properties. | [optional] [default to null]
**ActivityPathBundle** | **int32** | If we had map information, this would tell us something cool about the path this location wants you to take. I wish we had map information. | [optional] [default to null]
**ActivityPathDestination** | **int32** | If we had map information, this would tell us about path information related to destination on the map. Sad. Maybe you can do something cool with it. Go to town man. | [optional] [default to null]
**NavPointType** | [***interface{}**](interface{}.md) | The type of Nav Point that this represents. See the enumeration for more info. | [optional] [default to null]
**WorldPosition** | **[]int32** | Looks like it should be the position on the map, but sadly it does not look populated... yet? | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


