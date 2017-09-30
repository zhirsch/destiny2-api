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




// This appears to be information used when rendering rewards. We don't currently use it on BNet.
type DestinyDefinitionsDestinyItemSummaryBlockDefinition struct {
	// Apparently when rendering an item in a reward, this should be used as a sort priority. We're not doing it presently.
	SortPriority int32 `json:"sortPriority,omitempty"`
}

