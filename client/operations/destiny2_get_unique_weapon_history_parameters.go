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

// NewDestiny2GetUniqueWeaponHistoryParams creates a new Destiny2GetUniqueWeaponHistoryParams object
// with the default values initialized.
func NewDestiny2GetUniqueWeaponHistoryParams() *Destiny2GetUniqueWeaponHistoryParams {
	var ()
	return &Destiny2GetUniqueWeaponHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetUniqueWeaponHistoryParamsWithTimeout creates a new Destiny2GetUniqueWeaponHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetUniqueWeaponHistoryParamsWithTimeout(timeout time.Duration) *Destiny2GetUniqueWeaponHistoryParams {
	var ()
	return &Destiny2GetUniqueWeaponHistoryParams{

		timeout: timeout,
	}
}

// NewDestiny2GetUniqueWeaponHistoryParamsWithContext creates a new Destiny2GetUniqueWeaponHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetUniqueWeaponHistoryParamsWithContext(ctx context.Context) *Destiny2GetUniqueWeaponHistoryParams {
	var ()
	return &Destiny2GetUniqueWeaponHistoryParams{

		Context: ctx,
	}
}

// NewDestiny2GetUniqueWeaponHistoryParamsWithHTTPClient creates a new Destiny2GetUniqueWeaponHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetUniqueWeaponHistoryParamsWithHTTPClient(client *http.Client) *Destiny2GetUniqueWeaponHistoryParams {
	var ()
	return &Destiny2GetUniqueWeaponHistoryParams{
		HTTPClient: client,
	}
}

/*Destiny2GetUniqueWeaponHistoryParams contains all the parameters to send to the API endpoint
for the destiny2 get unique weapon history operation typically these are written to a http.Request
*/
type Destiny2GetUniqueWeaponHistoryParams struct {

	/*CharacterID
	  The id of the character to retrieve.

	*/
	CharacterID int64
	/*DestinyMembershipID
	  The Destiny membershipId of the user to retrieve.

	*/
	DestinyMembershipID int64
	/*MembershipType
	  A valid non-BungieNet membership type.

	*/
	MembershipType int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithTimeout(timeout time.Duration) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithContext(ctx context.Context) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithHTTPClient(client *http.Client) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithCharacterID(characterID int64) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetCharacterID(characterID int64) {
	o.CharacterID = characterID
}

// WithDestinyMembershipID adds the destinyMembershipID to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithDestinyMembershipID(destinyMembershipID int64) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetDestinyMembershipID(destinyMembershipID)
	return o
}

// SetDestinyMembershipID adds the destinyMembershipId to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetDestinyMembershipID(destinyMembershipID int64) {
	o.DestinyMembershipID = destinyMembershipID
}

// WithMembershipType adds the membershipType to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) WithMembershipType(membershipType int32) *Destiny2GetUniqueWeaponHistoryParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the destiny2 get unique weapon history params
func (o *Destiny2GetUniqueWeaponHistoryParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetUniqueWeaponHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param characterId
	if err := r.SetPathParam("characterId", swag.FormatInt64(o.CharacterID)); err != nil {
		return err
	}

	// path param destinyMembershipId
	if err := r.SetPathParam("destinyMembershipId", swag.FormatInt64(o.DestinyMembershipID)); err != nil {
		return err
	}

	// path param membershipType
	if err := r.SetPathParam("membershipType", swag.FormatInt32(o.MembershipType)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}