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

// NewGroupV2GetGroupsForMemberParams creates a new GroupV2GetGroupsForMemberParams object
// with the default values initialized.
func NewGroupV2GetGroupsForMemberParams() *GroupV2GetGroupsForMemberParams {
	var ()
	return &GroupV2GetGroupsForMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2GetGroupsForMemberParamsWithTimeout creates a new GroupV2GetGroupsForMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2GetGroupsForMemberParamsWithTimeout(timeout time.Duration) *GroupV2GetGroupsForMemberParams {
	var ()
	return &GroupV2GetGroupsForMemberParams{

		timeout: timeout,
	}
}

// NewGroupV2GetGroupsForMemberParamsWithContext creates a new GroupV2GetGroupsForMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2GetGroupsForMemberParamsWithContext(ctx context.Context) *GroupV2GetGroupsForMemberParams {
	var ()
	return &GroupV2GetGroupsForMemberParams{

		Context: ctx,
	}
}

// NewGroupV2GetGroupsForMemberParamsWithHTTPClient creates a new GroupV2GetGroupsForMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2GetGroupsForMemberParamsWithHTTPClient(client *http.Client) *GroupV2GetGroupsForMemberParams {
	var ()
	return &GroupV2GetGroupsForMemberParams{
		HTTPClient: client,
	}
}

/*GroupV2GetGroupsForMemberParams contains all the parameters to send to the API endpoint
for the group v2 get groups for member operation typically these are written to a http.Request
*/
type GroupV2GetGroupsForMemberParams struct {

	/*Filter
	  Filter apply to list of joined groups.

	*/
	Filter int32
	/*GroupType
	  Type of group the supplied member founded.

	*/
	GroupType int32
	/*MembershipID
	  Membership ID to for which to find founded groups.

	*/
	MembershipID int64
	/*MembershipType
	  Membership type of the supplied membership ID.

	*/
	MembershipType int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithTimeout(timeout time.Duration) *GroupV2GetGroupsForMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithContext(ctx context.Context) *GroupV2GetGroupsForMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithHTTPClient(client *http.Client) *GroupV2GetGroupsForMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithFilter(filter int32) *GroupV2GetGroupsForMemberParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetFilter(filter int32) {
	o.Filter = filter
}

// WithGroupType adds the groupType to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithGroupType(groupType int32) *GroupV2GetGroupsForMemberParams {
	o.SetGroupType(groupType)
	return o
}

// SetGroupType adds the groupType to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetGroupType(groupType int32) {
	o.GroupType = groupType
}

// WithMembershipID adds the membershipID to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithMembershipID(membershipID int64) *GroupV2GetGroupsForMemberParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WithMembershipType adds the membershipType to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) WithMembershipType(membershipType int32) *GroupV2GetGroupsForMemberParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the group v2 get groups for member params
func (o *GroupV2GetGroupsForMemberParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2GetGroupsForMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filter
	if err := r.SetPathParam("filter", swag.FormatInt32(o.Filter)); err != nil {
		return err
	}

	// path param groupType
	if err := r.SetPathParam("groupType", swag.FormatInt32(o.GroupType)); err != nil {
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
