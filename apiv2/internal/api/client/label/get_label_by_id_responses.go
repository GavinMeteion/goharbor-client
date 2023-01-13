// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetLabelByIDReader is a Reader for the GetLabelByID structure.
type GetLabelByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLabelByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLabelByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLabelByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelByIDOK creates a GetLabelByIDOK with default headers values
func NewGetLabelByIDOK() *GetLabelByIDOK {
	return &GetLabelByIDOK{}
}

/*
GetLabelByIDOK handles this case with default header values.

Get successfully.
*/
type GetLabelByIDOK struct {
	Payload *model.Label
}

func (o *GetLabelByIDOK) Error() string {
	return fmt.Sprintf("[GET /labels/{label_id}][%d] getLabelByIdOK  %+v", 200, o.Payload)
}

func (o *GetLabelByIDOK) GetPayload() *model.Label {
	return o.Payload
}

func (o *GetLabelByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelByIDUnauthorized creates a GetLabelByIDUnauthorized with default headers values
func NewGetLabelByIDUnauthorized() *GetLabelByIDUnauthorized {
	return &GetLabelByIDUnauthorized{}
}

/*
GetLabelByIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLabelByIDUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLabelByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /labels/{label_id}][%d] getLabelByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLabelByIDUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLabelByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelByIDNotFound creates a GetLabelByIDNotFound with default headers values
func NewGetLabelByIDNotFound() *GetLabelByIDNotFound {
	return &GetLabelByIDNotFound{}
}

/*
GetLabelByIDNotFound handles this case with default header values.

Not found
*/
type GetLabelByIDNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLabelByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /labels/{label_id}][%d] getLabelByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLabelByIDNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLabelByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelByIDInternalServerError creates a GetLabelByIDInternalServerError with default headers values
func NewGetLabelByIDInternalServerError() *GetLabelByIDInternalServerError {
	return &GetLabelByIDInternalServerError{}
}

/*
GetLabelByIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetLabelByIDInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetLabelByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /labels/{label_id}][%d] getLabelByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLabelByIDInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetLabelByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
