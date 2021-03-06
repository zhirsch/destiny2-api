// Code generated by go-swagger; DO NOT EDIT.

package trending

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

// NewTrendingGetTrendingCategoryParams creates a new TrendingGetTrendingCategoryParams object
// with the default values initialized.
func NewTrendingGetTrendingCategoryParams() *TrendingGetTrendingCategoryParams {
	var ()
	return &TrendingGetTrendingCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTrendingGetTrendingCategoryParamsWithTimeout creates a new TrendingGetTrendingCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrendingGetTrendingCategoryParamsWithTimeout(timeout time.Duration) *TrendingGetTrendingCategoryParams {
	var ()
	return &TrendingGetTrendingCategoryParams{

		timeout: timeout,
	}
}

// NewTrendingGetTrendingCategoryParamsWithContext creates a new TrendingGetTrendingCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrendingGetTrendingCategoryParamsWithContext(ctx context.Context) *TrendingGetTrendingCategoryParams {
	var ()
	return &TrendingGetTrendingCategoryParams{

		Context: ctx,
	}
}

// NewTrendingGetTrendingCategoryParamsWithHTTPClient creates a new TrendingGetTrendingCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrendingGetTrendingCategoryParamsWithHTTPClient(client *http.Client) *TrendingGetTrendingCategoryParams {
	var ()
	return &TrendingGetTrendingCategoryParams{
		HTTPClient: client,
	}
}

/*TrendingGetTrendingCategoryParams contains all the parameters to send to the API endpoint
for the trending get trending category operation typically these are written to a http.Request
*/
type TrendingGetTrendingCategoryParams struct {

	/*CategoryID
	  The ID of the category for whom you want additional results.

	*/
	CategoryID string
	/*PageNumber
	  The page # of results to return.

	*/
	PageNumber int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) WithTimeout(timeout time.Duration) *TrendingGetTrendingCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) WithContext(ctx context.Context) *TrendingGetTrendingCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) WithHTTPClient(client *http.Client) *TrendingGetTrendingCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) WithCategoryID(categoryID string) *TrendingGetTrendingCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithPageNumber adds the pageNumber to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) WithPageNumber(pageNumber int32) *TrendingGetTrendingCategoryParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the trending get trending category params
func (o *TrendingGetTrendingCategoryParams) SetPageNumber(pageNumber int32) {
	o.PageNumber = pageNumber
}

// WriteToRequest writes these params to a swagger request
func (o *TrendingGetTrendingCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
	}

	// path param pageNumber
	if err := r.SetPathParam("pageNumber", swag.FormatInt32(o.PageNumber)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
