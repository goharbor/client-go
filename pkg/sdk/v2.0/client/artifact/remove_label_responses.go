// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// RemoveLabelReader is a Reader for the RemoveLabel structure.
type RemoveLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveLabelOK creates a RemoveLabelOK with default headers values
func NewRemoveLabelOK() *RemoveLabelOK {
	return &RemoveLabelOK{}
}

/* RemoveLabelOK describes a response with status code 200, with default header values.

Success
*/
type RemoveLabelOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *RemoveLabelOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelOK ", 200)
}

func (o *RemoveLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewRemoveLabelUnauthorized creates a RemoveLabelUnauthorized with default headers values
func NewRemoveLabelUnauthorized() *RemoveLabelUnauthorized {
	return &RemoveLabelUnauthorized{}
}

/* RemoveLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RemoveLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *RemoveLabelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelUnauthorized  %+v", 401, o.Payload)
}
func (o *RemoveLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RemoveLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLabelForbidden creates a RemoveLabelForbidden with default headers values
func NewRemoveLabelForbidden() *RemoveLabelForbidden {
	return &RemoveLabelForbidden{}
}

/* RemoveLabelForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RemoveLabelForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *RemoveLabelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelForbidden  %+v", 403, o.Payload)
}
func (o *RemoveLabelForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RemoveLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLabelNotFound creates a RemoveLabelNotFound with default headers values
func NewRemoveLabelNotFound() *RemoveLabelNotFound {
	return &RemoveLabelNotFound{}
}

/* RemoveLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type RemoveLabelNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *RemoveLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelNotFound  %+v", 404, o.Payload)
}
func (o *RemoveLabelNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RemoveLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLabelConflict creates a RemoveLabelConflict with default headers values
func NewRemoveLabelConflict() *RemoveLabelConflict {
	return &RemoveLabelConflict{}
}

/* RemoveLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type RemoveLabelConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *RemoveLabelConflict) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelConflict  %+v", 409, o.Payload)
}
func (o *RemoveLabelConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RemoveLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLabelInternalServerError creates a RemoveLabelInternalServerError with default headers values
func NewRemoveLabelInternalServerError() *RemoveLabelInternalServerError {
	return &RemoveLabelInternalServerError{}
}

/* RemoveLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RemoveLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *RemoveLabelInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id}][%d] removeLabelInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RemoveLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}