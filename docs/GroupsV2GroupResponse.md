# GroupsV2GroupResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [***GroupsV2GroupV2**](GroupsV2.GroupV2.md) |  | [optional] [default to null]
**Founder** | [***GroupsV2GroupMember**](GroupsV2.GroupMember.md) |  | [optional] [default to null]
**AlliedIds** | **[]int64** |  | [optional] [default to null]
**ParentGroup** | [***GroupsV2GroupV2**](GroupsV2.GroupV2.md) |  | [optional] [default to null]
**AllianceStatus** | [***GroupsV2GroupAllianceStatus**](GroupsV2.GroupAllianceStatus.md) |  | [optional] [default to null]
**GroupJoinInviteCount** | **int32** |  | [optional] [default to null]
**CurrentUserMemberMap** | [**map[string]GroupsV2GroupMember**](GroupsV2.GroupMember.md) | This property will be populated if the authenticated user is a member of the group. Note that because of account linking, a user can sometimes be part of a clan more than once. As such, this returns the highest member type available. | [optional] [default to null]
**CurrentUserPotentialMemberMap** | [**map[string]GroupsV2GroupPotentialMember**](GroupsV2.GroupPotentialMember.md) | This property will be populated if the authenticated user is an applicant or has an outstanding invitation to join. Note that because of account linking, a user can sometimes be part of a clan more than once. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


