# DestinyDefinitionsDestinyEquippingBlockDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GearsetItemHash** | **int32** | If the item is part of a gearset, this is a reference to that gearset item. | [optional] [default to null]
**UniqueLabel** | **string** | If defined, this is the label used to check if the item has other items of matching types already equipped.   For instance, when you aren&#39;t allowed to equip more than one Exotic Weapon, that&#39;s because all exotic weapons have identical uniqueLabels and the game checks the to-be-equipped item&#39;s uniqueLabel vs. all other already equipped items (other than the item in the slot that&#39;s about to be occupied). | [optional] [default to null]
**UniqueLabelHash** | **int32** | The hash of that unique label. Does not point to a specific definition. | [optional] [default to null]
**EquipmentSlotTypeHash** | **int32** | An equipped item *must* be equipped in an Equipment Slot. This is the hash identifier of the DestinyEquipmentSlotDefinition into which it must be equipped. | [optional] [default to null]
**Attributes** | [***interface{}**](interface{}.md) | These are custom attributes on the equippability of the item.  For now, this can only be \&quot;equip on acquire\&quot;, which would mean that the item will be automatically equipped as soon as you pick it up. | [optional] [default to null]
**DisplayStrings** | **[]string** | These are strings that represent the possible Game/Account/Character state failure conditions that can occur when trying to equip the item. They match up one-to-one with requiredUnlockExpressions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


