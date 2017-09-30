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

// NewGroupV2GetInvitedIndividualsParams creates a new GroupV2GetInvitedIndividualsParams object
// with the default values initialized.
func NewGroupV2GetInvitedIndividualsParams() *GroupV2GetInvitedIndividualsParams {
	var ()
	return &GroupV2GetInvitedIndividualsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2GetInvitedIndividualsParamsWithTimeout creates a new GroupV2GetInvitedIndividualsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2GetInvitedIndividualsParamsWithTimeout(timeout time.Duration) *GroupV2GetInvitedIndividualsParams {
	var ()
	return &GroupV2GetInvitedIndividualsParams{

		timeout: timeout,
	}
}

// NewGroupV2GetInvitedIndividualsParamsWithContext creates a new GroupV2GetInvitedIndividualsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2GetInvitedIndividualsParamsWithContext(ctx context.Context) *GroupV2GetInvitedIndividualsParams {
	var ()
	return &GroupV2GetInvitedIndividualsParams{

		Context: ctx,
	}
}

// NewGroupV2GetInvitedIndividualsParamsWithHTTPClient creates a new GroupV2GetInvitedIndividualsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2GetInvitedIndividualsParamsWithHTTPClient(client *http.Client) *GroupV2GetInvitedIndividualsParams {
	var ()
	return &GroupV2GetInvitedIndividualsParams{
		HTTPClient: client,
	}
}

/*GroupV2GetInvitedIndividualsParams contains all the parameters to send to the API endpoint
for the group v2 get invited individuals operation typically these are written to a http.Request
*/
type GroupV2GetInvitedIndividualsParams struct {

	/*Currentpage
	  Page number (starting with 1). Each page has a fixed size of 50 items per page.

	*/
	Currentpage int32
	/*GroupID
	  ID of the group.

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) WithTimeout(timeout time.Duration) *GroupV2GetInvitedIndividualsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) WithContext(ctx context.Context) *GroupV2GetInvitedIndividualsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) WithHTTPClient(client *http.Client) *GroupV2GetInvitedIndividualsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrentpage adds the currentpage to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) WithCurrentpage(currentpage int32) *GroupV2GetInvitedIndividualsParams {
	o.SetCurrentpage(currentpage)
	return o
}

// SetCurrentpage adds the currentpage to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) SetCurrentpage(currentpage int32) {
	o.Currentpage = currentpage
}

// WithGroupID adds the groupID to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) WithGroupID(groupID int64) *GroupV2GetInvitedIndividualsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 get invited individuals params
func (o *GroupV2GetInvitedIndividualsParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2GetInvitedIndividualsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}