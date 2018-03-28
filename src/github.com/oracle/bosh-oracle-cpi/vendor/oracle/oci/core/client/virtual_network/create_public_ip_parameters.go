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

// NewCreatePublicIPParams creates a new CreatePublicIPParams object
// with the default values initialized.
func NewCreatePublicIPParams() *CreatePublicIPParams {
	var ()
	return &CreatePublicIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePublicIPParamsWithTimeout creates a new CreatePublicIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePublicIPParamsWithTimeout(timeout time.Duration) *CreatePublicIPParams {
	var ()
	return &CreatePublicIPParams{

		timeout: timeout,
	}
}

// NewCreatePublicIPParamsWithContext creates a new CreatePublicIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePublicIPParamsWithContext(ctx context.Context) *CreatePublicIPParams {
	var ()
	return &CreatePublicIPParams{

		Context: ctx,
	}
}

// NewCreatePublicIPParamsWithHTTPClient creates a new CreatePublicIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePublicIPParamsWithHTTPClient(client *http.Client) *CreatePublicIPParams {
	var ()
	return &CreatePublicIPParams{
		HTTPClient: client,
	}
}

/*CreatePublicIPParams contains all the parameters to send to the API endpoint
for the create public Ip operation typically these are written to a http.Request
*/
type CreatePublicIPParams struct {

	/*CreatePublicIPDetails
	  Create public IP details.

	*/
	CreatePublicIPDetails *models.CreatePublicIPDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create public Ip params
func (o *CreatePublicIPParams) WithTimeout(timeout time.Duration) *CreatePublicIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create public Ip params
func (o *CreatePublicIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create public Ip params
func (o *CreatePublicIPParams) WithContext(ctx context.Context) *CreatePublicIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create public Ip params
func (o *CreatePublicIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create public Ip params
func (o *CreatePublicIPParams) WithHTTPClient(client *http.Client) *CreatePublicIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create public Ip params
func (o *CreatePublicIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatePublicIPDetails adds the createPublicIPDetails to the create public Ip params
func (o *CreatePublicIPParams) WithCreatePublicIPDetails(createPublicIPDetails *models.CreatePublicIPDetails) *CreatePublicIPParams {
	o.SetCreatePublicIPDetails(createPublicIPDetails)
	return o
}

// SetCreatePublicIPDetails adds the createPublicIpDetails to the create public Ip params
func (o *CreatePublicIPParams) SetCreatePublicIPDetails(createPublicIPDetails *models.CreatePublicIPDetails) {
	o.CreatePublicIPDetails = createPublicIPDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create public Ip params
func (o *CreatePublicIPParams) WithOpcRetryToken(opcRetryToken *string) *CreatePublicIPParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create public Ip params
func (o *CreatePublicIPParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePublicIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatePublicIPDetails != nil {
		if err := r.SetBodyParam(o.CreatePublicIPDetails); err != nil {
			return err
		}
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}