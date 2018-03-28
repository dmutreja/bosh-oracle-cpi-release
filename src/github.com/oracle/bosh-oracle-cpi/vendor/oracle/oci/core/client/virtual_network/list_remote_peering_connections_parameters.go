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

// NewListRemotePeeringConnectionsParams creates a new ListRemotePeeringConnectionsParams object
// with the default values initialized.
func NewListRemotePeeringConnectionsParams() *ListRemotePeeringConnectionsParams {
	var ()
	return &ListRemotePeeringConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRemotePeeringConnectionsParamsWithTimeout creates a new ListRemotePeeringConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRemotePeeringConnectionsParamsWithTimeout(timeout time.Duration) *ListRemotePeeringConnectionsParams {
	var ()
	return &ListRemotePeeringConnectionsParams{

		timeout: timeout,
	}
}

// NewListRemotePeeringConnectionsParamsWithContext creates a new ListRemotePeeringConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRemotePeeringConnectionsParamsWithContext(ctx context.Context) *ListRemotePeeringConnectionsParams {
	var ()
	return &ListRemotePeeringConnectionsParams{

		Context: ctx,
	}
}

// NewListRemotePeeringConnectionsParamsWithHTTPClient creates a new ListRemotePeeringConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRemotePeeringConnectionsParamsWithHTTPClient(client *http.Client) *ListRemotePeeringConnectionsParams {
	var ()
	return &ListRemotePeeringConnectionsParams{
		HTTPClient: client,
	}
}

/*ListRemotePeeringConnectionsParams contains all the parameters to send to the API endpoint
for the list remote peering connections operation typically these are written to a http.Request
*/
type ListRemotePeeringConnectionsParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*DrgID
	  The OCID of the DRG.

	*/
	DrgID *string
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

// WithTimeout adds the timeout to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithTimeout(timeout time.Duration) *ListRemotePeeringConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithContext(ctx context.Context) *ListRemotePeeringConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithHTTPClient(client *http.Client) *ListRemotePeeringConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithCompartmentID(compartmentID string) *ListRemotePeeringConnectionsParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithDrgID adds the drgID to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithDrgID(drgID *string) *ListRemotePeeringConnectionsParams {
	o.SetDrgID(drgID)
	return o
}

// SetDrgID adds the drgId to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetDrgID(drgID *string) {
	o.DrgID = drgID
}

// WithLimit adds the limit to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithLimit(limit *int64) *ListRemotePeeringConnectionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) WithPage(page *string) *ListRemotePeeringConnectionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list remote peering connections params
func (o *ListRemotePeeringConnectionsParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListRemotePeeringConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DrgID != nil {

		// query param drgId
		var qrDrgID string
		if o.DrgID != nil {
			qrDrgID = *o.DrgID
		}
		qDrgID := qrDrgID
		if qDrgID != "" {
			if err := r.SetQueryParam("drgId", qDrgID); err != nil {
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
