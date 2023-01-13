// Code generated by go-swagger; DO NOT EDIT.

package system_cve_allowlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// PutSystemCVEAllowlistReader is a Reader for the PutSystemCVEAllowlist structure.
type PutSystemCVEAllowlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemCVEAllowlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutSystemCVEAllowlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutSystemCVEAllowlistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutSystemCVEAllowlistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutSystemCVEAllowlistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutSystemCVEAllowlistOK creates a PutSystemCVEAllowlistOK with default headers values
func NewPutSystemCVEAllowlistOK() *PutSystemCVEAllowlistOK {
	return &PutSystemCVEAllowlistOK{}
}

/*
PutSystemCVEAllowlistOK handles this case with default header values.

Successfully updated the CVE allowlist.
*/
type PutSystemCVEAllowlistOK struct {
}

func (o *PutSystemCVEAllowlistOK) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCveAllowlistOK ", 200)
}

func (o *PutSystemCVEAllowlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemCVEAllowlistUnauthorized creates a PutSystemCVEAllowlistUnauthorized with default headers values
func NewPutSystemCVEAllowlistUnauthorized() *PutSystemCVEAllowlistUnauthorized {
	return &PutSystemCVEAllowlistUnauthorized{}
}

/*
PutSystemCVEAllowlistUnauthorized handles this case with default header values.

Unauthorized
*/
type PutSystemCVEAllowlistUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PutSystemCVEAllowlistUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCveAllowlistUnauthorized  %+v", 401, o.Payload)
}

func (o *PutSystemCVEAllowlistUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PutSystemCVEAllowlistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSystemCVEAllowlistForbidden creates a PutSystemCVEAllowlistForbidden with default headers values
func NewPutSystemCVEAllowlistForbidden() *PutSystemCVEAllowlistForbidden {
	return &PutSystemCVEAllowlistForbidden{}
}

/*
PutSystemCVEAllowlistForbidden handles this case with default header values.

Forbidden
*/
type PutSystemCVEAllowlistForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PutSystemCVEAllowlistForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCveAllowlistForbidden  %+v", 403, o.Payload)
}

func (o *PutSystemCVEAllowlistForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PutSystemCVEAllowlistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSystemCVEAllowlistInternalServerError creates a PutSystemCVEAllowlistInternalServerError with default headers values
func NewPutSystemCVEAllowlistInternalServerError() *PutSystemCVEAllowlistInternalServerError {
	return &PutSystemCVEAllowlistInternalServerError{}
}

/*
PutSystemCVEAllowlistInternalServerError handles this case with default header values.

Internal server error
*/
type PutSystemCVEAllowlistInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PutSystemCVEAllowlistInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/CVEAllowlist][%d] putSystemCveAllowlistInternalServerError  %+v", 500, o.Payload)
}

func (o *PutSystemCVEAllowlistInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PutSystemCVEAllowlistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
