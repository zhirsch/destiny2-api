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

// NewGroupV2ApproveAllPendingParams creates a new GroupV2ApproveAllPendingParams object
// with the default values initialized.
func NewGroupV2ApproveAllPendingParams() *GroupV2ApproveAllPendingParams {
	var ()
	return &GroupV2ApproveAllPendingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2ApproveAllPendingParamsWithTimeout creates a new GroupV2ApproveAllPendingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2ApproveAllPendingParamsWithTimeout(timeout time.Duration) *GroupV2ApproveAllPendingParams {
	var ()
	return &GroupV2ApproveAllPendingParams{

		timeout: timeout,
	}
}

// NewGroupV2ApproveAllPendingParamsWithContext creates a new GroupV2ApproveAllPendingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2ApproveAllPendingParamsWithContext(ctx context.Context) *GroupV2ApproveAllPendingParams {
	var ()
	return &GroupV2ApproveAllPendingParams{

		Context: ctx,
	}
}

// NewGroupV2ApproveAllPendingParamsWithHTTPClient creates a new GroupV2ApproveAllPendingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2ApproveAllPendingParamsWithHTTPClient(client *http.Client) *GroupV2ApproveAllPendingParams {
	var ()
	return &GroupV2ApproveAllPendingParams{
		HTTPClient: client,
	}
}

/*GroupV2ApproveAllPendingParams contains all the parameters to send to the API endpoint
for the group v2 approve all pending operation typically these are written to a http.Request
*/
type GroupV2ApproveAllPendingParams struct {

	/*GroupID
	  ID of the group.

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) WithTimeout(timeout time.Duration) *GroupV2ApproveAllPendingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) WithContext(ctx context.Context) *GroupV2ApproveAllPendingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) WithHTTPClient(client *http.Client) *GroupV2ApproveAllPendingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) WithGroupID(groupID int64) *GroupV2ApproveAllPendingParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 approve all pending params
func (o *GroupV2ApproveAllPendingParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2ApproveAllPendingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}