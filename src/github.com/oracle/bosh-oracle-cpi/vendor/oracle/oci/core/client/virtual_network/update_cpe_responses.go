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

// UpdateCpeReader is a Reader for the UpdateCpe structure.
type UpdateCpeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCpeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCpeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateCpeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateCpeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateCpeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateCpeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateCpePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateCpeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateCpeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCpeOK creates a UpdateCpeOK with default headers values
func NewUpdateCpeOK() *UpdateCpeOK {
	return &UpdateCpeOK{}
}

/*UpdateCpeOK handles this case with default header values.

The CPE was updated.
*/
type UpdateCpeOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Cpe
}

func (o *UpdateCpeOK) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeOK  %+v", 200, o.Payload)
}

func (o *UpdateCpeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Cpe)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeBadRequest creates a UpdateCpeBadRequest with default headers values
func NewUpdateCpeBadRequest() *UpdateCpeBadRequest {
	return &UpdateCpeBadRequest{}
}

/*UpdateCpeBadRequest handles this case with default header values.

Bad Request
*/
type UpdateCpeBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCpeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeUnauthorized creates a UpdateCpeUnauthorized with default headers values
func NewUpdateCpeUnauthorized() *UpdateCpeUnauthorized {
	return &UpdateCpeUnauthorized{}
}

/*UpdateCpeUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCpeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCpeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeNotFound creates a UpdateCpeNotFound with default headers values
func NewUpdateCpeNotFound() *UpdateCpeNotFound {
	return &UpdateCpeNotFound{}
}

/*UpdateCpeNotFound handles this case with default header values.

Not Found
*/
type UpdateCpeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpeNotFound) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCpeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeConflict creates a UpdateCpeConflict with default headers values
func NewUpdateCpeConflict() *UpdateCpeConflict {
	return &UpdateCpeConflict{}
}

/*UpdateCpeConflict handles this case with default header values.

Conflict
*/
type UpdateCpeConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpeConflict) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeConflict  %+v", 409, o.Payload)
}

func (o *UpdateCpeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpePreconditionFailed creates a UpdateCpePreconditionFailed with default headers values
func NewUpdateCpePreconditionFailed() *UpdateCpePreconditionFailed {
	return &UpdateCpePreconditionFailed{}
}

/*UpdateCpePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateCpePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpePreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpePreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateCpePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeInternalServerError creates a UpdateCpeInternalServerError with default headers values
func NewUpdateCpeInternalServerError() *UpdateCpeInternalServerError {
	return &UpdateCpeInternalServerError{}
}

/*UpdateCpeInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateCpeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateCpeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] updateCpeInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCpeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCpeDefault creates a UpdateCpeDefault with default headers values
func NewUpdateCpeDefault(code int) *UpdateCpeDefault {
	return &UpdateCpeDefault{
		_statusCode: code,
	}
}

/*UpdateCpeDefault handles this case with default header values.

An error has occurred.
*/
type UpdateCpeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update cpe default response
func (o *UpdateCpeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCpeDefault) Error() string {
	return fmt.Sprintf("[PUT /cpes/{cpeId}][%d] UpdateCpe default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCpeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
