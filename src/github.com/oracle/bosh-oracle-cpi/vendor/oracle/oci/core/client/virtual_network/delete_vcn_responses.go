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

// DeleteVcnReader is a Reader for the DeleteVcn structure.
type DeleteVcnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVcnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVcnNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteVcnUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteVcnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteVcnPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteVcnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVcnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVcnNoContent creates a DeleteVcnNoContent with default headers values
func NewDeleteVcnNoContent() *DeleteVcnNoContent {
	return &DeleteVcnNoContent{}
}

/*DeleteVcnNoContent handles this case with default header values.

The VCN is being deleted.
*/
type DeleteVcnNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteVcnNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] deleteVcnNoContent ", 204)
}

func (o *DeleteVcnNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteVcnUnauthorized creates a DeleteVcnUnauthorized with default headers values
func NewDeleteVcnUnauthorized() *DeleteVcnUnauthorized {
	return &DeleteVcnUnauthorized{}
}

/*DeleteVcnUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteVcnUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVcnUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] deleteVcnUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteVcnUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVcnNotFound creates a DeleteVcnNotFound with default headers values
func NewDeleteVcnNotFound() *DeleteVcnNotFound {
	return &DeleteVcnNotFound{}
}

/*DeleteVcnNotFound handles this case with default header values.

Not Found
*/
type DeleteVcnNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVcnNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] deleteVcnNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVcnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVcnPreconditionFailed creates a DeleteVcnPreconditionFailed with default headers values
func NewDeleteVcnPreconditionFailed() *DeleteVcnPreconditionFailed {
	return &DeleteVcnPreconditionFailed{}
}

/*DeleteVcnPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteVcnPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVcnPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] deleteVcnPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteVcnPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVcnInternalServerError creates a DeleteVcnInternalServerError with default headers values
func NewDeleteVcnInternalServerError() *DeleteVcnInternalServerError {
	return &DeleteVcnInternalServerError{}
}

/*DeleteVcnInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteVcnInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVcnInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] deleteVcnInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVcnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVcnDefault creates a DeleteVcnDefault with default headers values
func NewDeleteVcnDefault(code int) *DeleteVcnDefault {
	return &DeleteVcnDefault{
		_statusCode: code,
	}
}

/*DeleteVcnDefault handles this case with default header values.

An error has occurred.
*/
type DeleteVcnDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete vcn default response
func (o *DeleteVcnDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVcnDefault) Error() string {
	return fmt.Sprintf("[DELETE /vcns/{vcnId}][%d] DeleteVcn default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVcnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
