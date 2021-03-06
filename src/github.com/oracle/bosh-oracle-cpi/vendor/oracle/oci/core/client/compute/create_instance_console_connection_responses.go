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

// CreateInstanceConsoleConnectionReader is a Reader for the CreateInstanceConsoleConnection structure.
type CreateInstanceConsoleConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstanceConsoleConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateInstanceConsoleConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateInstanceConsoleConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateInstanceConsoleConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateInstanceConsoleConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateInstanceConsoleConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateInstanceConsoleConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateInstanceConsoleConnectionOK creates a CreateInstanceConsoleConnectionOK with default headers values
func NewCreateInstanceConsoleConnectionOK() *CreateInstanceConsoleConnectionOK {
	return &CreateInstanceConsoleConnectionOK{}
}

/*CreateInstanceConsoleConnectionOK handles this case with default header values.

The instance console connection is being configured.
*/
type CreateInstanceConsoleConnectionOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.InstanceConsoleConnection
}

func (o *CreateInstanceConsoleConnectionOK) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] createInstanceConsoleConnectionOK  %+v", 200, o.Payload)
}

func (o *CreateInstanceConsoleConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.InstanceConsoleConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConsoleConnectionBadRequest creates a CreateInstanceConsoleConnectionBadRequest with default headers values
func NewCreateInstanceConsoleConnectionBadRequest() *CreateInstanceConsoleConnectionBadRequest {
	return &CreateInstanceConsoleConnectionBadRequest{}
}

/*CreateInstanceConsoleConnectionBadRequest handles this case with default header values.

Bad Request
*/
type CreateInstanceConsoleConnectionBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInstanceConsoleConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] createInstanceConsoleConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstanceConsoleConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConsoleConnectionUnauthorized creates a CreateInstanceConsoleConnectionUnauthorized with default headers values
func NewCreateInstanceConsoleConnectionUnauthorized() *CreateInstanceConsoleConnectionUnauthorized {
	return &CreateInstanceConsoleConnectionUnauthorized{}
}

/*CreateInstanceConsoleConnectionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateInstanceConsoleConnectionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInstanceConsoleConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] createInstanceConsoleConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstanceConsoleConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConsoleConnectionConflict creates a CreateInstanceConsoleConnectionConflict with default headers values
func NewCreateInstanceConsoleConnectionConflict() *CreateInstanceConsoleConnectionConflict {
	return &CreateInstanceConsoleConnectionConflict{}
}

/*CreateInstanceConsoleConnectionConflict handles this case with default header values.

Conflict
*/
type CreateInstanceConsoleConnectionConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInstanceConsoleConnectionConflict) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] createInstanceConsoleConnectionConflict  %+v", 409, o.Payload)
}

func (o *CreateInstanceConsoleConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConsoleConnectionInternalServerError creates a CreateInstanceConsoleConnectionInternalServerError with default headers values
func NewCreateInstanceConsoleConnectionInternalServerError() *CreateInstanceConsoleConnectionInternalServerError {
	return &CreateInstanceConsoleConnectionInternalServerError{}
}

/*CreateInstanceConsoleConnectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateInstanceConsoleConnectionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateInstanceConsoleConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] createInstanceConsoleConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstanceConsoleConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstanceConsoleConnectionDefault creates a CreateInstanceConsoleConnectionDefault with default headers values
func NewCreateInstanceConsoleConnectionDefault(code int) *CreateInstanceConsoleConnectionDefault {
	return &CreateInstanceConsoleConnectionDefault{
		_statusCode: code,
	}
}

/*CreateInstanceConsoleConnectionDefault handles this case with default header values.

An error has occurred.
*/
type CreateInstanceConsoleConnectionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create instance console connection default response
func (o *CreateInstanceConsoleConnectionDefault) Code() int {
	return o._statusCode
}

func (o *CreateInstanceConsoleConnectionDefault) Error() string {
	return fmt.Sprintf("[POST /instanceConsoleConnections][%d] CreateInstanceConsoleConnection default  %+v", o._statusCode, o.Payload)
}

func (o *CreateInstanceConsoleConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
