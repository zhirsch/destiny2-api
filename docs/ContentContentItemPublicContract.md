# ContentContentItemPublicContract

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **int64** |  | [optional] [default to null]
**CType** | **string** |  | [optional] [default to null]
**CmsPath** | **string** |  | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ModifyDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**AllowComments** | **bool** |  | [optional] [default to null]
**HasAgeGate** | **bool** |  | [optional] [default to null]
**MinimumAge** | **int32** |  | [optional] [default to null]
**RatingImagePath** | **string** |  | [optional] [default to null]
**Author** | [***UserGeneralUser**](User.GeneralUser.md) |  | [optional] [default to null]
**AutoEnglishPropertyFallback** | **bool** |  | [optional] [default to null]
**Properties** | [**map[string]interface{}**](interface{}.md) | Firehose content is really a collection of metadata and \&quot;properties\&quot;, which are the potentially-but-not-strictly localizable data that comprises the meat of whatever content is being shown.  As Cole Porter would have crooned, \&quot;Anything Goes\&quot; with Firehose properties. They are most often strings, but they can theoretically be anything. They are JSON encoded, and could be JSON structures, simple strings, numbers etc... The Content Type of the item (cType) will describe the properties, and thus how they ought to be deserialized. | [optional] [default to null]
**Representations** | [**[]ContentContentRepresentation**](Content.ContentRepresentation.md) |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**CommentSummary** | [***ContentCommentSummary**](Content.CommentSummary.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


