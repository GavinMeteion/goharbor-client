// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// DeleteRobotV1Reader is a Reader for the DeleteRobotV1 structure.
type DeleteRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRobotV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRobotV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRobotV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRobotV1OK creates a DeleteRobotV1OK with default headers values
func NewDeleteRobotV1OK() *DeleteRobotV1OK {
	return &DeleteRobotV1OK{}
}

/*
DeleteRobotV1OK handles this case with default header values.

Success
*/
type DeleteRobotV1OK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteRobotV1OK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1OK ", 200)
}

func (o *DeleteRobotV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteRobotV1BadRequest creates a DeleteRobotV1BadRequest with default headers values
func NewDeleteRobotV1BadRequest() *DeleteRobotV1BadRequest {
	return &DeleteRobotV1BadRequest{}
}

/*
DeleteRobotV1BadRequest handles this case with default header values.

Bad request
*/
type DeleteRobotV1BadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRobotV1BadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRobotV1Unauthorized creates a DeleteRobotV1Unauthorized with default headers values
func NewDeleteRobotV1Unauthorized() *DeleteRobotV1Unauthorized {
	return &DeleteRobotV1Unauthorized{}
}

/*
DeleteRobotV1Unauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRobotV1Unauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRobotV1Unauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRobotV1Forbidden creates a DeleteRobotV1Forbidden with default headers values
func NewDeleteRobotV1Forbidden() *DeleteRobotV1Forbidden {
	return &DeleteRobotV1Forbidden{}
}

/*
DeleteRobotV1Forbidden handles this case with default header values.

Forbidden
*/
type DeleteRobotV1Forbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRobotV1Forbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRobotV1NotFound creates a DeleteRobotV1NotFound with default headers values
func NewDeleteRobotV1NotFound() *DeleteRobotV1NotFound {
	return &DeleteRobotV1NotFound{}
}

/*
DeleteRobotV1NotFound handles this case with default header values.

Not found
*/
type DeleteRobotV1NotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRobotV1NotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRobotV1InternalServerError creates a DeleteRobotV1InternalServerError with default headers values
func NewDeleteRobotV1InternalServerError() *DeleteRobotV1InternalServerError {
	return &DeleteRobotV1InternalServerError{}
}

/*
DeleteRobotV1InternalServerError handles this case with default header values.

Internal server error
*/
type DeleteRobotV1InternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRobotV1InternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
