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

// GetIPSecConnectionReader is a Reader for the GetIPSecConnection structure.
type GetIPSecConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPSecConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPSecConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetIPSecConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetIPSecConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetIPSecConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetIPSecConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPSecConnectionOK creates a GetIPSecConnectionOK with default headers values
func NewGetIPSecConnectionOK() *GetIPSecConnectionOK {
	return &GetIPSecConnectionOK{}
}

/*GetIPSecConnectionOK handles this case with default header values.

The information was retrieved.
*/
type GetIPSecConnectionOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.IPSecConnection
}

func (o *GetIPSecConnectionOK) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}][%d] getIpSecConnectionOK  %+v", 200, o.Payload)
}

func (o *GetIPSecConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.IPSecConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionUnauthorized creates a GetIPSecConnectionUnauthorized with default headers values
func NewGetIPSecConnectionUnauthorized() *GetIPSecConnectionUnauthorized {
	return &GetIPSecConnectionUnauthorized{}
}

/*GetIPSecConnectionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetIPSecConnectionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}][%d] getIpSecConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIPSecConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionNotFound creates a GetIPSecConnectionNotFound with default headers values
func NewGetIPSecConnectionNotFound() *GetIPSecConnectionNotFound {
	return &GetIPSecConnectionNotFound{}
}

/*GetIPSecConnectionNotFound handles this case with default header values.

Not Found
*/
type GetIPSecConnectionNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}][%d] getIpSecConnectionNotFound  %+v", 404, o.Payload)
}

func (o *GetIPSecConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionInternalServerError creates a GetIPSecConnectionInternalServerError with default headers values
func NewGetIPSecConnectionInternalServerError() *GetIPSecConnectionInternalServerError {
	return &GetIPSecConnectionInternalServerError{}
}

/*GetIPSecConnectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetIPSecConnectionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}][%d] getIpSecConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIPSecConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDefault creates a GetIPSecConnectionDefault with default headers values
func NewGetIPSecConnectionDefault(code int) *GetIPSecConnectionDefault {
	return &GetIPSecConnectionDefault{
		_statusCode: code,
	}
}

/*GetIPSecConnectionDefault handles this case with default header values.

An error has occurred.
*/
type GetIPSecConnectionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get IP sec connection default response
func (o *GetIPSecConnectionDefault) Code() int {
	return o._statusCode
}

func (o *GetIPSecConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}][%d] GetIPSecConnection default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPSecConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
