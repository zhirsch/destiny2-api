// Code generated by go-swagger; DO NOT EDIT.

package destiny2

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

// NewDestiny2GetCharacterParams creates a new Destiny2GetCharacterParams object
// with the default values initialized.
func NewDestiny2GetCharacterParams() *Destiny2GetCharacterParams {
	var ()
	return &Destiny2GetCharacterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetCharacterParamsWithTimeout creates a new Destiny2GetCharacterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetCharacterParamsWithTimeout(timeout time.Duration) *Destiny2GetCharacterParams {
	var ()
	return &Destiny2GetCharacterParams{

		timeout: timeout,
	}
}

// NewDestiny2GetCharacterParamsWithContext creates a new Destiny2GetCharacterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetCharacterParamsWithContext(ctx context.Context) *Destiny2GetCharacterParams {
	var ()
	return &Destiny2GetCharacterParams{

		Context: ctx,
	}
}

// NewDestiny2GetCharacterParamsWithHTTPClient creates a new Destiny2GetCharacterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetCharacterParamsWithHTTPClient(client *http.Client) *Destiny2GetCharacterParams {
	var ()
	return &Destiny2GetCharacterParams{
		HTTPClient: client,
	}
}

/*Destiny2GetCharacterParams contains all the parameters to send to the API endpoint
for the destiny2 get character operation typically these are written to a http.Request
*/
type Destiny2GetCharacterParams struct {

	/*CharacterID
	  ID of the character.

	*/
	CharacterID int64
	/*Components
	  A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results.

	*/
	Components []int64
	/*DestinyMembershipID
	  Destiny membership ID.

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

// WithTimeout adds the timeout to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithTimeout(timeout time.Duration) *Destiny2GetCharacterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithContext(ctx context.Context) *Destiny2GetCharacterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithHTTPClient(client *http.Client) *Destiny2GetCharacterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithCharacterID(characterID int64) *Destiny2GetCharacterParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetCharacterID(characterID int64) {
	o.CharacterID = characterID
}

// WithComponents adds the components to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithComponents(components []int64) *Destiny2GetCharacterParams {
	o.SetComponents(components)
	return o
}

// SetComponents adds the components to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetComponents(components []int64) {
	o.Components = components
}

// WithDestinyMembershipID adds the destinyMembershipID to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithDestinyMembershipID(destinyMembershipID int64) *Destiny2GetCharacterParams {
	o.SetDestinyMembershipID(destinyMembershipID)
	return o
}

// SetDestinyMembershipID adds the destinyMembershipId to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetDestinyMembershipID(destinyMembershipID int64) {
	o.DestinyMembershipID = destinyMembershipID
}

// WithMembershipType adds the membershipType to the destiny2 get character params
func (o *Destiny2GetCharacterParams) WithMembershipType(membershipType int32) *Destiny2GetCharacterParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the destiny2 get character params
func (o *Destiny2GetCharacterParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetCharacterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param characterId
	if err := r.SetPathParam("characterId", swag.FormatInt64(o.CharacterID)); err != nil {
		return err
	}

	var valuesComponents []string
	for _, v := range o.Components {
		valuesComponents = append(valuesComponents, swag.FormatInt64(v))
	}

	joinedComponents := swag.JoinByFormat(valuesComponents, "csv")
	// query array param components
	if err := r.SetQueryParam("components", joinedComponents...); err != nil {
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
