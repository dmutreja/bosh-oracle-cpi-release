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

// CreateSubnetReader is a Reader for the CreateSubnet structure.
type CreateSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateSubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSubnetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSubnetOK creates a CreateSubnetOK with default headers values
func NewCreateSubnetOK() *CreateSubnetOK {
	return &CreateSubnetOK{}
}

/*CreateSubnetOK handles this case with default header values.

The subnet was created.
*/
type CreateSubnetOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Subnet
}

func (o *CreateSubnetOK) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetOK  %+v", 200, o.Payload)
}

func (o *CreateSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Subnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetBadRequest creates a CreateSubnetBadRequest with default headers values
func NewCreateSubnetBadRequest() *CreateSubnetBadRequest {
	return &CreateSubnetBadRequest{}
}

/*CreateSubnetBadRequest handles this case with default header values.

Bad Request
*/
type CreateSubnetBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSubnetBadRequest) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetUnauthorized creates a CreateSubnetUnauthorized with default headers values
func NewCreateSubnetUnauthorized() *CreateSubnetUnauthorized {
	return &CreateSubnetUnauthorized{}
}

/*CreateSubnetUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSubnetUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSubnetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateSubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetNotFound creates a CreateSubnetNotFound with default headers values
func NewCreateSubnetNotFound() *CreateSubnetNotFound {
	return &CreateSubnetNotFound{}
}

/*CreateSubnetNotFound handles this case with default header values.

Not Found
*/
type CreateSubnetNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSubnetNotFound) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetNotFound  %+v", 404, o.Payload)
}

func (o *CreateSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetConflict creates a CreateSubnetConflict with default headers values
func NewCreateSubnetConflict() *CreateSubnetConflict {
	return &CreateSubnetConflict{}
}

/*CreateSubnetConflict handles this case with default header values.

Conflict
*/
type CreateSubnetConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSubnetConflict) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetConflict  %+v", 409, o.Payload)
}

func (o *CreateSubnetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetInternalServerError creates a CreateSubnetInternalServerError with default headers values
func NewCreateSubnetInternalServerError() *CreateSubnetInternalServerError {
	return &CreateSubnetInternalServerError{}
}

/*CreateSubnetInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateSubnetInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] createSubnetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSubnetDefault creates a CreateSubnetDefault with default headers values
func NewCreateSubnetDefault(code int) *CreateSubnetDefault {
	return &CreateSubnetDefault{
		_statusCode: code,
	}
}

/*CreateSubnetDefault handles this case with default header values.

An error has occurred.
*/
type CreateSubnetDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create subnet default response
func (o *CreateSubnetDefault) Code() int {
	return o._statusCode
}

func (o *CreateSubnetDefault) Error() string {
	return fmt.Sprintf("[POST /subnets][%d] CreateSubnet default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
