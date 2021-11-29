// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// ListRobotReader is a Reader for the ListRobot structure.
type ListRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRobotOK creates a ListRobotOK with default headers values
func NewListRobotOK() *ListRobotOK {
	return &ListRobotOK{}
}

/*ListRobotOK handles this case with default header values.

Success
*/
type ListRobotOK struct {
	/*Link refers to the previous page and next page
	 */
	Link string
	/*The total count of robot accounts
	 */
	XTotalCount int64

	Payload []*models.Robot
}

func (o *ListRobotOK) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotOK  %+v", 200, o.Payload)
}

func (o *ListRobotOK) GetPayload() []*models.Robot {
	return o.Payload
}

func (o *ListRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotBadRequest creates a ListRobotBadRequest with default headers values
func NewListRobotBadRequest() *ListRobotBadRequest {
	return &ListRobotBadRequest{}
}

/*ListRobotBadRequest handles this case with default header values.

Bad request
*/
type ListRobotBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotBadRequest) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotBadRequest  %+v", 400, o.Payload)
}

func (o *ListRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotNotFound creates a ListRobotNotFound with default headers values
func NewListRobotNotFound() *ListRobotNotFound {
	return &ListRobotNotFound{}
}

/*ListRobotNotFound handles this case with default header values.

Not found
*/
type ListRobotNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotNotFound) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotNotFound  %+v", 404, o.Payload)
}

func (o *ListRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotInternalServerError creates a ListRobotInternalServerError with default headers values
func NewListRobotInternalServerError() *ListRobotInternalServerError {
	return &ListRobotInternalServerError{}
}

/*ListRobotInternalServerError handles this case with default header values.

Internal server error
*/
type ListRobotInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotInternalServerError) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
