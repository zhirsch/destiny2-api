# DestinyResponsesDestinyProfileResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorReceipts** | [***interface{}**](interface{}.md) | Recent, refundable purchases you have made from vendors. When will you use it? Couldn&#39;t say...  COMPONENT TYPE: VendorReceipts | [optional] [default to null]
**ProfileInventory** | [***interface{}**](interface{}.md) | The profile-level inventory of the Destiny Profile.  COMPONENT TYPE: ProfileInventories | [optional] [default to null]
**ProfileCurrencies** | [***interface{}**](interface{}.md) | The profile-level currencies owned by the Destiny Profile.  COMPONENT TYPE: ProfileCurrencies | [optional] [default to null]
**Profile** | [***interface{}**](interface{}.md) | The basic information about the Destiny Profile (formerly \&quot;Account\&quot;).  COMPONENT TYPE: Profiles | [optional] [default to null]
**ProfileKiosks** | [***interface{}**](interface{}.md) | Items available from Kiosks that are available Profile-wide (i.e. across all characters)  This component returns information about what Kiosk items are available to you on a *Profile* level. It is theoretically possible for Kiosks to have items gated by specific Character as well. If you ever have those, you will find them on the characterKiosks property.  COMPONENT TYPE: Kiosks | [optional] [default to null]
**Characters** | [***interface{}**](interface{}.md) | Basic information about each character, keyed by the CharacterId.  COMPONENT TYPE: Characters | [optional] [default to null]
**CharacterInventories** | [***interface{}**](interface{}.md) | The character-level non-equipped inventory items, keyed by the Character&#39;s Id.  COMPONENT TYPE: CharacterInventories | [optional] [default to null]
**CharacterProgressions** | [***interface{}**](interface{}.md) | Character-level progression data, keyed by the Character&#39;s Id.  COMPONENT TYPE: CharacterProgressions | [optional] [default to null]
**CharacterRenderData** | [***interface{}**](interface{}.md) | Character rendering data - a minimal set of info needed to render a character in 3D - keyed by the Character&#39;s Id.  COMPONENT TYPE: CharacterRenderData | [optional] [default to null]
**CharacterActivities** | [***interface{}**](interface{}.md) | Character activity data - the activities available to this character and its status, keyed by the Character&#39;s Id.  COMPONENT TYPE: CharacterActivities | [optional] [default to null]
**CharacterEquipment** | [***interface{}**](interface{}.md) | The character&#39;s equipped items, keyed by the Character&#39;s Id.  COMPONENT TYPE: CharacterEquipment | [optional] [default to null]
**CharacterKiosks** | [***interface{}**](interface{}.md) | Items available from Kiosks that are available to a specific character as opposed to the account as a whole. It must be combined with data from the profileKiosks property to get a full picture of the character&#39;s available items to check out of a kiosk.  This component returns information about what Kiosk items are available to you on a *Character* level. Usually, kiosk items will be earned for the entire Profile (all characters) at once. To find those, look in the profileKiosks property.  COMPONENT TYPE: Kiosks | [optional] [default to null]
**ItemComponents** | [***interface{}**](interface{}.md) | Information about instanced items across all returned characters, keyed by the item&#39;s instance ID.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


