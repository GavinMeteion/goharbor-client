// Code generated by go-swagger; DO NOT EDIT.

package registry

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

// NewUpdateRegistryParams creates a new UpdateRegistryParams object
// with the default values initialized.
func NewUpdateRegistryParams() *UpdateRegistryParams {
	var ()
	return &UpdateRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRegistryParamsWithTimeout creates a new UpdateRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRegistryParamsWithTimeout(timeout time.Duration) *UpdateRegistryParams {
	var ()
	return &UpdateRegistryParams{

		timeout: timeout,
	}
}

// NewUpdateRegistryParamsWithContext creates a new UpdateRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRegistryParamsWithContext(ctx context.Context) *UpdateRegistryParams {
	var ()
	return &UpdateRegistryParams{

		Context: ctx,
	}
}

// NewUpdateRegistryParamsWithHTTPClient creates a new UpdateRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRegistryParamsWithHTTPClient(client *http.Client) *UpdateRegistryParams {
	var ()
	return &UpdateRegistryParams{
		HTTPClient: client,
	}
}

/*
UpdateRegistryParams contains all the parameters to send to the API endpoint
for the update registry operation typically these are written to a http.Request
*/
type UpdateRegistryParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ID
	  The registry ID

	*/
	ID int64
	/*Registry
	  The registry

	*/
	Registry *model.RegistryUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update registry params
func (o *UpdateRegistryParams) WithTimeout(timeout time.Duration) *UpdateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update registry params
func (o *UpdateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update registry params
func (o *UpdateRegistryParams) WithContext(ctx context.Context) *UpdateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update registry params
func (o *UpdateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update registry params
func (o *UpdateRegistryParams) WithHTTPClient(client *http.Client) *UpdateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update registry params
func (o *UpdateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update registry params
func (o *UpdateRegistryParams) WithXRequestID(xRequestID *string) *UpdateRegistryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update registry params
func (o *UpdateRegistryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the update registry params
func (o *UpdateRegistryParams) WithID(id int64) *UpdateRegistryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update registry params
func (o *UpdateRegistryParams) SetID(id int64) {
	o.ID = id
}

// WithRegistry adds the registry to the update registry params
func (o *UpdateRegistryParams) WithRegistry(registry *model.RegistryUpdate) *UpdateRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the update registry params
func (o *UpdateRegistryParams) SetRegistry(registry *model.RegistryUpdate) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Registry != nil {
		if err := r.SetBodyParam(o.Registry); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
