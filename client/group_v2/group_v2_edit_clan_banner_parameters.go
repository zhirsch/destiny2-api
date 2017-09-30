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

// NewGroupV2EditClanBannerParams creates a new GroupV2EditClanBannerParams object
// with the default values initialized.
func NewGroupV2EditClanBannerParams() *GroupV2EditClanBannerParams {
	var ()
	return &GroupV2EditClanBannerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2EditClanBannerParamsWithTimeout creates a new GroupV2EditClanBannerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2EditClanBannerParamsWithTimeout(timeout time.Duration) *GroupV2EditClanBannerParams {
	var ()
	return &GroupV2EditClanBannerParams{

		timeout: timeout,
	}
}

// NewGroupV2EditClanBannerParamsWithContext creates a new GroupV2EditClanBannerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2EditClanBannerParamsWithContext(ctx context.Context) *GroupV2EditClanBannerParams {
	var ()
	return &GroupV2EditClanBannerParams{

		Context: ctx,
	}
}

// NewGroupV2EditClanBannerParamsWithHTTPClient creates a new GroupV2EditClanBannerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2EditClanBannerParamsWithHTTPClient(client *http.Client) *GroupV2EditClanBannerParams {
	var ()
	return &GroupV2EditClanBannerParams{
		HTTPClient: client,
	}
}

/*GroupV2EditClanBannerParams contains all the parameters to send to the API endpoint
for the group v2 edit clan banner operation typically these are written to a http.Request
*/
type GroupV2EditClanBannerParams struct {

	/*GroupID
	  Group ID of the group to edit.

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) WithTimeout(timeout time.Duration) *GroupV2EditClanBannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) WithContext(ctx context.Context) *GroupV2EditClanBannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) WithHTTPClient(client *http.Client) *GroupV2EditClanBannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) WithGroupID(groupID int64) *GroupV2EditClanBannerParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 edit clan banner params
func (o *GroupV2EditClanBannerParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2EditClanBannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
