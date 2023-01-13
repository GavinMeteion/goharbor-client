// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// ListProvidersReader is a Reader for the ListProviders structure.
type ListProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProvidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProvidersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProvidersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProvidersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProvidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProvidersOK creates a ListProvidersOK with default headers values
func NewListProvidersOK() *ListProvidersOK {
	return &ListProvidersOK{}
}

/*
ListProvidersOK handles this case with default header values.

Success
*/
type ListProvidersOK struct {
	Payload []*model.Metadata
}

func (o *ListProvidersOK) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersOK  %+v", 200, o.Payload)
}

func (o *ListProvidersOK) GetPayload() []*model.Metadata {
	return o.Payload
}

func (o *ListProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersBadRequest creates a ListProvidersBadRequest with default headers values
func NewListProvidersBadRequest() *ListProvidersBadRequest {
	return &ListProvidersBadRequest{}
}

/*
ListProvidersBadRequest handles this case with default header values.

Bad request
*/
type ListProvidersBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *ListProvidersBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersUnauthorized creates a ListProvidersUnauthorized with default headers values
func NewListProvidersUnauthorized() *ListProvidersUnauthorized {
	return &ListProvidersUnauthorized{}
}

/*
ListProvidersUnauthorized handles this case with default header values.

Unauthorized
*/
type ListProvidersUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProvidersUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersForbidden creates a ListProvidersForbidden with default headers values
func NewListProvidersForbidden() *ListProvidersForbidden {
	return &ListProvidersForbidden{}
}

/*
ListProvidersForbidden handles this case with default header values.

Forbidden
*/
type ListProvidersForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersForbidden) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersForbidden  %+v", 403, o.Payload)
}

func (o *ListProvidersForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersNotFound creates a ListProvidersNotFound with default headers values
func NewListProvidersNotFound() *ListProvidersNotFound {
	return &ListProvidersNotFound{}
}

/*
ListProvidersNotFound handles this case with default header values.

Not found
*/
type ListProvidersNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersNotFound) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersNotFound  %+v", 404, o.Payload)
}

func (o *ListProvidersNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProvidersInternalServerError creates a ListProvidersInternalServerError with default headers values
func NewListProvidersInternalServerError() *ListProvidersInternalServerError {
	return &ListProvidersInternalServerError{}
}

/*
ListProvidersInternalServerError handles this case with default header values.

Internal server error
*/
type ListProvidersInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListProvidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/providers][%d] listProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProvidersInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListProvidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
