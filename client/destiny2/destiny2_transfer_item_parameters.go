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

// NewDestiny2TransferItemParams creates a new Destiny2TransferItemParams object
// with the default values initialized.
func NewDestiny2TransferItemParams() *Destiny2TransferItemParams {

	return &Destiny2TransferItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2TransferItemParamsWithTimeout creates a new Destiny2TransferItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2TransferItemParamsWithTimeout(timeout time.Duration) *Destiny2TransferItemParams {

	return &Destiny2TransferItemParams{

		timeout: timeout,
	}
}

// NewDestiny2TransferItemParamsWithContext creates a new Destiny2TransferItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2TransferItemParamsWithContext(ctx context.Context) *Destiny2TransferItemParams {

	return &Destiny2TransferItemParams{

		Context: ctx,
	}
}

// NewDestiny2TransferItemParamsWithHTTPClient creates a new Destiny2TransferItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2TransferItemParamsWithHTTPClient(client *http.Client) *Destiny2TransferItemParams {

	return &Destiny2TransferItemParams{
		HTTPClient: client,
	}
}

/*Destiny2TransferItemParams contains all the parameters to send to the API endpoint
for the destiny2 transfer item operation typically these are written to a http.Request
*/
type Destiny2TransferItemParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) WithTimeout(timeout time.Duration) *Destiny2TransferItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) WithContext(ctx context.Context) *Destiny2TransferItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) WithHTTPClient(client *http.Client) *Destiny2TransferItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 transfer item params
func (o *Destiny2TransferItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2TransferItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
