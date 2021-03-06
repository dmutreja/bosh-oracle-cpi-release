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

// ListDrgAttachmentsReader is a Reader for the ListDrgAttachments structure.
type ListDrgAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDrgAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDrgAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListDrgAttachmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListDrgAttachmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListDrgAttachmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListDrgAttachmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDrgAttachmentsOK creates a ListDrgAttachmentsOK with default headers values
func NewListDrgAttachmentsOK() *ListDrgAttachmentsOK {
	return &ListDrgAttachmentsOK{}
}

/*ListDrgAttachmentsOK handles this case with default header values.

The list is being retrieved.
*/
type ListDrgAttachmentsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.DrgAttachment
}

func (o *ListDrgAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /drgAttachments][%d] listDrgAttachmentsOK  %+v", 200, o.Payload)
}

func (o *ListDrgAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListDrgAttachmentsUnauthorized creates a ListDrgAttachmentsUnauthorized with default headers values
func NewListDrgAttachmentsUnauthorized() *ListDrgAttachmentsUnauthorized {
	return &ListDrgAttachmentsUnauthorized{}
}

/*ListDrgAttachmentsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListDrgAttachmentsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDrgAttachmentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /drgAttachments][%d] listDrgAttachmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListDrgAttachmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDrgAttachmentsNotFound creates a ListDrgAttachmentsNotFound with default headers values
func NewListDrgAttachmentsNotFound() *ListDrgAttachmentsNotFound {
	return &ListDrgAttachmentsNotFound{}
}

/*ListDrgAttachmentsNotFound handles this case with default header values.

Not Found
*/
type ListDrgAttachmentsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDrgAttachmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /drgAttachments][%d] listDrgAttachmentsNotFound  %+v", 404, o.Payload)
}

func (o *ListDrgAttachmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDrgAttachmentsInternalServerError creates a ListDrgAttachmentsInternalServerError with default headers values
func NewListDrgAttachmentsInternalServerError() *ListDrgAttachmentsInternalServerError {
	return &ListDrgAttachmentsInternalServerError{}
}

/*ListDrgAttachmentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListDrgAttachmentsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListDrgAttachmentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /drgAttachments][%d] listDrgAttachmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDrgAttachmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDrgAttachmentsDefault creates a ListDrgAttachmentsDefault with default headers values
func NewListDrgAttachmentsDefault(code int) *ListDrgAttachmentsDefault {
	return &ListDrgAttachmentsDefault{
		_statusCode: code,
	}
}

/*ListDrgAttachmentsDefault handles this case with default header values.

An error has occurred.
*/
type ListDrgAttachmentsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list drg attachments default response
func (o *ListDrgAttachmentsDefault) Code() int {
	return o._statusCode
}

func (o *ListDrgAttachmentsDefault) Error() string {
	return fmt.Sprintf("[GET /drgAttachments][%d] ListDrgAttachments default  %+v", o._statusCode, o.Payload)
}

func (o *ListDrgAttachmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
