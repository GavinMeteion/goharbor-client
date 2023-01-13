// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetCurrentUserPermissionsReader is a Reader for the GetCurrentUserPermissions structure.
type GetCurrentUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserPermissionsOK creates a GetCurrentUserPermissionsOK with default headers values
func NewGetCurrentUserPermissionsOK() *GetCurrentUserPermissionsOK {
	return &GetCurrentUserPermissionsOK{}
}

/*
GetCurrentUserPermissionsOK handles this case with default header values.

Get current user permission successfully.
*/
type GetCurrentUserPermissionsOK struct {
	Payload []*model.Permission
}

func (o *GetCurrentUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getCurrentUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserPermissionsOK) GetPayload() []*model.Permission {
	return o.Payload
}

func (o *GetCurrentUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserPermissionsUnauthorized creates a GetCurrentUserPermissionsUnauthorized with default headers values
func NewGetCurrentUserPermissionsUnauthorized() *GetCurrentUserPermissionsUnauthorized {
	return &GetCurrentUserPermissionsUnauthorized{}
}

/*
GetCurrentUserPermissionsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetCurrentUserPermissionsUnauthorized struct {
}

func (o *GetCurrentUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getCurrentUserPermissionsUnauthorized ", 401)
}

func (o *GetCurrentUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentUserPermissionsInternalServerError creates a GetCurrentUserPermissionsInternalServerError with default headers values
func NewGetCurrentUserPermissionsInternalServerError() *GetCurrentUserPermissionsInternalServerError {
	return &GetCurrentUserPermissionsInternalServerError{}
}

/*
GetCurrentUserPermissionsInternalServerError handles this case with default header values.

Internal errors.
*/
type GetCurrentUserPermissionsInternalServerError struct {
}

func (o *GetCurrentUserPermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/permissions][%d] getCurrentUserPermissionsInternalServerError ", 500)
}

func (o *GetCurrentUserPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
