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

// NewGroupV2EditGroupMembershipParams creates a new GroupV2EditGroupMembershipParams object
// with the default values initialized.
func NewGroupV2EditGroupMembershipParams() *GroupV2EditGroupMembershipParams {
	var ()
	return &GroupV2EditGroupMembershipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2EditGroupMembershipParamsWithTimeout creates a new GroupV2EditGroupMembershipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2EditGroupMembershipParamsWithTimeout(timeout time.Duration) *GroupV2EditGroupMembershipParams {
	var ()
	return &GroupV2EditGroupMembershipParams{

		timeout: timeout,
	}
}

// NewGroupV2EditGroupMembershipParamsWithContext creates a new GroupV2EditGroupMembershipParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2EditGroupMembershipParamsWithContext(ctx context.Context) *GroupV2EditGroupMembershipParams {
	var ()
	return &GroupV2EditGroupMembershipParams{

		Context: ctx,
	}
}

// NewGroupV2EditGroupMembershipParamsWithHTTPClient creates a new GroupV2EditGroupMembershipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2EditGroupMembershipParamsWithHTTPClient(client *http.Client) *GroupV2EditGroupMembershipParams {
	var ()
	return &GroupV2EditGroupMembershipParams{
		HTTPClient: client,
	}
}

/*GroupV2EditGroupMembershipParams contains all the parameters to send to the API endpoint
for the group v2 edit group membership operation typically these are written to a http.Request
*/
type GroupV2EditGroupMembershipParams struct {

	/*GroupID
	  ID of the group to which the member belongs.

	*/
	GroupID int64
	/*MemberType
	  New membertype for the specified member.

	*/
	MemberType int32
	/*MembershipID
	  Membership ID to modify.

	*/
	MembershipID int64
	/*MembershipType
	  Membership type of the provide membership ID.

	*/
	MembershipType int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithTimeout(timeout time.Duration) *GroupV2EditGroupMembershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithContext(ctx context.Context) *GroupV2EditGroupMembershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithHTTPClient(client *http.Client) *GroupV2EditGroupMembershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithGroupID(groupID int64) *GroupV2EditGroupMembershipParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithMemberType adds the memberType to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithMemberType(memberType int32) *GroupV2EditGroupMembershipParams {
	o.SetMemberType(memberType)
	return o
}

// SetMemberType adds the memberType to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetMemberType(memberType int32) {
	o.MemberType = memberType
}

// WithMembershipID adds the membershipID to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithMembershipID(membershipID int64) *GroupV2EditGroupMembershipParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WithMembershipType adds the membershipType to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) WithMembershipType(membershipType int32) *GroupV2EditGroupMembershipParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the group v2 edit group membership params
func (o *GroupV2EditGroupMembershipParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2EditGroupMembershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	// path param memberType
	if err := r.SetPathParam("memberType", swag.FormatInt32(o.MemberType)); err != nil {
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
