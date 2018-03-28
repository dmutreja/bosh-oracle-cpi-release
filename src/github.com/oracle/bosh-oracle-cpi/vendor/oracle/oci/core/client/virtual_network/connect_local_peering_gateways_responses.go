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

// ConnectLocalPeeringGatewaysReader is a Reader for the ConnectLocalPeeringGateways structure.
type ConnectLocalPeeringGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectLocalPeeringGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewConnectLocalPeeringGatewaysNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConnectLocalPeeringGatewaysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConnectLocalPeeringGatewaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConnectLocalPeeringGatewaysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewConnectLocalPeeringGatewaysConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewConnectLocalPeeringGatewaysPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConnectLocalPeeringGatewaysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewConnectLocalPeeringGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectLocalPeeringGatewaysNoContent creates a ConnectLocalPeeringGatewaysNoContent with default headers values
func NewConnectLocalPeeringGatewaysNoContent() *ConnectLocalPeeringGatewaysNoContent {
	return &ConnectLocalPeeringGatewaysNoContent{}
}

/*ConnectLocalPeeringGatewaysNoContent handles this case with default header values.

The connect request was accepted and a peering is being established.
*/
type ConnectLocalPeeringGatewaysNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *ConnectLocalPeeringGatewaysNoContent) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysNoContent ", 204)
}

func (o *ConnectLocalPeeringGatewaysNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewConnectLocalPeeringGatewaysBadRequest creates a ConnectLocalPeeringGatewaysBadRequest with default headers values
func NewConnectLocalPeeringGatewaysBadRequest() *ConnectLocalPeeringGatewaysBadRequest {
	return &ConnectLocalPeeringGatewaysBadRequest{}
}

/*ConnectLocalPeeringGatewaysBadRequest handles this case with default header values.

Bad Request
*/
type ConnectLocalPeeringGatewaysBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysBadRequest) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysBadRequest  %+v", 400, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysUnauthorized creates a ConnectLocalPeeringGatewaysUnauthorized with default headers values
func NewConnectLocalPeeringGatewaysUnauthorized() *ConnectLocalPeeringGatewaysUnauthorized {
	return &ConnectLocalPeeringGatewaysUnauthorized{}
}

/*ConnectLocalPeeringGatewaysUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectLocalPeeringGatewaysUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysUnauthorized  %+v", 401, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysNotFound creates a ConnectLocalPeeringGatewaysNotFound with default headers values
func NewConnectLocalPeeringGatewaysNotFound() *ConnectLocalPeeringGatewaysNotFound {
	return &ConnectLocalPeeringGatewaysNotFound{}
}

/*ConnectLocalPeeringGatewaysNotFound handles this case with default header values.

Not Found
*/
type ConnectLocalPeeringGatewaysNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysNotFound) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysNotFound  %+v", 404, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysConflict creates a ConnectLocalPeeringGatewaysConflict with default headers values
func NewConnectLocalPeeringGatewaysConflict() *ConnectLocalPeeringGatewaysConflict {
	return &ConnectLocalPeeringGatewaysConflict{}
}

/*ConnectLocalPeeringGatewaysConflict handles this case with default header values.

Conflict
*/
type ConnectLocalPeeringGatewaysConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysConflict) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysConflict  %+v", 409, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysPreconditionFailed creates a ConnectLocalPeeringGatewaysPreconditionFailed with default headers values
func NewConnectLocalPeeringGatewaysPreconditionFailed() *ConnectLocalPeeringGatewaysPreconditionFailed {
	return &ConnectLocalPeeringGatewaysPreconditionFailed{}
}

/*ConnectLocalPeeringGatewaysPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type ConnectLocalPeeringGatewaysPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysInternalServerError creates a ConnectLocalPeeringGatewaysInternalServerError with default headers values
func NewConnectLocalPeeringGatewaysInternalServerError() *ConnectLocalPeeringGatewaysInternalServerError {
	return &ConnectLocalPeeringGatewaysInternalServerError{}
}

/*ConnectLocalPeeringGatewaysInternalServerError handles this case with default header values.

Internal Server Error
*/
type ConnectLocalPeeringGatewaysInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ConnectLocalPeeringGatewaysInternalServerError) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] connectLocalPeeringGatewaysInternalServerError  %+v", 500, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectLocalPeeringGatewaysDefault creates a ConnectLocalPeeringGatewaysDefault with default headers values
func NewConnectLocalPeeringGatewaysDefault(code int) *ConnectLocalPeeringGatewaysDefault {
	return &ConnectLocalPeeringGatewaysDefault{
		_statusCode: code,
	}
}

/*ConnectLocalPeeringGatewaysDefault handles this case with default header values.

An error has occurred.
*/
type ConnectLocalPeeringGatewaysDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the connect local peering gateways default response
func (o *ConnectLocalPeeringGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *ConnectLocalPeeringGatewaysDefault) Error() string {
	return fmt.Sprintf("[POST /localPeeringGateways/{localPeeringGatewayId}/actions/connect][%d] ConnectLocalPeeringGateways default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectLocalPeeringGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
