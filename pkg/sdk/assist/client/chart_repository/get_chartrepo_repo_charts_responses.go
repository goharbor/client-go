// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/assist/models"
)

// GetChartrepoRepoChartsReader is a Reader for the GetChartrepoRepoCharts structure.
type GetChartrepoRepoChartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChartrepoRepoChartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChartrepoRepoChartsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetChartrepoRepoChartsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetChartrepoRepoChartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetChartrepoRepoChartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChartrepoRepoChartsOK creates a GetChartrepoRepoChartsOK with default headers values
func NewGetChartrepoRepoChartsOK() *GetChartrepoRepoChartsOK {
	return &GetChartrepoRepoChartsOK{}
}

/*GetChartrepoRepoChartsOK handles this case with default header values.

Searched for charts of project in Harbor successfully.
*/
type GetChartrepoRepoChartsOK struct {
	Payload []*models.ChartInfoEntry
}

func (o *GetChartrepoRepoChartsOK) Error() string {
	return fmt.Sprintf("[GET /chartrepo/{repo}/charts][%d] getChartrepoRepoChartsOK  %+v", 200, o.Payload)
}

func (o *GetChartrepoRepoChartsOK) GetPayload() []*models.ChartInfoEntry {
	return o.Payload
}

func (o *GetChartrepoRepoChartsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartrepoRepoChartsUnauthorized creates a GetChartrepoRepoChartsUnauthorized with default headers values
func NewGetChartrepoRepoChartsUnauthorized() *GetChartrepoRepoChartsUnauthorized {
	return &GetChartrepoRepoChartsUnauthorized{}
}

/*GetChartrepoRepoChartsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChartrepoRepoChartsUnauthorized struct {
}

func (o *GetChartrepoRepoChartsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /chartrepo/{repo}/charts][%d] getChartrepoRepoChartsUnauthorized ", 401)
}

func (o *GetChartrepoRepoChartsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChartrepoRepoChartsForbidden creates a GetChartrepoRepoChartsForbidden with default headers values
func NewGetChartrepoRepoChartsForbidden() *GetChartrepoRepoChartsForbidden {
	return &GetChartrepoRepoChartsForbidden{}
}

/*GetChartrepoRepoChartsForbidden handles this case with default header values.

Operation is forbidden or quota exceeded
*/
type GetChartrepoRepoChartsForbidden struct {
}

func (o *GetChartrepoRepoChartsForbidden) Error() string {
	return fmt.Sprintf("[GET /chartrepo/{repo}/charts][%d] getChartrepoRepoChartsForbidden ", 403)
}

func (o *GetChartrepoRepoChartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChartrepoRepoChartsInternalServerError creates a GetChartrepoRepoChartsInternalServerError with default headers values
func NewGetChartrepoRepoChartsInternalServerError() *GetChartrepoRepoChartsInternalServerError {
	return &GetChartrepoRepoChartsInternalServerError{}
}

/*GetChartrepoRepoChartsInternalServerError handles this case with default header values.

Internal server error occurred
*/
type GetChartrepoRepoChartsInternalServerError struct {
}

func (o *GetChartrepoRepoChartsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chartrepo/{repo}/charts][%d] getChartrepoRepoChartsInternalServerError ", 500)
}

func (o *GetChartrepoRepoChartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
