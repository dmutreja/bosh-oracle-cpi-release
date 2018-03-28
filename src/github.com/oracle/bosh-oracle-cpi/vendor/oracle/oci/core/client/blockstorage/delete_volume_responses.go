// Code generated by go-swagger; DO NOT EDIT.

package blockstorage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// DeleteVolumeReader is a Reader for the DeleteVolume structure.
type DeleteVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVolumeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteVolumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteVolumeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVolumeNoContent creates a DeleteVolumeNoContent with default headers values
func NewDeleteVolumeNoContent() *DeleteVolumeNoContent {
	return &DeleteVolumeNoContent{}
}

/*DeleteVolumeNoContent handles this case with default header values.

The volume is being deleted.
*/
type DeleteVolumeNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteVolumeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeNoContent ", 204)
}

func (o *DeleteVolumeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteVolumeBadRequest creates a DeleteVolumeBadRequest with default headers values
func NewDeleteVolumeBadRequest() *DeleteVolumeBadRequest {
	return &DeleteVolumeBadRequest{}
}

/*DeleteVolumeBadRequest handles this case with default header values.

Bad Request
*/
type DeleteVolumeBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVolumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeUnauthorized creates a DeleteVolumeUnauthorized with default headers values
func NewDeleteVolumeUnauthorized() *DeleteVolumeUnauthorized {
	return &DeleteVolumeUnauthorized{}
}

/*DeleteVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteVolumeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeNotFound creates a DeleteVolumeNotFound with default headers values
func NewDeleteVolumeNotFound() *DeleteVolumeNotFound {
	return &DeleteVolumeNotFound{}
}

/*DeleteVolumeNotFound handles this case with default header values.

Not Found
*/
type DeleteVolumeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeConflict creates a DeleteVolumeConflict with default headers values
func NewDeleteVolumeConflict() *DeleteVolumeConflict {
	return &DeleteVolumeConflict{}
}

/*DeleteVolumeConflict handles this case with default header values.

Conflict
*/
type DeleteVolumeConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeConflict) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeConflict  %+v", 409, o.Payload)
}

func (o *DeleteVolumeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeInternalServerError creates a DeleteVolumeInternalServerError with default headers values
func NewDeleteVolumeInternalServerError() *DeleteVolumeInternalServerError {
	return &DeleteVolumeInternalServerError{}
}

/*DeleteVolumeInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteVolumeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] deleteVolumeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumeDefault creates a DeleteVolumeDefault with default headers values
func NewDeleteVolumeDefault(code int) *DeleteVolumeDefault {
	return &DeleteVolumeDefault{
		_statusCode: code,
	}
}

/*DeleteVolumeDefault handles this case with default header values.

An error has occurred.
*/
type DeleteVolumeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete volume default response
func (o *DeleteVolumeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVolumeDefault) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{volumeId}][%d] DeleteVolume default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
