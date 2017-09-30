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





type DestinyItemComponentSetOfint64 struct {
	
	Instances *DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent `json:"instances,omitempty"`
	
	Objectives *DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent `json:"objectives,omitempty"`
	
	Perks *DictionaryComponentResponseOfint64AndDestinyItemPerksComponent `json:"perks,omitempty"`
	
	RenderData *DictionaryComponentResponseOfint64AndDestinyItemRenderComponent `json:"renderData,omitempty"`
	
	Stats *DictionaryComponentResponseOfint64AndDestinyItemStatsComponent `json:"stats,omitempty"`
	
	Sockets *DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent `json:"sockets,omitempty"`
	
	TalentGrids *DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent `json:"talentGrids,omitempty"`
	
	PlugStates *DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent `json:"plugStates,omitempty"`
}

