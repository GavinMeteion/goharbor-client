// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetGCLogReader is a Reader for the GetGCLog structure.
type GetGCLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGCLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGCLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGCLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGCLogOK creates a GetGCLogOK with default headers values
func NewGetGCLogOK() *GetGCLogOK {
	return &GetGCLogOK{}
}

/*
GetGCLogOK handles this case with default header values.

Get successfully.
*/
type GetGCLogOK struct {
	Payload string
}

func (o *GetGCLogOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogOK  %+v", 200, o.Payload)
}

func (o *GetGCLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetGCLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCLogBadRequest creates a GetGCLogBadRequest with default headers values
func NewGetGCLogBadRequest() *GetGCLogBadRequest {
	return &GetGCLogBadRequest{}
}

/*
GetGCLogBadRequest handles this case with default header values.

Bad request
*/
type GetGCLogBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetGCLogBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCLogUnauthorized creates a GetGCLogUnauthorized with default headers values
func NewGetGCLogUnauthorized() *GetGCLogUnauthorized {
	return &GetGCLogUnauthorized{}
}

/*
GetGCLogUnauthorized handles this case with default header values.

Unauthorized
*/
type GetGCLogUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCLogUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCLogForbidden creates a GetGCLogForbidden with default headers values
func NewGetGCLogForbidden() *GetGCLogForbidden {
	return &GetGCLogForbidden{}
}

/*
GetGCLogForbidden handles this case with default header values.

Forbidden
*/
type GetGCLogForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCLogForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogForbidden  %+v", 403, o.Payload)
}

func (o *GetGCLogForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCLogNotFound creates a GetGCLogNotFound with default headers values
func NewGetGCLogNotFound() *GetGCLogNotFound {
	return &GetGCLogNotFound{}
}

/*
GetGCLogNotFound handles this case with default header values.

Not found
*/
type GetGCLogNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCLogNotFound) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogNotFound  %+v", 404, o.Payload)
}

func (o *GetGCLogNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCLogInternalServerError creates a GetGCLogInternalServerError with default headers values
func NewGetGCLogInternalServerError() *GetGCLogInternalServerError {
	return &GetGCLogInternalServerError{}
}

/*
GetGCLogInternalServerError handles this case with default header values.

Internal server error
*/
type GetGCLogInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetGCLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}/log][%d] getGcLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCLogInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetGCLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
