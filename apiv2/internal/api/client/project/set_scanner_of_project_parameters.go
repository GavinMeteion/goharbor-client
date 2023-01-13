// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// NewSetScannerOfProjectParams creates a new SetScannerOfProjectParams object
// with the default values initialized.
func NewSetScannerOfProjectParams() *SetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &SetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSetScannerOfProjectParamsWithTimeout creates a new SetScannerOfProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetScannerOfProjectParamsWithTimeout(timeout time.Duration) *SetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &SetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: timeout,
	}
}

// NewSetScannerOfProjectParamsWithContext creates a new SetScannerOfProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetScannerOfProjectParamsWithContext(ctx context.Context) *SetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &SetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,

		Context: ctx,
	}
}

// NewSetScannerOfProjectParamsWithHTTPClient creates a new SetScannerOfProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetScannerOfProjectParamsWithHTTPClient(client *http.Client) *SetScannerOfProjectParams {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &SetScannerOfProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
		HTTPClient:      client,
	}
}

/*
SetScannerOfProjectParams contains all the parameters to send to the API endpoint
for the set scanner of project operation typically these are written to a http.Request
*/
type SetScannerOfProjectParams struct {

	/*XIsResourceName
	  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.

	*/
	XIsResourceName *bool
	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Payload*/
	Payload *model.ProjectScanner
	/*ProjectNameOrID
	  The name or id of the project

	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set scanner of project params
func (o *SetScannerOfProjectParams) WithTimeout(timeout time.Duration) *SetScannerOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set scanner of project params
func (o *SetScannerOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set scanner of project params
func (o *SetScannerOfProjectParams) WithContext(ctx context.Context) *SetScannerOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set scanner of project params
func (o *SetScannerOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set scanner of project params
func (o *SetScannerOfProjectParams) WithHTTPClient(client *http.Client) *SetScannerOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set scanner of project params
func (o *SetScannerOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the set scanner of project params
func (o *SetScannerOfProjectParams) WithXIsResourceName(xIsResourceName *bool) *SetScannerOfProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the set scanner of project params
func (o *SetScannerOfProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the set scanner of project params
func (o *SetScannerOfProjectParams) WithXRequestID(xRequestID *string) *SetScannerOfProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the set scanner of project params
func (o *SetScannerOfProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the set scanner of project params
func (o *SetScannerOfProjectParams) WithPayload(payload *model.ProjectScanner) *SetScannerOfProjectParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the set scanner of project params
func (o *SetScannerOfProjectParams) SetPayload(payload *model.ProjectScanner) {
	o.Payload = payload
}

// WithProjectNameOrID adds the projectNameOrID to the set scanner of project params
func (o *SetScannerOfProjectParams) WithProjectNameOrID(projectNameOrID string) *SetScannerOfProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the set scanner of project params
func (o *SetScannerOfProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *SetScannerOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
