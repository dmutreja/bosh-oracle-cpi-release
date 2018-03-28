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

// GetFastConnectProviderServiceReader is a Reader for the GetFastConnectProviderService structure.
type GetFastConnectProviderServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFastConnectProviderServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFastConnectProviderServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFastConnectProviderServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetFastConnectProviderServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFastConnectProviderServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFastConnectProviderServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetFastConnectProviderServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFastConnectProviderServiceOK creates a GetFastConnectProviderServiceOK with default headers values
func NewGetFastConnectProviderServiceOK() *GetFastConnectProviderServiceOK {
	return &GetFastConnectProviderServiceOK{}
}

/*GetFastConnectProviderServiceOK handles this case with default header values.

The provider service is being retrieved.
*/
type GetFastConnectProviderServiceOK struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.FastConnectProviderService
}

func (o *GetFastConnectProviderServiceOK) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] getFastConnectProviderServiceOK  %+v", 200, o.Payload)
}

func (o *GetFastConnectProviderServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.FastConnectProviderService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFastConnectProviderServiceBadRequest creates a GetFastConnectProviderServiceBadRequest with default headers values
func NewGetFastConnectProviderServiceBadRequest() *GetFastConnectProviderServiceBadRequest {
	return &GetFastConnectProviderServiceBadRequest{}
}

/*GetFastConnectProviderServiceBadRequest handles this case with default header values.

Bad Request
*/
type GetFastConnectProviderServiceBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetFastConnectProviderServiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] getFastConnectProviderServiceBadRequest  %+v", 400, o.Payload)
}

func (o *GetFastConnectProviderServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFastConnectProviderServiceUnauthorized creates a GetFastConnectProviderServiceUnauthorized with default headers values
func NewGetFastConnectProviderServiceUnauthorized() *GetFastConnectProviderServiceUnauthorized {
	return &GetFastConnectProviderServiceUnauthorized{}
}

/*GetFastConnectProviderServiceUnauthorized handles this case with default header values.

Unauthorized
*/
type GetFastConnectProviderServiceUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetFastConnectProviderServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] getFastConnectProviderServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFastConnectProviderServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFastConnectProviderServiceNotFound creates a GetFastConnectProviderServiceNotFound with default headers values
func NewGetFastConnectProviderServiceNotFound() *GetFastConnectProviderServiceNotFound {
	return &GetFastConnectProviderServiceNotFound{}
}

/*GetFastConnectProviderServiceNotFound handles this case with default header values.

Not Found
*/
type GetFastConnectProviderServiceNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetFastConnectProviderServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] getFastConnectProviderServiceNotFound  %+v", 404, o.Payload)
}

func (o *GetFastConnectProviderServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFastConnectProviderServiceInternalServerError creates a GetFastConnectProviderServiceInternalServerError with default headers values
func NewGetFastConnectProviderServiceInternalServerError() *GetFastConnectProviderServiceInternalServerError {
	return &GetFastConnectProviderServiceInternalServerError{}
}

/*GetFastConnectProviderServiceInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetFastConnectProviderServiceInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetFastConnectProviderServiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] getFastConnectProviderServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFastConnectProviderServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFastConnectProviderServiceDefault creates a GetFastConnectProviderServiceDefault with default headers values
func NewGetFastConnectProviderServiceDefault(code int) *GetFastConnectProviderServiceDefault {
	return &GetFastConnectProviderServiceDefault{
		_statusCode: code,
	}
}

/*GetFastConnectProviderServiceDefault handles this case with default header values.

An error has occurred.
*/
type GetFastConnectProviderServiceDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get fast connect provider service default response
func (o *GetFastConnectProviderServiceDefault) Code() int {
	return o._statusCode
}

func (o *GetFastConnectProviderServiceDefault) Error() string {
	return fmt.Sprintf("[GET /fastConnectProviderServices/{providerServiceId}][%d] GetFastConnectProviderService default  %+v", o._statusCode, o.Payload)
}

func (o *GetFastConnectProviderServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
