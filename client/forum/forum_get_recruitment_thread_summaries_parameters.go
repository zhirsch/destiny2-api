// Code generated by go-swagger; DO NOT EDIT.

package forum

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

// NewForumGetRecruitmentThreadSummariesParams creates a new ForumGetRecruitmentThreadSummariesParams object
// with the default values initialized.
func NewForumGetRecruitmentThreadSummariesParams() *ForumGetRecruitmentThreadSummariesParams {

	return &ForumGetRecruitmentThreadSummariesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewForumGetRecruitmentThreadSummariesParamsWithTimeout creates a new ForumGetRecruitmentThreadSummariesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForumGetRecruitmentThreadSummariesParamsWithTimeout(timeout time.Duration) *ForumGetRecruitmentThreadSummariesParams {

	return &ForumGetRecruitmentThreadSummariesParams{

		timeout: timeout,
	}
}

// NewForumGetRecruitmentThreadSummariesParamsWithContext creates a new ForumGetRecruitmentThreadSummariesParams object
// with the default values initialized, and the ability to set a context for a request
func NewForumGetRecruitmentThreadSummariesParamsWithContext(ctx context.Context) *ForumGetRecruitmentThreadSummariesParams {

	return &ForumGetRecruitmentThreadSummariesParams{

		Context: ctx,
	}
}

// NewForumGetRecruitmentThreadSummariesParamsWithHTTPClient creates a new ForumGetRecruitmentThreadSummariesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForumGetRecruitmentThreadSummariesParamsWithHTTPClient(client *http.Client) *ForumGetRecruitmentThreadSummariesParams {

	return &ForumGetRecruitmentThreadSummariesParams{
		HTTPClient: client,
	}
}

/*ForumGetRecruitmentThreadSummariesParams contains all the parameters to send to the API endpoint
for the forum get recruitment thread summaries operation typically these are written to a http.Request
*/
type ForumGetRecruitmentThreadSummariesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) WithTimeout(timeout time.Duration) *ForumGetRecruitmentThreadSummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) WithContext(ctx context.Context) *ForumGetRecruitmentThreadSummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) WithHTTPClient(client *http.Client) *ForumGetRecruitmentThreadSummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forum get recruitment thread summaries params
func (o *ForumGetRecruitmentThreadSummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ForumGetRecruitmentThreadSummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
