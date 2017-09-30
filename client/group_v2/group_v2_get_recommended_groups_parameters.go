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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGroupV2GetRecommendedGroupsParams creates a new GroupV2GetRecommendedGroupsParams object
// with the default values initialized.
func NewGroupV2GetRecommendedGroupsParams() *GroupV2GetRecommendedGroupsParams {

	return &GroupV2GetRecommendedGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2GetRecommendedGroupsParamsWithTimeout creates a new GroupV2GetRecommendedGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2GetRecommendedGroupsParamsWithTimeout(timeout time.Duration) *GroupV2GetRecommendedGroupsParams {

	return &GroupV2GetRecommendedGroupsParams{

		timeout: timeout,
	}
}

// NewGroupV2GetRecommendedGroupsParamsWithContext creates a new GroupV2GetRecommendedGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2GetRecommendedGroupsParamsWithContext(ctx context.Context) *GroupV2GetRecommendedGroupsParams {

	return &GroupV2GetRecommendedGroupsParams{

		Context: ctx,
	}
}

// NewGroupV2GetRecommendedGroupsParamsWithHTTPClient creates a new GroupV2GetRecommendedGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2GetRecommendedGroupsParamsWithHTTPClient(client *http.Client) *GroupV2GetRecommendedGroupsParams {

	return &GroupV2GetRecommendedGroupsParams{
		HTTPClient: client,
	}
}

/*GroupV2GetRecommendedGroupsParams contains all the parameters to send to the API endpoint
for the group v2 get recommended groups operation typically these are written to a http.Request
*/
type GroupV2GetRecommendedGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) WithTimeout(timeout time.Duration) *GroupV2GetRecommendedGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) WithContext(ctx context.Context) *GroupV2GetRecommendedGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) WithHTTPClient(client *http.Client) *GroupV2GetRecommendedGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 get recommended groups params
func (o *GroupV2GetRecommendedGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2GetRecommendedGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
