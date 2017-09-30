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





type GroupsV2GroupResponse struct {
	
	Detail *GroupsV2GroupV2 `json:"detail,omitempty"`
	
	Founder *GroupsV2GroupMember `json:"founder,omitempty"`
	
	AlliedIds []int64 `json:"alliedIds,omitempty"`
	
	ParentGroup *GroupsV2GroupV2 `json:"parentGroup,omitempty"`
	
	AllianceStatus *GroupsV2GroupAllianceStatus `json:"allianceStatus,omitempty"`
	
	GroupJoinInviteCount int32 `json:"groupJoinInviteCount,omitempty"`
	// This property will be populated if the authenticated user is a member of the group. Note that because of account linking, a user can sometimes be part of a clan more than once. As such, this returns the highest member type available.
	CurrentUserMemberMap map[string]GroupsV2GroupMember `json:"currentUserMemberMap,omitempty"`
	// This property will be populated if the authenticated user is an applicant or has an outstanding invitation to join. Note that because of account linking, a user can sometimes be part of a clan more than once.
	CurrentUserPotentialMemberMap map[string]GroupsV2GroupPotentialMember `json:"currentUserPotentialMemberMap,omitempty"`
}

