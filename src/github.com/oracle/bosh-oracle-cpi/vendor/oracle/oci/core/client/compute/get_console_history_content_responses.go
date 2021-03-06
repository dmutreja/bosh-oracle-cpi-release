// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// GetConsoleHistoryContentReader is a Reader for the GetConsoleHistoryContent structure.
type GetConsoleHistoryContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsoleHistoryContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConsoleHistoryContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetConsoleHistoryContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetConsoleHistoryContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetConsoleHistoryContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetConsoleHistoryContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetConsoleHistoryContentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConsoleHistoryContentOK creates a GetConsoleHistoryContentOK with default headers values
func NewGetConsoleHistoryContentOK() *GetConsoleHistoryContentOK {
	return &GetConsoleHistoryContentOK{}
}

/*GetConsoleHistoryContentOK handles this case with default header values.

The information is being retrieved.
*/
type GetConsoleHistoryContentOK struct {
	/*The number of bytes remaining in the snapshot.
	 */
	OpcBytesRemaining int32
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload string
}

func (o *GetConsoleHistoryContentOK) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] getConsoleHistoryContentOK  %+v", 200, o.Payload)
}

func (o *GetConsoleHistoryContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-bytes-remaining
	opcBytesRemaining, err := swag.ConvertInt32(response.GetHeader("opc-bytes-remaining"))
	if err != nil {
		return errors.InvalidType("opc-bytes-remaining", "header", "int32", response.GetHeader("opc-bytes-remaining"))
	}
	o.OpcBytesRemaining = opcBytesRemaining

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsoleHistoryContentBadRequest creates a GetConsoleHistoryContentBadRequest with default headers values
func NewGetConsoleHistoryContentBadRequest() *GetConsoleHistoryContentBadRequest {
	return &GetConsoleHistoryContentBadRequest{}
}

/*GetConsoleHistoryContentBadRequest handles this case with default header values.

Bad Request
*/
type GetConsoleHistoryContentBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetConsoleHistoryContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] getConsoleHistoryContentBadRequest  %+v", 400, o.Payload)
}

func (o *GetConsoleHistoryContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsoleHistoryContentUnauthorized creates a GetConsoleHistoryContentUnauthorized with default headers values
func NewGetConsoleHistoryContentUnauthorized() *GetConsoleHistoryContentUnauthorized {
	return &GetConsoleHistoryContentUnauthorized{}
}

/*GetConsoleHistoryContentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetConsoleHistoryContentUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetConsoleHistoryContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] getConsoleHistoryContentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConsoleHistoryContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsoleHistoryContentNotFound creates a GetConsoleHistoryContentNotFound with default headers values
func NewGetConsoleHistoryContentNotFound() *GetConsoleHistoryContentNotFound {
	return &GetConsoleHistoryContentNotFound{}
}

/*GetConsoleHistoryContentNotFound handles this case with default header values.

Not Found
*/
type GetConsoleHistoryContentNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetConsoleHistoryContentNotFound) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] getConsoleHistoryContentNotFound  %+v", 404, o.Payload)
}

func (o *GetConsoleHistoryContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsoleHistoryContentInternalServerError creates a GetConsoleHistoryContentInternalServerError with default headers values
func NewGetConsoleHistoryContentInternalServerError() *GetConsoleHistoryContentInternalServerError {
	return &GetConsoleHistoryContentInternalServerError{}
}

/*GetConsoleHistoryContentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetConsoleHistoryContentInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetConsoleHistoryContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] getConsoleHistoryContentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConsoleHistoryContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsoleHistoryContentDefault creates a GetConsoleHistoryContentDefault with default headers values
func NewGetConsoleHistoryContentDefault(code int) *GetConsoleHistoryContentDefault {
	return &GetConsoleHistoryContentDefault{
		_statusCode: code,
	}
}

/*GetConsoleHistoryContentDefault handles this case with default header values.

An error has occurred.
*/
type GetConsoleHistoryContentDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get console history content default response
func (o *GetConsoleHistoryContentDefault) Code() int {
	return o._statusCode
}

func (o *GetConsoleHistoryContentDefault) Error() string {
	return fmt.Sprintf("[GET /instanceConsoleHistories/{instanceConsoleHistoryId}/data][%d] GetConsoleHistoryContent default  %+v", o._statusCode, o.Payload)
}

func (o *GetConsoleHistoryContentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
