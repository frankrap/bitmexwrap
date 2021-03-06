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

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserCancelWithdrawalParams creates a new UserCancelWithdrawalParams object
// with the default values initialized.
func NewUserCancelWithdrawalParams() *UserCancelWithdrawalParams {
	var ()
	return &UserCancelWithdrawalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCancelWithdrawalParamsWithTimeout creates a new UserCancelWithdrawalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCancelWithdrawalParamsWithTimeout(timeout time.Duration) *UserCancelWithdrawalParams {
	var ()
	return &UserCancelWithdrawalParams{

		timeout: timeout,
	}
}

// NewUserCancelWithdrawalParamsWithContext creates a new UserCancelWithdrawalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCancelWithdrawalParamsWithContext(ctx context.Context) *UserCancelWithdrawalParams {
	var ()
	return &UserCancelWithdrawalParams{

		Context: ctx,
	}
}

// NewUserCancelWithdrawalParamsWithHTTPClient creates a new UserCancelWithdrawalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCancelWithdrawalParamsWithHTTPClient(client *http.Client) *UserCancelWithdrawalParams {
	var ()
	return &UserCancelWithdrawalParams{
		HTTPClient: client,
	}
}

/*UserCancelWithdrawalParams contains all the parameters to send to the API endpoint
for the user cancel withdrawal operation typically these are written to a http.Request
*/
type UserCancelWithdrawalParams struct {

	/*Token*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) WithTimeout(timeout time.Duration) *UserCancelWithdrawalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) WithContext(ctx context.Context) *UserCancelWithdrawalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) WithHTTPClient(client *http.Client) *UserCancelWithdrawalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithToken adds the token to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) WithToken(token string) *UserCancelWithdrawalParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the user cancel withdrawal params
func (o *UserCancelWithdrawalParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *UserCancelWithdrawalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
