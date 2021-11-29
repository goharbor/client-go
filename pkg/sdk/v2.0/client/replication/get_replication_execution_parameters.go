// Code generated by go-swagger; DO NOT EDIT.

package replication

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
)

// NewGetReplicationExecutionParams creates a new GetReplicationExecutionParams object
// with the default values initialized.
func NewGetReplicationExecutionParams() *GetReplicationExecutionParams {
	var ()
	return &GetReplicationExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationExecutionParamsWithTimeout creates a new GetReplicationExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReplicationExecutionParamsWithTimeout(timeout time.Duration) *GetReplicationExecutionParams {
	var ()
	return &GetReplicationExecutionParams{

		timeout: timeout,
	}
}

// NewGetReplicationExecutionParamsWithContext creates a new GetReplicationExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReplicationExecutionParamsWithContext(ctx context.Context) *GetReplicationExecutionParams {
	var ()
	return &GetReplicationExecutionParams{

		Context: ctx,
	}
}

// NewGetReplicationExecutionParamsWithHTTPClient creates a new GetReplicationExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReplicationExecutionParamsWithHTTPClient(client *http.Client) *GetReplicationExecutionParams {
	var ()
	return &GetReplicationExecutionParams{
		HTTPClient: client,
	}
}

/*GetReplicationExecutionParams contains all the parameters to send to the API endpoint
for the get replication execution operation typically these are written to a http.Request
*/
type GetReplicationExecutionParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ID
	  The ID of the execution.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get replication execution params
func (o *GetReplicationExecutionParams) WithTimeout(timeout time.Duration) *GetReplicationExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication execution params
func (o *GetReplicationExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication execution params
func (o *GetReplicationExecutionParams) WithContext(ctx context.Context) *GetReplicationExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication execution params
func (o *GetReplicationExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication execution params
func (o *GetReplicationExecutionParams) WithHTTPClient(client *http.Client) *GetReplicationExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication execution params
func (o *GetReplicationExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get replication execution params
func (o *GetReplicationExecutionParams) WithXRequestID(xRequestID *string) *GetReplicationExecutionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get replication execution params
func (o *GetReplicationExecutionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get replication execution params
func (o *GetReplicationExecutionParams) WithID(id int64) *GetReplicationExecutionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get replication execution params
func (o *GetReplicationExecutionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
