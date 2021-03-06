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

// NewUserGetPartnershipsParams creates a new UserGetPartnershipsParams object
// with the default values initialized.
func NewUserGetPartnershipsParams() *UserGetPartnershipsParams {
	var ()
	return &UserGetPartnershipsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserGetPartnershipsParamsWithTimeout creates a new UserGetPartnershipsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserGetPartnershipsParamsWithTimeout(timeout time.Duration) *UserGetPartnershipsParams {
	var ()
	return &UserGetPartnershipsParams{

		timeout: timeout,
	}
}

// NewUserGetPartnershipsParamsWithContext creates a new UserGetPartnershipsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserGetPartnershipsParamsWithContext(ctx context.Context) *UserGetPartnershipsParams {
	var ()
	return &UserGetPartnershipsParams{

		Context: ctx,
	}
}

// NewUserGetPartnershipsParamsWithHTTPClient creates a new UserGetPartnershipsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserGetPartnershipsParamsWithHTTPClient(client *http.Client) *UserGetPartnershipsParams {
	var ()
	return &UserGetPartnershipsParams{
		HTTPClient: client,
	}
}

/*UserGetPartnershipsParams contains all the parameters to send to the API endpoint
for the user get partnerships operation typically these are written to a http.Request
*/
type UserGetPartnershipsParams struct {

	/*MembershipID
	  The ID of the member for whom partnerships should be returned.

	*/
	MembershipID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user get partnerships params
func (o *UserGetPartnershipsParams) WithTimeout(timeout time.Duration) *UserGetPartnershipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user get partnerships params
func (o *UserGetPartnershipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user get partnerships params
func (o *UserGetPartnershipsParams) WithContext(ctx context.Context) *UserGetPartnershipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user get partnerships params
func (o *UserGetPartnershipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user get partnerships params
func (o *UserGetPartnershipsParams) WithHTTPClient(client *http.Client) *UserGetPartnershipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user get partnerships params
func (o *UserGetPartnershipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMembershipID adds the membershipID to the user get partnerships params
func (o *UserGetPartnershipsParams) WithMembershipID(membershipID int64) *UserGetPartnershipsParams {
	o.SetMembershipID(membershipID)
	return o
}

// SetMembershipID adds the membershipId to the user get partnerships params
func (o *UserGetPartnershipsParams) SetMembershipID(membershipID int64) {
	o.MembershipID = membershipID
}

// WriteToRequest writes these params to a swagger request
func (o *UserGetPartnershipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param membershipId
	if err := r.SetPathParam("membershipId", swag.FormatInt64(o.MembershipID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
