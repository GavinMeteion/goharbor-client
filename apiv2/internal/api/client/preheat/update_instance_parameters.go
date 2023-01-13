// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

	"github.com/GavinMeteion/goharbor-client/v5/apiv2/model"
)

// NewUpdateInstanceParams creates a new UpdateInstanceParams object
// with the default values initialized.
func NewUpdateInstanceParams() *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInstanceParamsWithTimeout creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInstanceParamsWithTimeout(timeout time.Duration) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		timeout: timeout,
	}
}

// NewUpdateInstanceParamsWithContext creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInstanceParamsWithContext(ctx context.Context) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		Context: ctx,
	}
}

// NewUpdateInstanceParamsWithHTTPClient creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInstanceParamsWithHTTPClient(client *http.Client) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{
		HTTPClient: client,
	}
}

/*
UpdateInstanceParams contains all the parameters to send to the API endpoint
for the update instance operation typically these are written to a http.Request
*/
type UpdateInstanceParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Instance
	  The instance to update

	*/
	Instance *model.Instance
	/*PreheatInstanceName
	  Instance Name

	*/
	PreheatInstanceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update instance params
func (o *UpdateInstanceParams) WithTimeout(timeout time.Duration) *UpdateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update instance params
func (o *UpdateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update instance params
func (o *UpdateInstanceParams) WithContext(ctx context.Context) *UpdateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update instance params
func (o *UpdateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update instance params
func (o *UpdateInstanceParams) WithHTTPClient(client *http.Client) *UpdateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update instance params
func (o *UpdateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update instance params
func (o *UpdateInstanceParams) WithXRequestID(xRequestID *string) *UpdateInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update instance params
func (o *UpdateInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithInstance adds the instance to the update instance params
func (o *UpdateInstanceParams) WithInstance(instance *model.Instance) *UpdateInstanceParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the update instance params
func (o *UpdateInstanceParams) SetInstance(instance *model.Instance) {
	o.Instance = instance
}

// WithPreheatInstanceName adds the preheatInstanceName to the update instance params
func (o *UpdateInstanceParams) WithPreheatInstanceName(preheatInstanceName string) *UpdateInstanceParams {
	o.SetPreheatInstanceName(preheatInstanceName)
	return o
}

// SetPreheatInstanceName adds the preheatInstanceName to the update instance params
func (o *UpdateInstanceParams) SetPreheatInstanceName(preheatInstanceName string) {
	o.PreheatInstanceName = preheatInstanceName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Instance != nil {
		if err := r.SetBodyParam(o.Instance); err != nil {
			return err
		}
	}

	// path param preheat_instance_name
	if err := r.SetPathParam("preheat_instance_name", o.PreheatInstanceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
