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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewForumGetPostsThreadedPagedFromChildParams creates a new ForumGetPostsThreadedPagedFromChildParams object
// with the default values initialized.
func NewForumGetPostsThreadedPagedFromChildParams() *ForumGetPostsThreadedPagedFromChildParams {
	var ()
	return &ForumGetPostsThreadedPagedFromChildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewForumGetPostsThreadedPagedFromChildParamsWithTimeout creates a new ForumGetPostsThreadedPagedFromChildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForumGetPostsThreadedPagedFromChildParamsWithTimeout(timeout time.Duration) *ForumGetPostsThreadedPagedFromChildParams {
	var ()
	return &ForumGetPostsThreadedPagedFromChildParams{

		timeout: timeout,
	}
}

// NewForumGetPostsThreadedPagedFromChildParamsWithContext creates a new ForumGetPostsThreadedPagedFromChildParams object
// with the default values initialized, and the ability to set a context for a request
func NewForumGetPostsThreadedPagedFromChildParamsWithContext(ctx context.Context) *ForumGetPostsThreadedPagedFromChildParams {
	var ()
	return &ForumGetPostsThreadedPagedFromChildParams{

		Context: ctx,
	}
}

// NewForumGetPostsThreadedPagedFromChildParamsWithHTTPClient creates a new ForumGetPostsThreadedPagedFromChildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForumGetPostsThreadedPagedFromChildParamsWithHTTPClient(client *http.Client) *ForumGetPostsThreadedPagedFromChildParams {
	var ()
	return &ForumGetPostsThreadedPagedFromChildParams{
		HTTPClient: client,
	}
}

/*ForumGetPostsThreadedPagedFromChildParams contains all the parameters to send to the API endpoint
for the forum get posts threaded paged from child operation typically these are written to a http.Request
*/
type ForumGetPostsThreadedPagedFromChildParams struct {

	/*ChildPostID*/
	ChildPostID uint64
	/*Page*/
	Page int32
	/*PageSize*/
	PageSize int32
	/*ReplySize*/
	ReplySize int32
	/*RootThreadMode*/
	RootThreadMode bool
	/*Showbanned
	  If this value is not null or empty, banned posts are requested to be returned

	*/
	Showbanned *string
	/*SortMode*/
	SortMode int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithTimeout(timeout time.Duration) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithContext(ctx context.Context) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithHTTPClient(client *http.Client) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildPostID adds the childPostID to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithChildPostID(childPostID uint64) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetChildPostID(childPostID)
	return o
}

// SetChildPostID adds the childPostId to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetChildPostID(childPostID uint64) {
	o.ChildPostID = childPostID
}

// WithPage adds the page to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithPage(page int32) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetPage(page int32) {
	o.Page = page
}

// WithPageSize adds the pageSize to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithPageSize(pageSize int32) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetPageSize(pageSize int32) {
	o.PageSize = pageSize
}

// WithReplySize adds the replySize to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithReplySize(replySize int32) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetReplySize(replySize)
	return o
}

// SetReplySize adds the replySize to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetReplySize(replySize int32) {
	o.ReplySize = replySize
}

// WithRootThreadMode adds the rootThreadMode to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithRootThreadMode(rootThreadMode bool) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetRootThreadMode(rootThreadMode)
	return o
}

// SetRootThreadMode adds the rootThreadMode to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetRootThreadMode(rootThreadMode bool) {
	o.RootThreadMode = rootThreadMode
}

// WithShowbanned adds the showbanned to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithShowbanned(showbanned *string) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetShowbanned(showbanned)
	return o
}

// SetShowbanned adds the showbanned to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetShowbanned(showbanned *string) {
	o.Showbanned = showbanned
}

// WithSortMode adds the sortMode to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) WithSortMode(sortMode int32) *ForumGetPostsThreadedPagedFromChildParams {
	o.SetSortMode(sortMode)
	return o
}

// SetSortMode adds the sortMode to the forum get posts threaded paged from child params
func (o *ForumGetPostsThreadedPagedFromChildParams) SetSortMode(sortMode int32) {
	o.SortMode = sortMode
}

// WriteToRequest writes these params to a swagger request
func (o *ForumGetPostsThreadedPagedFromChildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childPostId
	if err := r.SetPathParam("childPostId", swag.FormatUint64(o.ChildPostID)); err != nil {
		return err
	}

	// path param page
	if err := r.SetPathParam("page", swag.FormatInt32(o.Page)); err != nil {
		return err
	}

	// path param pageSize
	if err := r.SetPathParam("pageSize", swag.FormatInt32(o.PageSize)); err != nil {
		return err
	}

	// path param replySize
	if err := r.SetPathParam("replySize", swag.FormatInt32(o.ReplySize)); err != nil {
		return err
	}

	// path param rootThreadMode
	if err := r.SetPathParam("rootThreadMode", swag.FormatBool(o.RootThreadMode)); err != nil {
		return err
	}

	if o.Showbanned != nil {

		// query param showbanned
		var qrShowbanned string
		if o.Showbanned != nil {
			qrShowbanned = *o.Showbanned
		}
		qShowbanned := qrShowbanned
		if qShowbanned != "" {
			if err := r.SetQueryParam("showbanned", qShowbanned); err != nil {
				return err
			}
		}

	}

	// path param sortMode
	if err := r.SetPathParam("sortMode", swag.FormatInt32(o.SortMode)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
