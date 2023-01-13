// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetPurgeHistoryReader is a Reader for the GetPurgeHistory structure.
type GetPurgeHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPurgeHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPurgeHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPurgeHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPurgeHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPurgeHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPurgeHistoryOK creates a GetPurgeHistoryOK with default headers values
func NewGetPurgeHistoryOK() *GetPurgeHistoryOK {
	return &GetPurgeHistoryOK{}
}

/*
GetPurgeHistoryOK handles this case with default header values.

Get purge job results successfully.
*/
type GetPurgeHistoryOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of history
	 */
	XTotalCount int64

	Payload []*model.ExecHistory
}

func (o *GetPurgeHistoryOK) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit][%d] getPurgeHistoryOK  %+v", 200, o.Payload)
}

func (o *GetPurgeHistoryOK) GetPayload() []*model.ExecHistory {
	return o.Payload
}

func (o *GetPurgeHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeHistoryUnauthorized creates a GetPurgeHistoryUnauthorized with default headers values
func NewGetPurgeHistoryUnauthorized() *GetPurgeHistoryUnauthorized {
	return &GetPurgeHistoryUnauthorized{}
}

/*
GetPurgeHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPurgeHistoryUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit][%d] getPurgeHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPurgeHistoryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeHistoryForbidden creates a GetPurgeHistoryForbidden with default headers values
func NewGetPurgeHistoryForbidden() *GetPurgeHistoryForbidden {
	return &GetPurgeHistoryForbidden{}
}

/*
GetPurgeHistoryForbidden handles this case with default header values.

Forbidden
*/
type GetPurgeHistoryForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit][%d] getPurgeHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetPurgeHistoryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPurgeHistoryInternalServerError creates a GetPurgeHistoryInternalServerError with default headers values
func NewGetPurgeHistoryInternalServerError() *GetPurgeHistoryInternalServerError {
	return &GetPurgeHistoryInternalServerError{}
}

/*
GetPurgeHistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetPurgeHistoryInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPurgeHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/purgeaudit][%d] getPurgeHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPurgeHistoryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPurgeHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
