// Code generated by go-swagger; DO NOT EDIT.

package robot

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

// NewUpdateRobotParams creates a new UpdateRobotParams object
// with the default values initialized.
func NewUpdateRobotParams() *UpdateRobotParams {
	var ()
	return &UpdateRobotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRobotParamsWithTimeout creates a new UpdateRobotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRobotParamsWithTimeout(timeout time.Duration) *UpdateRobotParams {
	var ()
	return &UpdateRobotParams{

		timeout: timeout,
	}
}

// NewUpdateRobotParamsWithContext creates a new UpdateRobotParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRobotParamsWithContext(ctx context.Context) *UpdateRobotParams {
	var ()
	return &UpdateRobotParams{

		Context: ctx,
	}
}

// NewUpdateRobotParamsWithHTTPClient creates a new UpdateRobotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRobotParamsWithHTTPClient(client *http.Client) *UpdateRobotParams {
	var ()
	return &UpdateRobotParams{
		HTTPClient: client,
	}
}

/*
UpdateRobotParams contains all the parameters to send to the API endpoint
for the update robot operation typically these are written to a http.Request
*/
type UpdateRobotParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Robot
	  The JSON object of a robot account.

	*/
	Robot *model.Robot
	/*RobotID
	  Robot ID

	*/
	RobotID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update robot params
func (o *UpdateRobotParams) WithTimeout(timeout time.Duration) *UpdateRobotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update robot params
func (o *UpdateRobotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update robot params
func (o *UpdateRobotParams) WithContext(ctx context.Context) *UpdateRobotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update robot params
func (o *UpdateRobotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update robot params
func (o *UpdateRobotParams) WithHTTPClient(client *http.Client) *UpdateRobotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update robot params
func (o *UpdateRobotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update robot params
func (o *UpdateRobotParams) WithXRequestID(xRequestID *string) *UpdateRobotParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update robot params
func (o *UpdateRobotParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRobot adds the robot to the update robot params
func (o *UpdateRobotParams) WithRobot(robot *model.Robot) *UpdateRobotParams {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the update robot params
func (o *UpdateRobotParams) SetRobot(robot *model.Robot) {
	o.Robot = robot
}

// WithRobotID adds the robotID to the update robot params
func (o *UpdateRobotParams) WithRobotID(robotID int64) *UpdateRobotParams {
	o.SetRobotID(robotID)
	return o
}

// SetRobotID adds the robotId to the update robot params
func (o *UpdateRobotParams) SetRobotID(robotID int64) {
	o.RobotID = robotID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRobotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
			return err
		}
	}

	// path param robot_id
	if err := r.SetPathParam("robot_id", swag.FormatInt64(o.RobotID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
