// Code generated by go-swagger; DO NOT EDIT.

package number_search

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

// NewGetNumbersearchParams creates a new GetNumbersearchParams object
// with the default values initialized.
func NewGetNumbersearchParams() *GetNumbersearchParams {
	var (
		aPIVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetNumbersearchParams{
		APIVersion: &aPIVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNumbersearchParamsWithTimeout creates a new GetNumbersearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNumbersearchParamsWithTimeout(timeout time.Duration) *GetNumbersearchParams {
	var (
		aPIVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetNumbersearchParams{
		APIVersion: &aPIVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		timeout: timeout,
	}
}

// NewGetNumbersearchParamsWithContext creates a new GetNumbersearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNumbersearchParamsWithContext(ctx context.Context) *GetNumbersearchParams {
	var (
		apiVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetNumbersearchParams{
		APIVersion: &apiVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,

		Context: ctx,
	}
}

// NewGetNumbersearchParamsWithHTTPClient creates a new GetNumbersearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNumbersearchParamsWithHTTPClient(client *http.Client) *GetNumbersearchParams {
	var (
		apiVersionDefault = string("1.0")
		countryDefault    = string("US")
		limitDefault      = int64(10)
		offsetDefault     = int64(0)
	)
	return &GetNumbersearchParams{
		APIVersion: &apiVersionDefault,
		Country:    &countryDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetNumbersearchParams contains all the parameters to send to the API endpoint
for the get numbersearch operation typically these are written to a http.Request
*/
type GetNumbersearchParams struct {

	/*APIVersion
	  API Version. If not specified your pinned verison is used.

	*/
	APIVersion *string
	/*Contains
	  Filter by numbers which contain this value

	*/
	Contains *string
	/*Country
	  Filter by country ISO. Only one country can be filtered at a time.
	If no country filter is provided then results for United States are returned by default.


	*/
	Country *string
	/*Limit
	  The numbers of items to return.

	*/
	Limit *int64
	/*NumberType
	  Filter by number type: fixed, mobile, tollfree

	*/
	NumberType []string
	/*Offset
	  The number of items to skip before starting to collect the result set.

	*/
	Offset *int64
	/*Prefix
	  Filter by numbers with this prefix after country code

	*/
	Prefix *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get numbersearch params
func (o *GetNumbersearchParams) WithTimeout(timeout time.Duration) *GetNumbersearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get numbersearch params
func (o *GetNumbersearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get numbersearch params
func (o *GetNumbersearchParams) WithContext(ctx context.Context) *GetNumbersearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get numbersearch params
func (o *GetNumbersearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get numbersearch params
func (o *GetNumbersearchParams) WithHTTPClient(client *http.Client) *GetNumbersearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get numbersearch params
func (o *GetNumbersearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get numbersearch params
func (o *GetNumbersearchParams) WithAPIVersion(aPIVersion *string) *GetNumbersearchParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get numbersearch params
func (o *GetNumbersearchParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithContains adds the contains to the get numbersearch params
func (o *GetNumbersearchParams) WithContains(contains *string) *GetNumbersearchParams {
	o.SetContains(contains)
	return o
}

// SetContains adds the contains to the get numbersearch params
func (o *GetNumbersearchParams) SetContains(contains *string) {
	o.Contains = contains
}

// WithCountry adds the country to the get numbersearch params
func (o *GetNumbersearchParams) WithCountry(country *string) *GetNumbersearchParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the get numbersearch params
func (o *GetNumbersearchParams) SetCountry(country *string) {
	o.Country = country
}

// WithLimit adds the limit to the get numbersearch params
func (o *GetNumbersearchParams) WithLimit(limit *int64) *GetNumbersearchParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get numbersearch params
func (o *GetNumbersearchParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNumberType adds the numberType to the get numbersearch params
func (o *GetNumbersearchParams) WithNumberType(numberType []string) *GetNumbersearchParams {
	o.SetNumberType(numberType)
	return o
}

// SetNumberType adds the numberType to the get numbersearch params
func (o *GetNumbersearchParams) SetNumberType(numberType []string) {
	o.NumberType = numberType
}

// WithOffset adds the offset to the get numbersearch params
func (o *GetNumbersearchParams) WithOffset(offset *int64) *GetNumbersearchParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get numbersearch params
func (o *GetNumbersearchParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPrefix adds the prefix to the get numbersearch params
func (o *GetNumbersearchParams) WithPrefix(prefix *string) *GetNumbersearchParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the get numbersearch params
func (o *GetNumbersearchParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetNumbersearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Contains != nil {

		// query param contains
		var qrContains string
		if o.Contains != nil {
			qrContains = *o.Contains
		}
		qContains := qrContains
		if qContains != "" {
			if err := r.SetQueryParam("contains", qContains); err != nil {
				return err
			}
		}

	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesNumberType := o.NumberType

	joinedNumberType := swag.JoinByFormat(valuesNumberType, "csv")
	// query array param number_type
	if err := r.SetQueryParam("number_type", joinedNumberType...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string
		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {
			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
