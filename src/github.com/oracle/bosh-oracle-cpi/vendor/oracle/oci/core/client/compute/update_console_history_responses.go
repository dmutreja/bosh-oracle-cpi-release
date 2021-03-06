// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// UpdateConsoleHistoryReader is a Reader for the UpdateConsoleHistory structure.
type UpdateConsoleHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConsoleHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateConsoleHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateConsoleHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateConsoleHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateConsoleHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateConsoleHistoryPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateConsoleHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateConsoleHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateConsoleHistoryOK creates a UpdateConsoleHistoryOK with default headers values
func NewUpdateConsoleHistoryOK() *UpdateConsoleHistoryOK {
	return &UpdateConsoleHistoryOK{}
}

/*UpdateConsoleHistoryOK handles this case with default header values.

The request was accepted by the server
*/
type UpdateConsoleHistoryOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.ConsoleHistory
}

func (o *UpdateConsoleHistoryOK) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryOK  %+v", 200, o.Payload)
}

func (o *UpdateConsoleHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.ConsoleHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryBadRequest creates a UpdateConsoleHistoryBadRequest with default headers values
func NewUpdateConsoleHistoryBadRequest() *UpdateConsoleHistoryBadRequest {
	return &UpdateConsoleHistoryBadRequest{}
}

/*UpdateConsoleHistoryBadRequest handles this case with default header values.

Bad Request
*/
type UpdateConsoleHistoryBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateConsoleHistoryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConsoleHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryUnauthorized creates a UpdateConsoleHistoryUnauthorized with default headers values
func NewUpdateConsoleHistoryUnauthorized() *UpdateConsoleHistoryUnauthorized {
	return &UpdateConsoleHistoryUnauthorized{}
}

/*UpdateConsoleHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateConsoleHistoryUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateConsoleHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateConsoleHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryNotFound creates a UpdateConsoleHistoryNotFound with default headers values
func NewUpdateConsoleHistoryNotFound() *UpdateConsoleHistoryNotFound {
	return &UpdateConsoleHistoryNotFound{}
}

/*UpdateConsoleHistoryNotFound handles this case with default header values.

Not Found
*/
type UpdateConsoleHistoryNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateConsoleHistoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConsoleHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryPreconditionFailed creates a UpdateConsoleHistoryPreconditionFailed with default headers values
func NewUpdateConsoleHistoryPreconditionFailed() *UpdateConsoleHistoryPreconditionFailed {
	return &UpdateConsoleHistoryPreconditionFailed{}
}

/*UpdateConsoleHistoryPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateConsoleHistoryPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateConsoleHistoryPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateConsoleHistoryPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryInternalServerError creates a UpdateConsoleHistoryInternalServerError with default headers values
func NewUpdateConsoleHistoryInternalServerError() *UpdateConsoleHistoryInternalServerError {
	return &UpdateConsoleHistoryInternalServerError{}
}

/*UpdateConsoleHistoryInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateConsoleHistoryInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateConsoleHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] updateConsoleHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateConsoleHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsoleHistoryDefault creates a UpdateConsoleHistoryDefault with default headers values
func NewUpdateConsoleHistoryDefault(code int) *UpdateConsoleHistoryDefault {
	return &UpdateConsoleHistoryDefault{
		_statusCode: code,
	}
}

/*UpdateConsoleHistoryDefault handles this case with default header values.

An error has occurred.
*/
type UpdateConsoleHistoryDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update console history default response
func (o *UpdateConsoleHistoryDefault) Code() int {
	return o._statusCode
}

func (o *UpdateConsoleHistoryDefault) Error() string {
	return fmt.Sprintf("[PUT /instanceConsoleHistories/{instanceConsoleHistoryId}][%d] UpdateConsoleHistory default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateConsoleHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
