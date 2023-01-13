// Code generated by go-swagger; DO NOT EDIT.

package gc

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

// NewUpdateGCScheduleParams creates a new UpdateGCScheduleParams object
// with the default values initialized.
func NewUpdateGCScheduleParams() *UpdateGCScheduleParams {
	var ()
	return &UpdateGCScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGCScheduleParamsWithTimeout creates a new UpdateGCScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGCScheduleParamsWithTimeout(timeout time.Duration) *UpdateGCScheduleParams {
	var ()
	return &UpdateGCScheduleParams{

		timeout: timeout,
	}
}

// NewUpdateGCScheduleParamsWithContext creates a new UpdateGCScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGCScheduleParamsWithContext(ctx context.Context) *UpdateGCScheduleParams {
	var ()
	return &UpdateGCScheduleParams{

		Context: ctx,
	}
}

// NewUpdateGCScheduleParamsWithHTTPClient creates a new UpdateGCScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGCScheduleParamsWithHTTPClient(client *http.Client) *UpdateGCScheduleParams {
	var ()
	return &UpdateGCScheduleParams{
		HTTPClient: client,
	}
}

/*
UpdateGCScheduleParams contains all the parameters to send to the API endpoint
for the update GC schedule operation typically these are written to a http.Request
*/
type UpdateGCScheduleParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Schedule
	  Updates of gc's schedule.

	*/
	Schedule *model.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update GC schedule params
func (o *UpdateGCScheduleParams) WithTimeout(timeout time.Duration) *UpdateGCScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update GC schedule params
func (o *UpdateGCScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update GC schedule params
func (o *UpdateGCScheduleParams) WithContext(ctx context.Context) *UpdateGCScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update GC schedule params
func (o *UpdateGCScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update GC schedule params
func (o *UpdateGCScheduleParams) WithHTTPClient(client *http.Client) *UpdateGCScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update GC schedule params
func (o *UpdateGCScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update GC schedule params
func (o *UpdateGCScheduleParams) WithXRequestID(xRequestID *string) *UpdateGCScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update GC schedule params
func (o *UpdateGCScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the update GC schedule params
func (o *UpdateGCScheduleParams) WithSchedule(schedule *model.Schedule) *UpdateGCScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the update GC schedule params
func (o *UpdateGCScheduleParams) SetSchedule(schedule *model.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGCScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
