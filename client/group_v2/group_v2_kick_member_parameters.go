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

// NewGroupV2KickMemberParams creates a new GroupV2KickMemberParams object
// with the default values initialized.
func NewGroupV2KickMemberParams() *GroupV2KickMemberParams {
	var ()
	return &GroupV2KickMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2KickMemberParamsWithTimeout creates a new GroupV2KickMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2KickMemberParamsWithTimeout(timeout time.Duration) *GroupV2KickMemberParams {
	var ()
	return &GroupV2KickMemberParams{

		timeout: timeout,
	}
}

// NewGroupV2KickMemberParamsWithContext creates a new GroupV2KickMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2KickMemberParamsWithContext(ctx context.Context) *GroupV2KickMemberParams {
	var ()
	return &GroupV2KickMemberParams{

		Context: ctx,
	}
}

// NewGroupV2KickMemberParamsWithHTTPClient creates a new GroupV2KickMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2KickMemberParamsWithHTTPClient(client *http.Client) *GroupV2KickMemberParams {
	var ()
	return &GroupV2KickMemberParams{
		HTTPClient: client,
	}
}

/*GroupV2KickMemberParams contains all the parameters to send to the API endpoint
for the group v2 kick member operation typically these are written to a http.Request
*/
type GroupV2KickMemberParams struct {

	/*GroupID
	  Group ID to kick the user from.

	*/
	GroupID int64
	/*MembershipID
	  Membership ID to kick.

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

// WithTimeout adds the timeout to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithTimeout(timeout time.Duration) *GroupV2KickMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithContext(ctx context.Context) *GroupV2KickMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithHTTPClient(client *http.Client) *GroupV2KickMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithGroupID(groupID int64) *GroupV2KickMemberParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithMembershipID adds the membershipID to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithMembershipID(membershipID int64) *GroupV2KickMemberParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WithMembershipType adds the membershipType to the group v2 kick member params
func (o *GroupV2KickMemberParams) WithMembershipType(membershipType int32) *GroupV2KickMemberParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the group v2 kick member params
func (o *GroupV2KickMemberParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2KickMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
