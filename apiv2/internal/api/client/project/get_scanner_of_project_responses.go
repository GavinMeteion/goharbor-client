// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// GetScannerOfProjectReader is a Reader for the GetScannerOfProject structure.
type GetScannerOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScannerOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScannerOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScannerOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScannerOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScannerOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScannerOfProjectOK creates a GetScannerOfProjectOK with default headers values
func NewGetScannerOfProjectOK() *GetScannerOfProjectOK {
	return &GetScannerOfProjectOK{}
}

/*
GetScannerOfProjectOK handles this case with default header values.

The details of the scanner registration.
*/
type GetScannerOfProjectOK struct {
	Payload *model.ScannerRegistration
}

func (o *GetScannerOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetScannerOfProjectOK) GetPayload() *model.ScannerRegistration {
	return o.Payload
}

func (o *GetScannerOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ScannerRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScannerOfProjectBadRequest creates a GetScannerOfProjectBadRequest with default headers values
func NewGetScannerOfProjectBadRequest() *GetScannerOfProjectBadRequest {
	return &GetScannerOfProjectBadRequest{}
}

/*
GetScannerOfProjectBadRequest handles this case with default header values.

Bad project ID
*/
type GetScannerOfProjectBadRequest struct {
}

func (o *GetScannerOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectBadRequest ", 400)
}

func (o *GetScannerOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectUnauthorized creates a GetScannerOfProjectUnauthorized with default headers values
func NewGetScannerOfProjectUnauthorized() *GetScannerOfProjectUnauthorized {
	return &GetScannerOfProjectUnauthorized{}
}

/*
GetScannerOfProjectUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetScannerOfProjectUnauthorized struct {
}

func (o *GetScannerOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectUnauthorized ", 401)
}

func (o *GetScannerOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectForbidden creates a GetScannerOfProjectForbidden with default headers values
func NewGetScannerOfProjectForbidden() *GetScannerOfProjectForbidden {
	return &GetScannerOfProjectForbidden{}
}

/*
GetScannerOfProjectForbidden handles this case with default header values.

Request is not allowed
*/
type GetScannerOfProjectForbidden struct {
}

func (o *GetScannerOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectForbidden ", 403)
}

func (o *GetScannerOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectNotFound creates a GetScannerOfProjectNotFound with default headers values
func NewGetScannerOfProjectNotFound() *GetScannerOfProjectNotFound {
	return &GetScannerOfProjectNotFound{}
}

/*
GetScannerOfProjectNotFound handles this case with default header values.

The requested object is not found
*/
type GetScannerOfProjectNotFound struct {
}

func (o *GetScannerOfProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectNotFound ", 404)
}

func (o *GetScannerOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScannerOfProjectInternalServerError creates a GetScannerOfProjectInternalServerError with default headers values
func NewGetScannerOfProjectInternalServerError() *GetScannerOfProjectInternalServerError {
	return &GetScannerOfProjectInternalServerError{}
}

/*
GetScannerOfProjectInternalServerError handles this case with default header values.

Internal server error happened
*/
type GetScannerOfProjectInternalServerError struct {
}

func (o *GetScannerOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/scanner][%d] getScannerOfProjectInternalServerError ", 500)
}

func (o *GetScannerOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
