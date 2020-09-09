// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/apiv2/model/legacy"
)

// GetProjectsProjectIDWebhookPoliciesPolicyIDReader is a Reader for the GetProjectsProjectIDWebhookPoliciesPolicyID structure.
type GetProjectsProjectIDWebhookPoliciesPolicyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDOK creates a GetProjectsProjectIDWebhookPoliciesPolicyIDOK with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDOK() *GetProjectsProjectIDWebhookPoliciesPolicyIDOK {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDOK{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDOK handles this case with default header values.

Get webhook policy successfully.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDOK struct {
	Payload *legacy.WebhookPolicy
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdOK  %+v", 200, o.Payload)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDOK) GetPayload() *legacy.WebhookPolicy {
	return o.Payload
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.WebhookPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest creates a GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest() *GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest handles this case with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdBadRequest ", 400)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized creates a GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized() *GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdUnauthorized ", 401)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDForbidden creates a GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDForbidden() *GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden handles this case with default header values.

User have no permission to get webhook policy of the project.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdForbidden ", 403)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDNotFound creates a GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDNotFound() *GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound handles this case with default header values.

Webhook policy ID does not exist.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdNotFound ", 404)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError creates a GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError with default headers values
func NewGetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError() *GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError {
	return &GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError{}
}

/*GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError handles this case with default header values.

Internal server errors.
*/
type GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError struct {
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/policies/{policy_id}][%d] getProjectsProjectIdWebhookPoliciesPolicyIdInternalServerError ", 500)
}

func (o *GetProjectsProjectIDWebhookPoliciesPolicyIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}