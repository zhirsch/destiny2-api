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

// NewGroupV2EditOptionalConversationParams creates a new GroupV2EditOptionalConversationParams object
// with the default values initialized.
func NewGroupV2EditOptionalConversationParams() *GroupV2EditOptionalConversationParams {
	var ()
	return &GroupV2EditOptionalConversationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupV2EditOptionalConversationParamsWithTimeout creates a new GroupV2EditOptionalConversationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupV2EditOptionalConversationParamsWithTimeout(timeout time.Duration) *GroupV2EditOptionalConversationParams {
	var ()
	return &GroupV2EditOptionalConversationParams{

		timeout: timeout,
	}
}

// NewGroupV2EditOptionalConversationParamsWithContext creates a new GroupV2EditOptionalConversationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupV2EditOptionalConversationParamsWithContext(ctx context.Context) *GroupV2EditOptionalConversationParams {
	var ()
	return &GroupV2EditOptionalConversationParams{

		Context: ctx,
	}
}

// NewGroupV2EditOptionalConversationParamsWithHTTPClient creates a new GroupV2EditOptionalConversationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupV2EditOptionalConversationParamsWithHTTPClient(client *http.Client) *GroupV2EditOptionalConversationParams {
	var ()
	return &GroupV2EditOptionalConversationParams{
		HTTPClient: client,
	}
}

/*GroupV2EditOptionalConversationParams contains all the parameters to send to the API endpoint
for the group v2 edit optional conversation operation typically these are written to a http.Request
*/
type GroupV2EditOptionalConversationParams struct {

	/*ConversationID
	  Conversation Id of the channel being edited.

	*/
	ConversationID int64
	/*GroupID
	  Group ID of the group to edit.

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) WithTimeout(timeout time.Duration) *GroupV2EditOptionalConversationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) WithContext(ctx context.Context) *GroupV2EditOptionalConversationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) WithHTTPClient(client *http.Client) *GroupV2EditOptionalConversationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConversationID adds the conversationID to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) WithConversationID(conversationID int64) *GroupV2EditOptionalConversationParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) SetConversationID(conversationID int64) {
	o.ConversationID = conversationID
}

// WithGroupID adds the groupID to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) WithGroupID(groupID int64) *GroupV2EditOptionalConversationParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the group v2 edit optional conversation params
func (o *GroupV2EditOptionalConversationParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GroupV2EditOptionalConversationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param conversationId
	if err := r.SetPathParam("conversationId", swag.FormatInt64(o.ConversationID)); err != nil {
		return err
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
