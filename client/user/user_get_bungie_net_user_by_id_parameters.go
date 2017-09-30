// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUserGetBungieNetUserByIDParams creates a new UserGetBungieNetUserByIDParams object
// with the default values initialized.
func NewUserGetBungieNetUserByIDParams() *UserGetBungieNetUserByIDParams {
	var ()
	return &UserGetBungieNetUserByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetBungieNetUserByIDParamsWithTimeout creates a new UserGetBungieNetUserByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetBungieNetUserByIDParamsWithTimeout(timeout time.Duration) *UserGetBungieNetUserByIDParams {
	var ()
	return &UserGetBungieNetUserByIDParams{

		timeout: timeout,
	}
}

// NewUserGetBungieNetUserByIDParamsWithContext creates a new UserGetBungieNetUserByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetBungieNetUserByIDParamsWithContext(ctx context.Context) *UserGetBungieNetUserByIDParams {
	var ()
	return &UserGetBungieNetUserByIDParams{

		Context: ctx,
	}
}

// NewUserGetBungieNetUserByIDParamsWithHTTPClient creates a new UserGetBungieNetUserByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetBungieNetUserByIDParamsWithHTTPClient(client *http.Client) *UserGetBungieNetUserByIDParams {
	var ()
	return &UserGetBungieNetUserByIDParams{
		HTTPClient: client,
	}
}

/*UserGetBungieNetUserByIDParams contains all the parameters to send to the API endpoint
for the user get bungie net user by Id operation typically these are written to a http.Request
*/
type UserGetBungieNetUserByIDParams struct {

	/*ID
	  The requested Bungie.net membership id.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) WithTimeout(timeout time.Duration) *UserGetBungieNetUserByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) WithContext(ctx context.Context) *UserGetBungieNetUserByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) WithHTTPClient(client *http.Client) *UserGetBungieNetUserByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) WithID(id int64) *UserGetBungieNetUserByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user get bungie net user by Id params
func (o *UserGetBungieNetUserByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetBungieNetUserByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
