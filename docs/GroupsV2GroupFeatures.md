# GroupsV2GroupFeatures

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumMembers** | **int32** |  | [optional] [default to null]
**MaximumMembershipsOfGroupType** | **int32** | Maximum number of groups of this type a typical membership may join. For example, a user may join about 50 General groups with their Bungie.net account. They may join one clan per Destiny membership. | [optional] [default to null]
**Capabilities** | [***GroupsV2Capabilities**](GroupsV2.Capabilities.md) |  | [optional] [default to null]
**MembershipTypes** | [**[]BungieMembershipType**](BungieMembershipType.md) |  | [optional] [default to null]
**InvitePermissionOverride** | **bool** | Minimum Member Level allowed to invite new members to group  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#39;t  Default is false for clans, true for groups. | [optional] [default to null]
**UpdateCulturePermissionOverride** | **bool** | Minimum Member Level allowed to update group culture  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#39;t  Default is false for clans, true for groups. | [optional] [default to null]
**HostGuidedGamePermissionOverride** | [***interface{}**](interface{}.md) | Minimum Member Level allowed to host guided games  Always Allowed: Founder, Acting Founder, Admin  Allowed Overrides: None, Member, Beginner  Default is Member for clans, None for groups, although this means nothing for groups. | [optional] [default to null]
**UpdateBannerPermissionOverride** | **bool** | Minimum Member Level allowed to update banner  Always Allowed: Founder, Acting Founder  True means admins have this power, false means they don&#39;t  Default is false for clans, true for groups. | [optional] [default to null]
**JoinLevel** | [***interface{}**](interface{}.md) | Level to join a member at when accepting an invite, application, or joining an open clan  Default is Beginner. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


