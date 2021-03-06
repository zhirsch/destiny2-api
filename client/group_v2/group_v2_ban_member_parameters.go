// Code generated by go-swagger; DO NOT EDIT.

package group_v2

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

// NewGroupV2BanMemberParams creates a new GroupV2BanMemberParams object
// with the default values initialized.
func NewGroupV2BanMemberParams() *GroupV2BanMemberParams {
	var ()
	return &GroupV2BanMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2BanMemberParamsWithTimeout creates a new GroupV2BanMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2BanMemberParamsWithTimeout(timeout time.Duration) *GroupV2BanMemberParams {
	var ()
	return &GroupV2BanMemberParams{

		timeout: timeout,
	}
}

// NewGroupV2BanMemberParamsWithContext creates a new GroupV2BanMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2BanMemberParamsWithContext(ctx context.Context) *GroupV2BanMemberParams {
	var ()
	return &GroupV2BanMemberParams{

		Context: ctx,
	}
}

// NewGroupV2BanMemberParamsWithHTTPClient creates a new GroupV2BanMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2BanMemberParamsWithHTTPClient(client *http.Client) *GroupV2BanMemberParams {
	var ()
	return &GroupV2BanMemberParams{
		HTTPClient: client,
	}
}

/*GroupV2BanMemberParams contains all the parameters to send to the API endpoint
for the group v2 ban member operation typically these are written to a http.Request
*/
type GroupV2BanMemberParams struct {

	/*GroupID
	  Group ID that has the member to ban.

	*/
	GroupID int64
	/*MembershipID
	  Membership ID of the member to ban from the group.

	*/
	MembershipID int64
	/*MembershipType
	  Membership type of the provided membership ID.

	*/
	MembershipType int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithTimeout(timeout time.Duration) *GroupV2BanMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithContext(ctx context.Context) *GroupV2BanMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithHTTPClient(client *http.Client) *GroupV2BanMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithGroupID(groupID int64) *GroupV2BanMemberParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithMembershipID adds the membershipID to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithMembershipID(membershipID int64) *GroupV2BanMemberParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WithMembershipType adds the membershipType to the group v2 ban member params
func (o *GroupV2BanMemberParams) WithMembershipType(membershipType int32) *GroupV2BanMemberParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the group v2 ban member params
func (o *GroupV2BanMemberParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2BanMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	// path param membershipId
	if err := r.SetPathParam("membershipId", swag.FormatInt64(o.MembershipID)); err != nil {
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
