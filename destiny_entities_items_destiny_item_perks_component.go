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




// Instanced items can have perks: benefits that the item bestows.  These are related to DestinySandboxPerkDefinition, and sometimes - but not always - have human readable info. When they do, they are the icons and text that you see in an item's tooltip.  Talent Grids, Sockets, and the item itself can apply Perks, which are then summarized here for your convenience.
type DestinyEntitiesItemsDestinyItemPerksComponent struct {
	// The list of perks to display in an item tooltip - and whether or not they have been activated.
	Perks []DestinyPerksDestinyPerkReference `json:"perks,omitempty"`
}

