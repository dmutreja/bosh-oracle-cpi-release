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

// ListIPSecConnectionsReader is a Reader for the ListIPSecConnections structure.
type ListIPSecConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIPSecConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListIPSecConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListIPSecConnectionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListIPSecConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListIPSecConnectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListIPSecConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListIPSecConnectionsOK creates a ListIPSecConnectionsOK with default headers values
func NewListIPSecConnectionsOK() *ListIPSecConnectionsOK {
	return &ListIPSecConnectionsOK{}
}

/*ListIPSecConnectionsOK handles this case with default header values.

The list is being retrieved.
*/
type ListIPSecConnectionsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.IPSecConnection
}

func (o *ListIPSecConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections][%d] listIpSecConnectionsOK  %+v", 200, o.Payload)
}

func (o *ListIPSecConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPSecConnectionsUnauthorized creates a ListIPSecConnectionsUnauthorized with default headers values
func NewListIPSecConnectionsUnauthorized() *ListIPSecConnectionsUnauthorized {
	return &ListIPSecConnectionsUnauthorized{}
}

/*ListIPSecConnectionsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListIPSecConnectionsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListIPSecConnectionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections][%d] listIpSecConnectionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListIPSecConnectionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPSecConnectionsNotFound creates a ListIPSecConnectionsNotFound with default headers values
func NewListIPSecConnectionsNotFound() *ListIPSecConnectionsNotFound {
	return &ListIPSecConnectionsNotFound{}
}

/*ListIPSecConnectionsNotFound handles this case with default header values.

Not Found
*/
type ListIPSecConnectionsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListIPSecConnectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections][%d] listIpSecConnectionsNotFound  %+v", 404, o.Payload)
}

func (o *ListIPSecConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPSecConnectionsInternalServerError creates a ListIPSecConnectionsInternalServerError with default headers values
func NewListIPSecConnectionsInternalServerError() *ListIPSecConnectionsInternalServerError {
	return &ListIPSecConnectionsInternalServerError{}
}

/*ListIPSecConnectionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListIPSecConnectionsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListIPSecConnectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections][%d] listIpSecConnectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIPSecConnectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPSecConnectionsDefault creates a ListIPSecConnectionsDefault with default headers values
func NewListIPSecConnectionsDefault(code int) *ListIPSecConnectionsDefault {
	return &ListIPSecConnectionsDefault{
		_statusCode: code,
	}
}

/*ListIPSecConnectionsDefault handles this case with default header values.

An error has occurred.
*/
type ListIPSecConnectionsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list IP sec connections default response
func (o *ListIPSecConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *ListIPSecConnectionsDefault) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections][%d] ListIPSecConnections default  %+v", o._statusCode, o.Payload)
}

func (o *ListIPSecConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
