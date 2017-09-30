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

// NewGroupV2GetMembersOfGroupParams creates a new GroupV2GetMembersOfGroupParams object
// with the default values initialized.
func NewGroupV2GetMembersOfGroupParams() *GroupV2GetMembersOfGroupParams {
	var ()
	return &GroupV2GetMembersOfGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2GetMembersOfGroupParamsWithTimeout creates a new GroupV2GetMembersOfGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2GetMembersOfGroupParamsWithTimeout(timeout time.Duration) *GroupV2GetMembersOfGroupParams {
	var ()
	return &GroupV2GetMembersOfGroupParams{

		timeout: timeout,
	}
}

// NewGroupV2GetMembersOfGroupParamsWithContext creates a new GroupV2GetMembersOfGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2GetMembersOfGroupParamsWithContext(ctx context.Context) *GroupV2GetMembersOfGroupParams {
	var ()
	return &GroupV2GetMembersOfGroupParams{

		Context: ctx,
	}
}

// NewGroupV2GetMembersOfGroupParamsWithHTTPClient creates a new GroupV2GetMembersOfGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2GetMembersOfGroupParamsWithHTTPClient(client *http.Client) *GroupV2GetMembersOfGroupParams {
	var ()
	return &GroupV2GetMembersOfGroupParams{
		HTTPClient: client,
	}
}

/*GroupV2GetMembersOfGroupParams contains all the parameters to send to the API endpoint
for the group v2 get members of group operation typically these are written to a http.Request
*/
type GroupV2GetMembersOfGroupParams struct {

	/*Currentpage
	  Page number (starting with 1). Each page has a fixed size of 50 items per page.

	*/
	Currentpage int32
	/*GroupID
	  The ID of the group.

	*/
	GroupID int64
	/*MemberType
	  Filter out other member types. Use None for all members.

	*/
	MemberType *int32
	/*NameSearch
	  The name fragment upon which a search should be executed for members with matching display or unique names.

	*/
	NameSearch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithTimeout(timeout time.Duration) *GroupV2GetMembersOfGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithContext(ctx context.Context) *GroupV2GetMembersOfGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithHTTPClient(client *http.Client) *GroupV2GetMembersOfGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrentpage adds the currentpage to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithCurrentpage(currentpage int32) *GroupV2GetMembersOfGroupParams {
	o.SetCurrentpage(currentpage)
	return o
}

// SetCurrentpage adds the currentpage to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetCurrentpage(currentpage int32) {
	o.Currentpage = currentpage
}

// WithGroupID adds the groupID to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithGroupID(groupID int64) *GroupV2GetMembersOfGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WithMemberType adds the memberType to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithMemberType(memberType *int32) *GroupV2GetMembersOfGroupParams {
	o.SetMemberType(memberType)
	return o
}

// SetMemberType adds the memberType to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetMemberType(memberType *int32) {
	o.MemberType = memberType
}

// WithNameSearch adds the nameSearch to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) WithNameSearch(nameSearch *string) *GroupV2GetMembersOfGroupParams {
	o.SetNameSearch(nameSearch)
	return o
}

// SetNameSearch adds the nameSearch to the group v2 get members of group params
func (o *GroupV2GetMembersOfGroupParams) SetNameSearch(nameSearch *string) {
	o.NameSearch = nameSearch
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2GetMembersOfGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param currentpage
	qrCurrentpage := o.Currentpage
	qCurrentpage := swag.FormatInt32(qrCurrentpage)
	if qCurrentpage != "" {
		if err := r.SetQueryParam("currentpage", qCurrentpage); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if o.MemberType != nil {

		// query param memberType
		var qrMemberType int32
		if o.MemberType != nil {
			qrMemberType = *o.MemberType
		}
		qMemberType := swag.FormatInt32(qrMemberType)
		if qMemberType != "" {
			if err := r.SetQueryParam("memberType", qMemberType); err != nil {
				return err
			}
		}

	}

	if o.NameSearch != nil {

		// query param nameSearch
		var qrNameSearch string
		if o.NameSearch != nil {
			qrNameSearch = *o.NameSearch
		}
		qNameSearch := qrNameSearch
		if qNameSearch != "" {
			if err := r.SetQueryParam("nameSearch", qNameSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
