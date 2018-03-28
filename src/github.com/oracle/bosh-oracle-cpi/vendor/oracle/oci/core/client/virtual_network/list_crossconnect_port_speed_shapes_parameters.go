// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

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

// NewListCrossconnectPortSpeedShapesParams creates a new ListCrossconnectPortSpeedShapesParams object
// with the default values initialized.
func NewListCrossconnectPortSpeedShapesParams() *ListCrossconnectPortSpeedShapesParams {
	var ()
	return &ListCrossconnectPortSpeedShapesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCrossconnectPortSpeedShapesParamsWithTimeout creates a new ListCrossconnectPortSpeedShapesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCrossconnectPortSpeedShapesParamsWithTimeout(timeout time.Duration) *ListCrossconnectPortSpeedShapesParams {
	var ()
	return &ListCrossconnectPortSpeedShapesParams{

		timeout: timeout,
	}
}

// NewListCrossconnectPortSpeedShapesParamsWithContext creates a new ListCrossconnectPortSpeedShapesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCrossconnectPortSpeedShapesParamsWithContext(ctx context.Context) *ListCrossconnectPortSpeedShapesParams {
	var ()
	return &ListCrossconnectPortSpeedShapesParams{

		Context: ctx,
	}
}

// NewListCrossconnectPortSpeedShapesParamsWithHTTPClient creates a new ListCrossconnectPortSpeedShapesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCrossconnectPortSpeedShapesParamsWithHTTPClient(client *http.Client) *ListCrossconnectPortSpeedShapesParams {
	var ()
	return &ListCrossconnectPortSpeedShapesParams{
		HTTPClient: client,
	}
}

/*ListCrossconnectPortSpeedShapesParams contains all the parameters to send to the API endpoint
for the list crossconnect port speed shapes operation typically these are written to a http.Request
*/
type ListCrossconnectPortSpeedShapesParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithTimeout(timeout time.Duration) *ListCrossconnectPortSpeedShapesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithContext(ctx context.Context) *ListCrossconnectPortSpeedShapesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithHTTPClient(client *http.Client) *ListCrossconnectPortSpeedShapesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithCompartmentID(compartmentID string) *ListCrossconnectPortSpeedShapesParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithLimit adds the limit to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithLimit(limit *int64) *ListCrossconnectPortSpeedShapesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) WithPage(page *string) *ListCrossconnectPortSpeedShapesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list crossconnect port speed shapes params
func (o *ListCrossconnectPortSpeedShapesParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListCrossconnectPortSpeedShapesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
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

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
