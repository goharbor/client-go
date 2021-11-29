// Code generated by go-swagger; DO NOT EDIT.

package member

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

// NewListProjectMembersParams creates a new ListProjectMembersParams object
// with the default values initialized.
func NewListProjectMembersParams() *ListProjectMembersParams {
	var (
		xIsResourceNameDefault = bool(false)
		pageDefault            = int64(1)
		pageSizeDefault        = int64(10)
	)
	return &ListProjectMembersParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectMembersParamsWithTimeout creates a new ListProjectMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProjectMembersParamsWithTimeout(timeout time.Duration) *ListProjectMembersParams {
	var (
		xIsResourceNameDefault = bool(false)
		pageDefault            = int64(1)
		pageSizeDefault        = int64(10)
	)
	return &ListProjectMembersParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListProjectMembersParamsWithContext creates a new ListProjectMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProjectMembersParamsWithContext(ctx context.Context) *ListProjectMembersParams {
	var (
		xIsResourceNameDefault = bool(false)
		pageDefault            = int64(1)
		pageSizeDefault        = int64(10)
	)
	return &ListProjectMembersParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,

		Context: ctx,
	}
}

// NewListProjectMembersParamsWithHTTPClient creates a new ListProjectMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProjectMembersParamsWithHTTPClient(client *http.Client) *ListProjectMembersParams {
	var (
		xIsResourceNameDefault = bool(false)
		pageDefault            = int64(1)
		pageSizeDefault        = int64(10)
	)
	return &ListProjectMembersParams{
		XIsResourceName: &xIsResourceNameDefault,
		Page:            &pageDefault,
		PageSize:        &pageSizeDefault,
		HTTPClient:      client,
	}
}

/*ListProjectMembersParams contains all the parameters to send to the API endpoint
for the list project members operation typically these are written to a http.Request
*/
type ListProjectMembersParams struct {

	/*XIsResourceName
	  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.

	*/
	XIsResourceName *bool
	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Entityname
	  The entity name to search.

	*/
	Entityname *string
	/*Page
	  The page number

	*/
	Page *int64
	/*PageSize
	  The size of per page

	*/
	PageSize *int64
	/*ProjectNameOrID
	  The name or id of the project

	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list project members params
func (o *ListProjectMembersParams) WithTimeout(timeout time.Duration) *ListProjectMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project members params
func (o *ListProjectMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project members params
func (o *ListProjectMembersParams) WithContext(ctx context.Context) *ListProjectMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project members params
func (o *ListProjectMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project members params
func (o *ListProjectMembersParams) WithHTTPClient(client *http.Client) *ListProjectMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project members params
func (o *ListProjectMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the list project members params
func (o *ListProjectMembersParams) WithXIsResourceName(xIsResourceName *bool) *ListProjectMembersParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the list project members params
func (o *ListProjectMembersParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the list project members params
func (o *ListProjectMembersParams) WithXRequestID(xRequestID *string) *ListProjectMembersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list project members params
func (o *ListProjectMembersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEntityname adds the entityname to the list project members params
func (o *ListProjectMembersParams) WithEntityname(entityname *string) *ListProjectMembersParams {
	o.SetEntityname(entityname)
	return o
}

// SetEntityname adds the entityname to the list project members params
func (o *ListProjectMembersParams) SetEntityname(entityname *string) {
	o.Entityname = entityname
}

// WithPage adds the page to the list project members params
func (o *ListProjectMembersParams) WithPage(page *int64) *ListProjectMembersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list project members params
func (o *ListProjectMembersParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list project members params
func (o *ListProjectMembersParams) WithPageSize(pageSize *int64) *ListProjectMembersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list project members params
func (o *ListProjectMembersParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProjectNameOrID adds the projectNameOrID to the list project members params
func (o *ListProjectMembersParams) WithProjectNameOrID(projectNameOrID string) *ListProjectMembersParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the list project members params
func (o *ListProjectMembersParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Entityname != nil {

		// query param entityname
		var qrEntityname string
		if o.Entityname != nil {
			qrEntityname = *o.Entityname
		}
		qEntityname := qrEntityname
		if qEntityname != "" {
			if err := r.SetQueryParam("entityname", qEntityname); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
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
