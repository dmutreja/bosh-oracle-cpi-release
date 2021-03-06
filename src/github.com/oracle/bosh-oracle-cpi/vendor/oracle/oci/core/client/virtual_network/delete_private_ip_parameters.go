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
)

// NewDeletePrivateIPParams creates a new DeletePrivateIPParams object
// with the default values initialized.
func NewDeletePrivateIPParams() *DeletePrivateIPParams {
	var ()
	return &DeletePrivateIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrivateIPParamsWithTimeout creates a new DeletePrivateIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePrivateIPParamsWithTimeout(timeout time.Duration) *DeletePrivateIPParams {
	var ()
	return &DeletePrivateIPParams{

		timeout: timeout,
	}
}

// NewDeletePrivateIPParamsWithContext creates a new DeletePrivateIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePrivateIPParamsWithContext(ctx context.Context) *DeletePrivateIPParams {
	var ()
	return &DeletePrivateIPParams{

		Context: ctx,
	}
}

// NewDeletePrivateIPParamsWithHTTPClient creates a new DeletePrivateIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePrivateIPParamsWithHTTPClient(client *http.Client) *DeletePrivateIPParams {
	var ()
	return &DeletePrivateIPParams{
		HTTPClient: client,
	}
}

/*DeletePrivateIPParams contains all the parameters to send to the API endpoint
for the delete private Ip operation typically these are written to a http.Request
*/
type DeletePrivateIPParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*PrivateIPID
	  The OCID of the private IP.

	*/
	PrivateIPID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete private Ip params
func (o *DeletePrivateIPParams) WithTimeout(timeout time.Duration) *DeletePrivateIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete private Ip params
func (o *DeletePrivateIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete private Ip params
func (o *DeletePrivateIPParams) WithContext(ctx context.Context) *DeletePrivateIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete private Ip params
func (o *DeletePrivateIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete private Ip params
func (o *DeletePrivateIPParams) WithHTTPClient(client *http.Client) *DeletePrivateIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete private Ip params
func (o *DeletePrivateIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete private Ip params
func (o *DeletePrivateIPParams) WithIfMatch(ifMatch *string) *DeletePrivateIPParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete private Ip params
func (o *DeletePrivateIPParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithPrivateIPID adds the privateIPID to the delete private Ip params
func (o *DeletePrivateIPParams) WithPrivateIPID(privateIPID string) *DeletePrivateIPParams {
	o.SetPrivateIPID(privateIPID)
	return o
}

// SetPrivateIPID adds the privateIpId to the delete private Ip params
func (o *DeletePrivateIPParams) SetPrivateIPID(privateIPID string) {
	o.PrivateIPID = privateIPID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrivateIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param privateIpId
	if err := r.SetPathParam("privateIpId", o.PrivateIPID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
