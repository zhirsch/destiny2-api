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




// A minimal view of a character's equipped items, for the purpose of rendering a summary screen or showing the character in 3D.
type DestinyCharacterDestinyCharacterPeerView struct {
	
	Equipment []DestinyCharacterDestinyItemPeerView `json:"equipment,omitempty"`
}

