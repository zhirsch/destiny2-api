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

// NewGroupV2IndividualGroupInviteParams creates a new GroupV2IndividualGroupInviteParams object
// with the default values initialized.
func NewGroupV2IndividualGroupInviteParams() *GroupV2IndividualGroupInviteParams {
	var ()
	return &GroupV2IndividualGroupInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2IndividualGroupInviteParamsWithTimeout creates a new GroupV2IndividualGroupInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2IndividualGroupInviteParamsWithTimeout(timeout time.Duration) *GroupV2IndividualGroupInviteParams {
	var ()
	return &GroupV2IndividualGroupInviteParams{

		timeout: timeout,
	}
}

// NewGroupV2IndividualGroupInviteParamsWithContext creates a new GroupV2IndividualGroupInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2IndividualGroupInviteParamsWithContext(ctx context.Context) *GroupV2IndividualGroupInviteParams {
	var ()
	return &GroupV2IndividualGroupInviteParams{

		Context: ctx,
	}
}

// NewGroupV2IndividualGroupInviteParamsWithHTTPClient creates a new GroupV2IndividualGroupInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2IndividualGroupInviteParamsWithHTTPClient(client *http.Client) *GroupV2IndividualGroupInviteParams {
	var ()
	return &GroupV2IndividualGroupInviteParams{
		HTTPClient: client,
	}
}

/*GroupV2IndividualGroupInviteParams contains all the parameters to send to the API endpoint
for the group v2 individual group invite operation typically these are written to a http.Request
*/
type GroupV2IndividualGroupInviteParams struct {

	/*GroupID
	  ID of the group you would like to join.

	*/
	GroupID int64
	/*MembershipID
	  Membership id of the account being invited.

	*/
	MembershipID int64
	/*MembershipType
	  MembershipType of the account being invited.

	*/
	MembershipType int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithTimeout(timeout time.Duration) *GroupV2IndividualGroupInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithContext(ctx context.Context) *GroupV2IndividualGroupInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithHTTPClient(client *http.Client) *GroupV2IndividualGroupInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithGroupID(groupID int64) *GroupV2IndividualGroupInviteParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithMembershipID adds the membershipID to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithMembershipID(membershipID int64) *GroupV2IndividualGroupInviteParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WithMembershipType adds the membershipType to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) WithMembershipType(membershipType int32) *GroupV2IndividualGroupInviteParams {
	o.SetMembershipType(membershipType)
	return o
}

// SetMembershipType adds the membershipType to the group v2 individual group invite params
func (o *GroupV2IndividualGroupInviteParams) SetMembershipType(membershipType int32) {
	o.MembershipType = membershipType
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2IndividualGroupInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
