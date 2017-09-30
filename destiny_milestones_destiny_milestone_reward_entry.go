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




// The character-specific data for a milestone's reward entry. See DestinyMilestoneDefinition for more information about Reward Entries.
type DestinyMilestonesDestinyMilestoneRewardEntry struct {
	// The identifier for the reward entry in question. It is important to look up the related DestinyMilestoneRewardEntryDefinition to get the static details about the reward, which you can do by looking up the milestone's DestinyMilestoneDefinition and examining the DestinyMilestoneDefinition.rewards[rewardCategoryHash].rewardEntries[rewardEntryHash] data.
	RewardEntryHash int32 `json:"rewardEntryHash,omitempty"`
	// If TRUE, the player has earned this reward.
	Earned bool `json:"earned,omitempty"`
	// If TRUE, the player has redeemed/picked up/obtained this reward. Feel free to alias this to \"gotTheShinyBauble\" in your own codebase.
	Redeemed bool `json:"redeemed,omitempty"`
}
