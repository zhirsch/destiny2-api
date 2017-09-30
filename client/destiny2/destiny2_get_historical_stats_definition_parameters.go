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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDestiny2GetHistoricalStatsDefinitionParams creates a new Destiny2GetHistoricalStatsDefinitionParams object
// with the default values initialized.
func NewDestiny2GetHistoricalStatsDefinitionParams() *Destiny2GetHistoricalStatsDefinitionParams {

	return &Destiny2GetHistoricalStatsDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetHistoricalStatsDefinitionParamsWithTimeout creates a new Destiny2GetHistoricalStatsDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetHistoricalStatsDefinitionParamsWithTimeout(timeout time.Duration) *Destiny2GetHistoricalStatsDefinitionParams {

	return &Destiny2GetHistoricalStatsDefinitionParams{

		timeout: timeout,
	}
}

// NewDestiny2GetHistoricalStatsDefinitionParamsWithContext creates a new Destiny2GetHistoricalStatsDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetHistoricalStatsDefinitionParamsWithContext(ctx context.Context) *Destiny2GetHistoricalStatsDefinitionParams {

	return &Destiny2GetHistoricalStatsDefinitionParams{

		Context: ctx,
	}
}

// NewDestiny2GetHistoricalStatsDefinitionParamsWithHTTPClient creates a new Destiny2GetHistoricalStatsDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetHistoricalStatsDefinitionParamsWithHTTPClient(client *http.Client) *Destiny2GetHistoricalStatsDefinitionParams {

	return &Destiny2GetHistoricalStatsDefinitionParams{
		HTTPClient: client,
	}
}

/*Destiny2GetHistoricalStatsDefinitionParams contains all the parameters to send to the API endpoint
for the destiny2 get historical stats definition operation typically these are written to a http.Request
*/
type Destiny2GetHistoricalStatsDefinitionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) WithTimeout(timeout time.Duration) *Destiny2GetHistoricalStatsDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) WithContext(ctx context.Context) *Destiny2GetHistoricalStatsDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) WithHTTPClient(client *http.Client) *Destiny2GetHistoricalStatsDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get historical stats definition params
func (o *Destiny2GetHistoricalStatsDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetHistoricalStatsDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
