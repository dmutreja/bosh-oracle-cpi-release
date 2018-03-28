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

// NewListSecurityListsParams creates a new ListSecurityListsParams object
// with the default values initialized.
func NewListSecurityListsParams() *ListSecurityListsParams {
	var ()
	return &ListSecurityListsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSecurityListsParamsWithTimeout creates a new ListSecurityListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSecurityListsParamsWithTimeout(timeout time.Duration) *ListSecurityListsParams {
	var ()
	return &ListSecurityListsParams{

		timeout: timeout,
	}
}

// NewListSecurityListsParamsWithContext creates a new ListSecurityListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSecurityListsParamsWithContext(ctx context.Context) *ListSecurityListsParams {
	var ()
	return &ListSecurityListsParams{

		Context: ctx,
	}
}

// NewListSecurityListsParamsWithHTTPClient creates a new ListSecurityListsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSecurityListsParamsWithHTTPClient(client *http.Client) *ListSecurityListsParams {
	var ()
	return &ListSecurityListsParams{
		HTTPClient: client,
	}
}

/*ListSecurityListsParams contains all the parameters to send to the API endpoint
for the list security lists operation typically these are written to a http.Request
*/
type ListSecurityListsParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*DisplayName
	  A filter to return only resources that match the given display name exactly.


	*/
	DisplayName *string
	/*LifecycleState
	  A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.


	*/
	LifecycleState *string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string
	/*SortBy
	  The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	sort order is case sensitive.

	**Note:** In general, some "List" operations (for example, `ListInstances`) let you
	optionally filter by Availability Domain if the scope of the resource type is within a
	single Availability Domain. If you call one of these "List" operations without specifying
	an Availability Domain, the resources are grouped by Availability Domain, then sorted.


	*/
	SortBy *string
	/*SortOrder
	  The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	is case sensitive.


	*/
	SortOrder *string
	/*VcnID
	  The OCID of the VCN.

	*/
	VcnID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list security lists params
func (o *ListSecurityListsParams) WithTimeout(timeout time.Duration) *ListSecurityListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list security lists params
func (o *ListSecurityListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list security lists params
func (o *ListSecurityListsParams) WithContext(ctx context.Context) *ListSecurityListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list security lists params
func (o *ListSecurityListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list security lists params
func (o *ListSecurityListsParams) WithHTTPClient(client *http.Client) *ListSecurityListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list security lists params
func (o *ListSecurityListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list security lists params
func (o *ListSecurityListsParams) WithCompartmentID(compartmentID string) *ListSecurityListsParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list security lists params
func (o *ListSecurityListsParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithDisplayName adds the displayName to the list security lists params
func (o *ListSecurityListsParams) WithDisplayName(displayName *string) *ListSecurityListsParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the list security lists params
func (o *ListSecurityListsParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithLifecycleState adds the lifecycleState to the list security lists params
func (o *ListSecurityListsParams) WithLifecycleState(lifecycleState *string) *ListSecurityListsParams {
	o.SetLifecycleState(lifecycleState)
	return o
}

// SetLifecycleState adds the lifecycleState to the list security lists params
func (o *ListSecurityListsParams) SetLifecycleState(lifecycleState *string) {
	o.LifecycleState = lifecycleState
}

// WithLimit adds the limit to the list security lists params
func (o *ListSecurityListsParams) WithLimit(limit *int64) *ListSecurityListsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list security lists params
func (o *ListSecurityListsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list security lists params
func (o *ListSecurityListsParams) WithPage(page *string) *ListSecurityListsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list security lists params
func (o *ListSecurityListsParams) SetPage(page *string) {
	o.Page = page
}

// WithSortBy adds the sortBy to the list security lists params
func (o *ListSecurityListsParams) WithSortBy(sortBy *string) *ListSecurityListsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list security lists params
func (o *ListSecurityListsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the list security lists params
func (o *ListSecurityListsParams) WithSortOrder(sortOrder *string) *ListSecurityListsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list security lists params
func (o *ListSecurityListsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithVcnID adds the vcnID to the list security lists params
func (o *ListSecurityListsParams) WithVcnID(vcnID string) *ListSecurityListsParams {
	o.SetVcnID(vcnID)
	return o
}

// SetVcnID adds the vcnId to the list security lists params
func (o *ListSecurityListsParams) SetVcnID(vcnID string) {
	o.VcnID = vcnID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSecurityListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if o.LifecycleState != nil {

		// query param lifecycleState
		var qrLifecycleState string
		if o.LifecycleState != nil {
			qrLifecycleState = *o.LifecycleState
		}
		qLifecycleState := qrLifecycleState
		if qLifecycleState != "" {
			if err := r.SetQueryParam("lifecycleState", qLifecycleState); err != nil {
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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	// query param vcnId
	qrVcnID := o.VcnID
	qVcnID := qrVcnID
	if qVcnID != "" {
		if err := r.SetQueryParam("vcnId", qVcnID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
