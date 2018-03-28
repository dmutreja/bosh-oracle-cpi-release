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

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// NewUpdateInternetGatewayParams creates a new UpdateInternetGatewayParams object
// with the default values initialized.
func NewUpdateInternetGatewayParams() *UpdateInternetGatewayParams {
	var ()
	return &UpdateInternetGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInternetGatewayParamsWithTimeout creates a new UpdateInternetGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInternetGatewayParamsWithTimeout(timeout time.Duration) *UpdateInternetGatewayParams {
	var ()
	return &UpdateInternetGatewayParams{

		timeout: timeout,
	}
}

// NewUpdateInternetGatewayParamsWithContext creates a new UpdateInternetGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInternetGatewayParamsWithContext(ctx context.Context) *UpdateInternetGatewayParams {
	var ()
	return &UpdateInternetGatewayParams{

		Context: ctx,
	}
}

// NewUpdateInternetGatewayParamsWithHTTPClient creates a new UpdateInternetGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInternetGatewayParamsWithHTTPClient(client *http.Client) *UpdateInternetGatewayParams {
	var ()
	return &UpdateInternetGatewayParams{
		HTTPClient: client,
	}
}

/*UpdateInternetGatewayParams contains all the parameters to send to the API endpoint
for the update internet gateway operation typically these are written to a http.Request
*/
type UpdateInternetGatewayParams struct {

	/*UpdateInternetGatewayDetails
	  Details for updating the Internet Gateway.

	*/
	UpdateInternetGatewayDetails *models.UpdateInternetGatewayDetails
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*IgID
	  The OCID of the Internet Gateway.

	*/
	IgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithTimeout(timeout time.Duration) *UpdateInternetGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithContext(ctx context.Context) *UpdateInternetGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithHTTPClient(client *http.Client) *UpdateInternetGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateInternetGatewayDetails adds the updateInternetGatewayDetails to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithUpdateInternetGatewayDetails(updateInternetGatewayDetails *models.UpdateInternetGatewayDetails) *UpdateInternetGatewayParams {
	o.SetUpdateInternetGatewayDetails(updateInternetGatewayDetails)
	return o
}

// SetUpdateInternetGatewayDetails adds the updateInternetGatewayDetails to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetUpdateInternetGatewayDetails(updateInternetGatewayDetails *models.UpdateInternetGatewayDetails) {
	o.UpdateInternetGatewayDetails = updateInternetGatewayDetails
}

// WithIfMatch adds the ifMatch to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithIfMatch(ifMatch *string) *UpdateInternetGatewayParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithIgID adds the igID to the update internet gateway params
func (o *UpdateInternetGatewayParams) WithIgID(igID string) *UpdateInternetGatewayParams {
	o.SetIgID(igID)
	return o
}

// SetIgID adds the igId to the update internet gateway params
func (o *UpdateInternetGatewayParams) SetIgID(igID string) {
	o.IgID = igID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInternetGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateInternetGatewayDetails != nil {
		if err := r.SetBodyParam(o.UpdateInternetGatewayDetails); err != nil {
			return err
		}
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param igId
	if err := r.SetPathParam("igId", o.IgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
