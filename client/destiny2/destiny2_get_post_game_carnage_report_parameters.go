// Code generated by go-swagger; DO NOT EDIT.

package destiny2

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

// NewDestiny2GetPostGameCarnageReportParams creates a new Destiny2GetPostGameCarnageReportParams object
// with the default values initialized.
func NewDestiny2GetPostGameCarnageReportParams() *Destiny2GetPostGameCarnageReportParams {
	var ()
	return &Destiny2GetPostGameCarnageReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetPostGameCarnageReportParamsWithTimeout creates a new Destiny2GetPostGameCarnageReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetPostGameCarnageReportParamsWithTimeout(timeout time.Duration) *Destiny2GetPostGameCarnageReportParams {
	var ()
	return &Destiny2GetPostGameCarnageReportParams{

		timeout: timeout,
	}
}

// NewDestiny2GetPostGameCarnageReportParamsWithContext creates a new Destiny2GetPostGameCarnageReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetPostGameCarnageReportParamsWithContext(ctx context.Context) *Destiny2GetPostGameCarnageReportParams {
	var ()
	return &Destiny2GetPostGameCarnageReportParams{

		Context: ctx,
	}
}

// NewDestiny2GetPostGameCarnageReportParamsWithHTTPClient creates a new Destiny2GetPostGameCarnageReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetPostGameCarnageReportParamsWithHTTPClient(client *http.Client) *Destiny2GetPostGameCarnageReportParams {
	var ()
	return &Destiny2GetPostGameCarnageReportParams{
		HTTPClient: client,
	}
}

/*Destiny2GetPostGameCarnageReportParams contains all the parameters to send to the API endpoint
for the destiny2 get post game carnage report operation typically these are written to a http.Request
*/
type Destiny2GetPostGameCarnageReportParams struct {

	/*ActivityID
	  The ID of the activity whose PGCR is requested.

	*/
	ActivityID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) WithTimeout(timeout time.Duration) *Destiny2GetPostGameCarnageReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) WithContext(ctx context.Context) *Destiny2GetPostGameCarnageReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) WithHTTPClient(client *http.Client) *Destiny2GetPostGameCarnageReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityID adds the activityID to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) WithActivityID(activityID int64) *Destiny2GetPostGameCarnageReportParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the destiny2 get post game carnage report params
func (o *Destiny2GetPostGameCarnageReportParams) SetActivityID(activityID int64) {
	o.ActivityID = activityID
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetPostGameCarnageReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param activityId
	if err := r.SetPathParam("activityId", swag.FormatInt64(o.ActivityID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
