// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetInstanceReader is a Reader for the GetInstance structure.
type GetInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceOK creates a GetInstanceOK with default headers values
func NewGetInstanceOK() *GetInstanceOK {
	return &GetInstanceOK{}
}

/*GetInstanceOK handles this case with default header values.

Success
*/
type GetInstanceOK struct {
	Payload *models.Instance
}

func (o *GetInstanceOK) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceOK  %+v", 200, o.Payload)
}

func (o *GetInstanceOK) GetPayload() *models.Instance {
	return o.Payload
}

func (o *GetInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceBadRequest creates a GetInstanceBadRequest with default headers values
func NewGetInstanceBadRequest() *GetInstanceBadRequest {
	return &GetInstanceBadRequest{}
}

/*GetInstanceBadRequest handles this case with default header values.

Bad request
*/
type GetInstanceBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetInstanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstanceBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceUnauthorized creates a GetInstanceUnauthorized with default headers values
func NewGetInstanceUnauthorized() *GetInstanceUnauthorized {
	return &GetInstanceUnauthorized{}
}

/*GetInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type GetInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstanceUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceForbidden creates a GetInstanceForbidden with default headers values
func NewGetInstanceForbidden() *GetInstanceForbidden {
	return &GetInstanceForbidden{}
}

/*GetInstanceForbidden handles this case with default header values.

Forbidden
*/
type GetInstanceForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetInstanceForbidden) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceForbidden  %+v", 403, o.Payload)
}

func (o *GetInstanceForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceNotFound creates a GetInstanceNotFound with default headers values
func NewGetInstanceNotFound() *GetInstanceNotFound {
	return &GetInstanceNotFound{}
}

/*GetInstanceNotFound handles this case with default header values.

Not found
*/
type GetInstanceNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetInstanceNotFound) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceInternalServerError creates a GetInstanceInternalServerError with default headers values
func NewGetInstanceInternalServerError() *GetInstanceInternalServerError {
	return &GetInstanceInternalServerError{}
}

/*GetInstanceInternalServerError handles this case with default header values.

Internal server error
*/
type GetInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances/{preheat_instance_name}][%d] getInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstanceInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
