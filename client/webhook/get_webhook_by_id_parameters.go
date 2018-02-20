// Code generated by go-swagger; DO NOT EDIT.

package webhook

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

// NewGetWebhookByIDParams creates a new GetWebhookByIDParams object
// with the default values initialized.
func NewGetWebhookByIDParams() *GetWebhookByIDParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &GetWebhookByIDParams{
		APIVersion: &aPIVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebhookByIDParamsWithTimeout creates a new GetWebhookByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebhookByIDParamsWithTimeout(timeout time.Duration) *GetWebhookByIDParams {
	var (
		aPIVersionDefault = string("1.0")
	)
	return &GetWebhookByIDParams{
		APIVersion: &aPIVersionDefault,

		timeout: timeout,
	}
}

// NewGetWebhookByIDParamsWithContext creates a new GetWebhookByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebhookByIDParamsWithContext(ctx context.Context) *GetWebhookByIDParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &GetWebhookByIDParams{
		APIVersion: &apiVersionDefault,

		Context: ctx,
	}
}

// NewGetWebhookByIDParamsWithHTTPClient creates a new GetWebhookByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebhookByIDParamsWithHTTPClient(client *http.Client) *GetWebhookByIDParams {
	var (
		apiVersionDefault = string("1.0")
	)
	return &GetWebhookByIDParams{
		APIVersion: &apiVersionDefault,
		HTTPClient: client,
	}
}

/*GetWebhookByIDParams contains all the parameters to send to the API endpoint
for the get webhook by Id operation typically these are written to a http.Request
*/
type GetWebhookByIDParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion *string
	/*UID
	  Alphanumeric ID of the webhook to get.

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get webhook by Id params
func (o *GetWebhookByIDParams) WithTimeout(timeout time.Duration) *GetWebhookByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get webhook by Id params
func (o *GetWebhookByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get webhook by Id params
func (o *GetWebhookByIDParams) WithContext(ctx context.Context) *GetWebhookByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get webhook by Id params
func (o *GetWebhookByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get webhook by Id params
func (o *GetWebhookByIDParams) WithHTTPClient(client *http.Client) *GetWebhookByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get webhook by Id params
func (o *GetWebhookByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get webhook by Id params
func (o *GetWebhookByIDParams) WithAPIVersion(aPIVersion *string) *GetWebhookByIDParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get webhook by Id params
func (o *GetWebhookByIDParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithUID adds the uid to the get webhook by Id params
func (o *GetWebhookByIDParams) WithUID(uid string) *GetWebhookByIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get webhook by Id params
func (o *GetWebhookByIDParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebhookByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// header param api-version
		if err := r.SetHeaderParam("api-version", *o.APIVersion); err != nil {
			return err
		}

	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}