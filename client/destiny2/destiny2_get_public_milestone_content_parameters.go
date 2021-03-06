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

// NewDestiny2GetPublicMilestoneContentParams creates a new Destiny2GetPublicMilestoneContentParams object
// with the default values initialized.
func NewDestiny2GetPublicMilestoneContentParams() *Destiny2GetPublicMilestoneContentParams {
	var ()
	return &Destiny2GetPublicMilestoneContentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDestiny2GetPublicMilestoneContentParamsWithTimeout creates a new Destiny2GetPublicMilestoneContentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDestiny2GetPublicMilestoneContentParamsWithTimeout(timeout time.Duration) *Destiny2GetPublicMilestoneContentParams {
	var ()
	return &Destiny2GetPublicMilestoneContentParams{

		timeout: timeout,
	}
}

// NewDestiny2GetPublicMilestoneContentParamsWithContext creates a new Destiny2GetPublicMilestoneContentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDestiny2GetPublicMilestoneContentParamsWithContext(ctx context.Context) *Destiny2GetPublicMilestoneContentParams {
	var ()
	return &Destiny2GetPublicMilestoneContentParams{

		Context: ctx,
	}
}

// NewDestiny2GetPublicMilestoneContentParamsWithHTTPClient creates a new Destiny2GetPublicMilestoneContentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDestiny2GetPublicMilestoneContentParamsWithHTTPClient(client *http.Client) *Destiny2GetPublicMilestoneContentParams {
	var ()
	return &Destiny2GetPublicMilestoneContentParams{
		HTTPClient: client,
	}
}

/*Destiny2GetPublicMilestoneContentParams contains all the parameters to send to the API endpoint
for the destiny2 get public milestone content operation typically these are written to a http.Request
*/
type Destiny2GetPublicMilestoneContentParams struct {

	/*MilestoneHash
	  The identifier for the milestone to be returned.

	*/
	MilestoneHash uint32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) WithTimeout(timeout time.Duration) *Destiny2GetPublicMilestoneContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) WithContext(ctx context.Context) *Destiny2GetPublicMilestoneContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) WithHTTPClient(client *http.Client) *Destiny2GetPublicMilestoneContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestoneHash adds the milestoneHash to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) WithMilestoneHash(milestoneHash uint32) *Destiny2GetPublicMilestoneContentParams {
	o.SetMilestoneHash(milestoneHash)
	return o
}

// SetMilestoneHash adds the milestoneHash to the destiny2 get public milestone content params
func (o *Destiny2GetPublicMilestoneContentParams) SetMilestoneHash(milestoneHash uint32) {
	o.MilestoneHash = milestoneHash
}

// WriteToRequest writes these params to a swagger request
func (o *Destiny2GetPublicMilestoneContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestoneHash
	if err := r.SetPathParam("milestoneHash", swag.FormatUint32(o.MilestoneHash)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
