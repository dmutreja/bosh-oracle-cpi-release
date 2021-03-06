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

// GetInstanceConsoleConnectionReader is a Reader for the GetInstanceConsoleConnection structure.
type GetInstanceConsoleConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceConsoleConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstanceConsoleConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstanceConsoleConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstanceConsoleConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstanceConsoleConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstanceConsoleConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstanceConsoleConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInstanceConsoleConnectionOK creates a GetInstanceConsoleConnectionOK with default headers values
func NewGetInstanceConsoleConnectionOK() *GetInstanceConsoleConnectionOK {
	return &GetInstanceConsoleConnectionOK{}
}

/*GetInstanceConsoleConnectionOK handles this case with default header values.

The instance console connection was retrieved.
*/
type GetInstanceConsoleConnectionOK struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.InstanceConsoleConnection
}

func (o *GetInstanceConsoleConnectionOK) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] getInstanceConsoleConnectionOK  %+v", 200, o.Payload)
}

func (o *GetInstanceConsoleConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.InstanceConsoleConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConsoleConnectionBadRequest creates a GetInstanceConsoleConnectionBadRequest with default headers values
func NewGetInstanceConsoleConnectionBadRequest() *GetInstanceConsoleConnectionBadRequest {
	return &GetInstanceConsoleConnectionBadRequest{}
}

/*GetInstanceConsoleConnectionBadRequest handles this case with default header values.

Bad Request
*/
type GetInstanceConsoleConnectionBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInstanceConsoleConnectionBadRequest) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] getInstanceConsoleConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstanceConsoleConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConsoleConnectionUnauthorized creates a GetInstanceConsoleConnectionUnauthorized with default headers values
func NewGetInstanceConsoleConnectionUnauthorized() *GetInstanceConsoleConnectionUnauthorized {
	return &GetInstanceConsoleConnectionUnauthorized{}
}

/*GetInstanceConsoleConnectionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetInstanceConsoleConnectionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInstanceConsoleConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] getInstanceConsoleConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstanceConsoleConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConsoleConnectionNotFound creates a GetInstanceConsoleConnectionNotFound with default headers values
func NewGetInstanceConsoleConnectionNotFound() *GetInstanceConsoleConnectionNotFound {
	return &GetInstanceConsoleConnectionNotFound{}
}

/*GetInstanceConsoleConnectionNotFound handles this case with default header values.

Not Found
*/
type GetInstanceConsoleConnectionNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInstanceConsoleConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] getInstanceConsoleConnectionNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceConsoleConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConsoleConnectionInternalServerError creates a GetInstanceConsoleConnectionInternalServerError with default headers values
func NewGetInstanceConsoleConnectionInternalServerError() *GetInstanceConsoleConnectionInternalServerError {
	return &GetInstanceConsoleConnectionInternalServerError{}
}

/*GetInstanceConsoleConnectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetInstanceConsoleConnectionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInstanceConsoleConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] getInstanceConsoleConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstanceConsoleConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConsoleConnectionDefault creates a GetInstanceConsoleConnectionDefault with default headers values
func NewGetInstanceConsoleConnectionDefault(code int) *GetInstanceConsoleConnectionDefault {
	return &GetInstanceConsoleConnectionDefault{
		_statusCode: code,
	}
}

/*GetInstanceConsoleConnectionDefault handles this case with default header values.

An error has occurred.
*/
type GetInstanceConsoleConnectionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get instance console connection default response
func (o *GetInstanceConsoleConnectionDefault) Code() int {
	return o._statusCode
}

func (o *GetInstanceConsoleConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleConnections/{instanceConsoleConnectionId}][%d] GetInstanceConsoleConnection default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstanceConsoleConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
