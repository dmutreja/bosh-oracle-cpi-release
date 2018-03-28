// Code generated by go-swagger; DO NOT EDIT.

package blockstorage

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

// NewDeleteVolumeBackupParams creates a new DeleteVolumeBackupParams object
// with the default values initialized.
func NewDeleteVolumeBackupParams() *DeleteVolumeBackupParams {
	var ()
	return &DeleteVolumeBackupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVolumeBackupParamsWithTimeout creates a new DeleteVolumeBackupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVolumeBackupParamsWithTimeout(timeout time.Duration) *DeleteVolumeBackupParams {
	var ()
	return &DeleteVolumeBackupParams{

		timeout: timeout,
	}
}

// NewDeleteVolumeBackupParamsWithContext creates a new DeleteVolumeBackupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVolumeBackupParamsWithContext(ctx context.Context) *DeleteVolumeBackupParams {
	var ()
	return &DeleteVolumeBackupParams{

		Context: ctx,
	}
}

// NewDeleteVolumeBackupParamsWithHTTPClient creates a new DeleteVolumeBackupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVolumeBackupParamsWithHTTPClient(client *http.Client) *DeleteVolumeBackupParams {
	var ()
	return &DeleteVolumeBackupParams{
		HTTPClient: client,
	}
}

/*DeleteVolumeBackupParams contains all the parameters to send to the API endpoint
for the delete volume backup operation typically these are written to a http.Request
*/
type DeleteVolumeBackupParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*VolumeBackupID
	  The OCID of the volume backup.

	*/
	VolumeBackupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete volume backup params
func (o *DeleteVolumeBackupParams) WithTimeout(timeout time.Duration) *DeleteVolumeBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete volume backup params
func (o *DeleteVolumeBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete volume backup params
func (o *DeleteVolumeBackupParams) WithContext(ctx context.Context) *DeleteVolumeBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete volume backup params
func (o *DeleteVolumeBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete volume backup params
func (o *DeleteVolumeBackupParams) WithHTTPClient(client *http.Client) *DeleteVolumeBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete volume backup params
func (o *DeleteVolumeBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete volume backup params
func (o *DeleteVolumeBackupParams) WithIfMatch(ifMatch *string) *DeleteVolumeBackupParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete volume backup params
func (o *DeleteVolumeBackupParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithVolumeBackupID adds the volumeBackupID to the delete volume backup params
func (o *DeleteVolumeBackupParams) WithVolumeBackupID(volumeBackupID string) *DeleteVolumeBackupParams {
	o.SetVolumeBackupID(volumeBackupID)
	return o
}

// SetVolumeBackupID adds the volumeBackupId to the delete volume backup params
func (o *DeleteVolumeBackupParams) SetVolumeBackupID(volumeBackupID string) {
	o.VolumeBackupID = volumeBackupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVolumeBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param volumeBackupId
	if err := r.SetPathParam("volumeBackupId", o.VolumeBackupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
