// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetWorkersReader is a Reader for the GetWorkers structure.
type GetWorkersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkersOK creates a GetWorkersOK with default headers values
func NewGetWorkersOK() *GetWorkersOK {
	return &GetWorkersOK{}
}

/*
GetWorkersOK handles this case with default header values.

Get workers successfully.
*/
type GetWorkersOK struct {
	Payload []*model.Worker
}

func (o *GetWorkersOK) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools/{pool_id}/workers][%d] getWorkersOK  %+v", 200, o.Payload)
}

func (o *GetWorkersOK) GetPayload() []*model.Worker {
	return o.Payload
}

func (o *GetWorkersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkersUnauthorized creates a GetWorkersUnauthorized with default headers values
func NewGetWorkersUnauthorized() *GetWorkersUnauthorized {
	return &GetWorkersUnauthorized{}
}

/*
GetWorkersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetWorkersUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools/{pool_id}/workers][%d] getWorkersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkersUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkersForbidden creates a GetWorkersForbidden with default headers values
func NewGetWorkersForbidden() *GetWorkersForbidden {
	return &GetWorkersForbidden{}
}

/*
GetWorkersForbidden handles this case with default header values.

Forbidden
*/
type GetWorkersForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkersForbidden) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools/{pool_id}/workers][%d] getWorkersForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkersForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkersNotFound creates a GetWorkersNotFound with default headers values
func NewGetWorkersNotFound() *GetWorkersNotFound {
	return &GetWorkersNotFound{}
}

/*
GetWorkersNotFound handles this case with default header values.

Not found
*/
type GetWorkersNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkersNotFound) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools/{pool_id}/workers][%d] getWorkersNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkersNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkersInternalServerError creates a GetWorkersInternalServerError with default headers values
func NewGetWorkersInternalServerError() *GetWorkersInternalServerError {
	return &GetWorkersInternalServerError{}
}

/*
GetWorkersInternalServerError handles this case with default header values.

Internal server error
*/
type GetWorkersInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetWorkersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobservice/pools/{pool_id}/workers][%d] getWorkersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkersInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetWorkersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
