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




// An intrinsic perk on an item, and the requirements for it to be activated.
type DestinyDefinitionsDestinyItemPerkEntryDefinition struct {
	// If this perk is not active, this is the string to show for why it's not providing its benefits.
	RequirementDisplayString string `json:"requirementDisplayString,omitempty"`
	// A hash identifier for the DestinySandboxPerkDefinition being provided on the item.
	PerkHash uint32 `json:"perkHash,omitempty"`
}

