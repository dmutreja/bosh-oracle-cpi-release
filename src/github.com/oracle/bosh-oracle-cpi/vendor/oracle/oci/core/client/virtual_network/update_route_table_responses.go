// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// UpdateRouteTableReader is a Reader for the UpdateRouteTable structure.
type UpdateRouteTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRouteTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRouteTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateRouteTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateRouteTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateRouteTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateRouteTableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateRouteTablePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateRouteTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateRouteTableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRouteTableOK creates a UpdateRouteTableOK with default headers values
func NewUpdateRouteTableOK() *UpdateRouteTableOK {
	return &UpdateRouteTableOK{}
}

/*UpdateRouteTableOK handles this case with default header values.

The route table was updated.
*/
type UpdateRouteTableOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.RouteTable
}

func (o *UpdateRouteTableOK) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableOK  %+v", 200, o.Payload)
}

func (o *UpdateRouteTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.RouteTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableBadRequest creates a UpdateRouteTableBadRequest with default headers values
func NewUpdateRouteTableBadRequest() *UpdateRouteTableBadRequest {
	return &UpdateRouteTableBadRequest{}
}

/*UpdateRouteTableBadRequest handles this case with default header values.

Bad Request
*/
type UpdateRouteTableBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRouteTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableUnauthorized creates a UpdateRouteTableUnauthorized with default headers values
func NewUpdateRouteTableUnauthorized() *UpdateRouteTableUnauthorized {
	return &UpdateRouteTableUnauthorized{}
}

/*UpdateRouteTableUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateRouteTableUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTableUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRouteTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableNotFound creates a UpdateRouteTableNotFound with default headers values
func NewUpdateRouteTableNotFound() *UpdateRouteTableNotFound {
	return &UpdateRouteTableNotFound{}
}

/*UpdateRouteTableNotFound handles this case with default header values.

Not Found
*/
type UpdateRouteTableNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTableNotFound) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRouteTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableConflict creates a UpdateRouteTableConflict with default headers values
func NewUpdateRouteTableConflict() *UpdateRouteTableConflict {
	return &UpdateRouteTableConflict{}
}

/*UpdateRouteTableConflict handles this case with default header values.

Conflict
*/
type UpdateRouteTableConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTableConflict) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableConflict  %+v", 409, o.Payload)
}

func (o *UpdateRouteTableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTablePreconditionFailed creates a UpdateRouteTablePreconditionFailed with default headers values
func NewUpdateRouteTablePreconditionFailed() *UpdateRouteTablePreconditionFailed {
	return &UpdateRouteTablePreconditionFailed{}
}

/*UpdateRouteTablePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateRouteTablePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTablePreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTablePreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateRouteTablePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableInternalServerError creates a UpdateRouteTableInternalServerError with default headers values
func NewUpdateRouteTableInternalServerError() *UpdateRouteTableInternalServerError {
	return &UpdateRouteTableInternalServerError{}
}

/*UpdateRouteTableInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateRouteTableInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateRouteTableInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] updateRouteTableInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRouteTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRouteTableDefault creates a UpdateRouteTableDefault with default headers values
func NewUpdateRouteTableDefault(code int) *UpdateRouteTableDefault {
	return &UpdateRouteTableDefault{
		_statusCode: code,
	}
}

/*UpdateRouteTableDefault handles this case with default header values.

An error has occurred.
*/
type UpdateRouteTableDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update route table default response
func (o *UpdateRouteTableDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRouteTableDefault) Error() string {
	return fmt.Sprintf("[PUT /routeTables/{rtId}][%d] UpdateRouteTable default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRouteTableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
