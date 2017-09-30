/* 
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@bungie.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package api





type DestinyRequestsDestinyItemTransferRequest struct {
	
	ItemReferenceHash int32 `json:"itemReferenceHash,omitempty"`
	
	StackSize int32 `json:"stackSize,omitempty"`
	
	TransferToVault bool `json:"transferToVault,omitempty"`
	
	ItemId int64 `json:"itemId,omitempty"`
	
	CharacterId int64 `json:"characterId,omitempty"`
	
	MembershipType *BungieMembershipType `json:"membershipType,omitempty"`
}

