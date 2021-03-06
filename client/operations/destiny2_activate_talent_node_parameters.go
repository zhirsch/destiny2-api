// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDestiny2ActivateTalentNodeParams creates a new Destiny2ActivateTalentNodeParams object
// with the default values initialized.
func NewDestiny2ActivateTalentNodeParams() *Destiny2ActivateTalentNodeParams {

	return &Destiny2ActivateTalentNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2ActivateTalentNodeParamsWithTimeout creates a new Destiny2ActivateTalentNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2ActivateTalentNodeParamsWithTimeout(timeout time.Duration) *Destiny2ActivateTalentNodeParams {

	return &Destiny2ActivateTalentNodeParams{

		timeout: timeout,
	}
}

// NewDestiny2ActivateTalentNodeParamsWithContext creates a new Destiny2ActivateTalentNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2ActivateTalentNodeParamsWithContext(ctx context.Context) *Destiny2ActivateTalentNodeParams {

	return &Destiny2ActivateTalentNodeParams{

		Context: ctx,
	}
}

// NewDestiny2ActivateTalentNodeParamsWithHTTPClient creates a new Destiny2ActivateTalentNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2ActivateTalentNodeParamsWithHTTPClient(client *http.Client) *Destiny2ActivateTalentNodeParams {

	return &Destiny2ActivateTalentNodeParams{
		HTTPClient: client,
	}
}

/*Destiny2ActivateTalentNodeParams contains all the parameters to send to the API endpoint
for the destiny2 activate talent node operation typically these are written to a http.Request
*/
type Destiny2ActivateTalentNodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) WithTimeout(timeout time.Duration) *Destiny2ActivateTalentNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) WithContext(ctx context.Context) *Destiny2ActivateTalentNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) WithHTTPClient(client *http.Client) *Destiny2ActivateTalentNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 activate talent node params
func (o *Destiny2ActivateTalentNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2ActivateTalentNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
