// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetTenantsTenantIDMetricsAPIV1SeriesReader is a Reader for the GetTenantsTenantIDMetricsAPIV1Series structure.
type GetTenantsTenantIDMetricsAPIV1SeriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTenantIDMetricsAPIV1SeriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTenantIDMetricsAPIV1SeriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTenantIDMetricsAPIV1SeriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTenantIDMetricsAPIV1SeriesOK creates a GetTenantsTenantIDMetricsAPIV1SeriesOK with default headers values
func NewGetTenantsTenantIDMetricsAPIV1SeriesOK() *GetTenantsTenantIDMetricsAPIV1SeriesOK {
	return &GetTenantsTenantIDMetricsAPIV1SeriesOK{}
}

/*GetTenantsTenantIDMetricsAPIV1SeriesOK handles this case with default header values.

List of metric names
*/
type GetTenantsTenantIDMetricsAPIV1SeriesOK struct {
	Payload []models.PrometheusLabelset
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/api/v1/series][%d] getTenantsTenantIdMetricsApiV1SeriesOK  %+v", 200, o.Payload)
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesOK) GetPayload() []models.PrometheusLabelset {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTenantIDMetricsAPIV1SeriesDefault creates a GetTenantsTenantIDMetricsAPIV1SeriesDefault with default headers values
func NewGetTenantsTenantIDMetricsAPIV1SeriesDefault(code int) *GetTenantsTenantIDMetricsAPIV1SeriesDefault {
	return &GetTenantsTenantIDMetricsAPIV1SeriesDefault{
		_statusCode: code,
	}
}

/*GetTenantsTenantIDMetricsAPIV1SeriesDefault handles this case with default header values.

Unexpected Error
*/
type GetTenantsTenantIDMetricsAPIV1SeriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants tenant ID metrics API v1 series default response
func (o *GetTenantsTenantIDMetricsAPIV1SeriesDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/api/v1/series][%d] GetTenantsTenantIDMetricsAPIV1Series default  %+v", o._statusCode, o.Payload)
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsAPIV1SeriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
