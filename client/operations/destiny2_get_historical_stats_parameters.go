// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDestiny2GetHistoricalStatsParams creates a new Destiny2GetHistoricalStatsParams object
// with the default values initialized.
func NewDestiny2GetHistoricalStatsParams() *Destiny2GetHistoricalStatsParams {
	var ()
	return &Destiny2GetHistoricalStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetHistoricalStatsParamsWithTimeout creates a new Destiny2GetHistoricalStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetHistoricalStatsParamsWithTimeout(timeout time.Duration) *Destiny2GetHistoricalStatsParams {
	var ()
	return &Destiny2GetHistoricalStatsParams{

		timeout: timeout,
	}
}

// NewDestiny2GetHistoricalStatsParamsWithContext creates a new Destiny2GetHistoricalStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetHistoricalStatsParamsWithContext(ctx context.Context) *Destiny2GetHistoricalStatsParams {
	var ()
	return &Destiny2GetHistoricalStatsParams{

		Context: ctx,
	}
}

// NewDestiny2GetHistoricalStatsParamsWithHTTPClient creates a new Destiny2GetHistoricalStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetHistoricalStatsParamsWithHTTPClient(client *http.Client) *Destiny2GetHistoricalStatsParams {
	var ()
	return &Destiny2GetHistoricalStatsParams{
		HTTPClient: client,
	}
}

/*Destiny2GetHistoricalStatsParams contains all the parameters to send to the API endpoint
for the destiny2 get historical stats operation typically these are written to a http.Request
*/
type Destiny2GetHistoricalStatsParams struct {

	/*CharacterID
	  The id of the character to retrieve. You can omit this character ID or set it to 0 to get aggregate stats across all characters.

	*/
	CharacterID int64
	/*Dayend
	  Last day to return when daily stats are requested. Use the format YYYY-MM-DD.

	*/
	Dayend *strfmt.DateTime
	/*Daystart
	  First day to return when daily stats are requested. Use the format YYYY-MM-DD

	*/
	Daystart *strfmt.DateTime
	/*DestinyMembershipID
	  The Destiny membershipId of the user to retrieve.

	*/
	DestinyMembershipID int64
	/*Groups
	  Group of stats to include, otherwise only general stats are returned. Comma separated list is allowed. Values: General, Weapons, Medals

	*/
	Groups []int64
	/*MembershipType
	  A valid non-BungieNet membership type.

	*/
	MembershipType int32
	/*Modes
	  Game modes to return. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited.

	*/
	Modes []int64
	/*PeriodType
	  Indicates a specific period type to return. Optional. May be: Daily, AllTime, or Activity

	*/
	PeriodType *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithTimeout(timeout time.Duration) *Destiny2GetHistoricalStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithContext(ctx context.Context) *Destiny2GetHistoricalStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithHTTPClient(client *http.Client) *Destiny2GetHistoricalStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithCharacterID(characterID int64) *Destiny2GetHistoricalStatsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetCharacterID(characterID int64) {
	o.CharacterID = characterID
}

// WithDayend adds the dayend to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithDayend(dayend *strfmt.DateTime) *Destiny2GetHistoricalStatsParams {
	o.SetDayend(dayend)
	return o
}

// SetDayend adds the dayend to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetDayend(dayend *strfmt.DateTime) {
	o.Dayend = dayend
}

// WithDaystart adds the daystart to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithDaystart(daystart *strfmt.DateTime) *Destiny2GetHistoricalStatsParams {
	o.SetDaystart(daystart)
	return o
}

// SetDaystart adds the daystart to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetDaystart(daystart *strfmt.DateTime) {
	o.Daystart = daystart
}

// WithDestinyMembershipID adds the destinyMembershipID to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithDestinyMembershipID(destinyMembershipID int64) *Destiny2GetHistoricalStatsParams {
	o.SetDestinyMembershipID(destinyMembershipID)
	return o
}

// SetDestinyMembershipID adds the destinyMembershipId to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetDestinyMembershipID(destinyMembershipID int64) {
	o.DestinyMembershipID = destinyMembershipID
}

// WithGroups adds the groups to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithGroups(groups []int64) *Destiny2GetHistoricalStatsParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetGroups(groups []int64) {
	o.Groups = groups
}

// WithMembershipType adds the membershipType to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithMembershipType(membershipType int32) *Destiny2GetHistoricalStatsParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WithModes adds the modes to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithModes(modes []int64) *Destiny2GetHistoricalStatsParams {
	o.SetModes(modes)
	return o
}

// SetModes adds the modes to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetModes(modes []int64) {
	o.Modes = modes
}

// WithPeriodType adds the periodType to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) WithPeriodType(periodType *int32) *Destiny2GetHistoricalStatsParams {
	o.SetPeriodType(periodType)
	return o
}

// SetPeriodType adds the periodType to the destiny2 get historical stats params
func (o *Destiny2GetHistoricalStatsParams) SetPeriodType(periodType *int32) {
	o.PeriodType = periodType
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetHistoricalStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param characterId
	if err := r.SetPathParam("characterId", swag.FormatInt64(o.CharacterID)); err != nil {
		return err
	}

	if o.Dayend != nil {

		// query param dayend
		var qrDayend strfmt.DateTime
		if o.Dayend != nil {
			qrDayend = *o.Dayend
		}
		qDayend := qrDayend.String()
		if qDayend != "" {
			if err := r.SetQueryParam("dayend", qDayend); err != nil {
				return err
			}
		}

	}

	if o.Daystart != nil {

		// query param daystart
		var qrDaystart strfmt.DateTime
		if o.Daystart != nil {
			qrDaystart = *o.Daystart
		}
		qDaystart := qrDaystart.String()
		if qDaystart != "" {
			if err := r.SetQueryParam("daystart", qDaystart); err != nil {
				return err
			}
		}

	}

	// path param destinyMembershipId
	if err := r.SetPathParam("destinyMembershipId", swag.FormatInt64(o.DestinyMembershipID)); err != nil {
		return err
	}

	var valuesGroups []string
	for _, v := range o.Groups {
		valuesGroups = append(valuesGroups, swag.FormatInt64(v))
	}

	joinedGroups := swag.JoinByFormat(valuesGroups, "csv")
	// query array param groups
	if err := r.SetQueryParam("groups", joinedGroups...); err != nil {
		return err
	}

	// path param membershipType
	if err := r.SetPathParam("membershipType", swag.FormatInt32(o.MembershipType)); err != nil {
		return err
	}

	var valuesModes []string
	for _, v := range o.Modes {
		valuesModes = append(valuesModes, swag.FormatInt64(v))
	}

	joinedModes := swag.JoinByFormat(valuesModes, "csv")
	// query array param modes
	if err := r.SetQueryParam("modes", joinedModes...); err != nil {
		return err
	}

	if o.PeriodType != nil {

		// query param periodType
		var qrPeriodType int32
		if o.PeriodType != nil {
			qrPeriodType = *o.PeriodType
		}
		qPeriodType := swag.FormatInt32(qrPeriodType)
		if qPeriodType != "" {
			if err := r.SetQueryParam("periodType", qPeriodType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
