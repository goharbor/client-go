// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ListProjectMetadatasReader is a Reader for the ListProjectMetadatas structure.
type ListProjectMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectMetadatasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectMetadatasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectMetadatasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectMetadatasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectMetadatasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProjectMetadatasOK creates a ListProjectMetadatasOK with default headers values
func NewListProjectMetadatasOK() *ListProjectMetadatasOK {
	return &ListProjectMetadatasOK{}
}

/*ListProjectMetadatasOK handles this case with default header values.

Success
*/
type ListProjectMetadatasOK struct {
	Payload map[string]string
}

func (o *ListProjectMetadatasOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasOK  %+v", 200, o.Payload)
}

func (o *ListProjectMetadatasOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *ListProjectMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasBadRequest creates a ListProjectMetadatasBadRequest with default headers values
func NewListProjectMetadatasBadRequest() *ListProjectMetadatasBadRequest {
	return &ListProjectMetadatasBadRequest{}
}

/*ListProjectMetadatasBadRequest handles this case with default header values.

Bad request
*/
type ListProjectMetadatasBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListProjectMetadatasBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectMetadatasBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasUnauthorized creates a ListProjectMetadatasUnauthorized with default headers values
func NewListProjectMetadatasUnauthorized() *ListProjectMetadatasUnauthorized {
	return &ListProjectMetadatasUnauthorized{}
}

/*ListProjectMetadatasUnauthorized handles this case with default header values.

Unauthorized
*/
type ListProjectMetadatasUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListProjectMetadatasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectMetadatasUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasForbidden creates a ListProjectMetadatasForbidden with default headers values
func NewListProjectMetadatasForbidden() *ListProjectMetadatasForbidden {
	return &ListProjectMetadatasForbidden{}
}

/*ListProjectMetadatasForbidden handles this case with default header values.

Forbidden
*/
type ListProjectMetadatasForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListProjectMetadatasForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectMetadatasForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasNotFound creates a ListProjectMetadatasNotFound with default headers values
func NewListProjectMetadatasNotFound() *ListProjectMetadatasNotFound {
	return &ListProjectMetadatasNotFound{}
}

/*ListProjectMetadatasNotFound handles this case with default header values.

Not found
*/
type ListProjectMetadatasNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListProjectMetadatasNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectMetadatasNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasInternalServerError creates a ListProjectMetadatasInternalServerError with default headers values
func NewListProjectMetadatasInternalServerError() *ListProjectMetadatasInternalServerError {
	return &ListProjectMetadatasInternalServerError{}
}

/*ListProjectMetadatasInternalServerError handles this case with default header values.

Internal server error
*/
type ListProjectMetadatasInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListProjectMetadatasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectMetadatasInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
