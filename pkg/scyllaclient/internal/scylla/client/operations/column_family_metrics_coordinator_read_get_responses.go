// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// ColumnFamilyMetricsCoordinatorReadGetReader is a Reader for the ColumnFamilyMetricsCoordinatorReadGet structure.
type ColumnFamilyMetricsCoordinatorReadGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsCoordinatorReadGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsCoordinatorReadGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsCoordinatorReadGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsCoordinatorReadGetOK creates a ColumnFamilyMetricsCoordinatorReadGetOK with default headers values
func NewColumnFamilyMetricsCoordinatorReadGetOK() *ColumnFamilyMetricsCoordinatorReadGetOK {
	return &ColumnFamilyMetricsCoordinatorReadGetOK{}
}

/*
ColumnFamilyMetricsCoordinatorReadGetOK handles this case with default header values.

ColumnFamilyMetricsCoordinatorReadGetOK column family metrics coordinator read get o k
*/
type ColumnFamilyMetricsCoordinatorReadGetOK struct {
}

func (o *ColumnFamilyMetricsCoordinatorReadGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsCoordinatorReadGetDefault creates a ColumnFamilyMetricsCoordinatorReadGetDefault with default headers values
func NewColumnFamilyMetricsCoordinatorReadGetDefault(code int) *ColumnFamilyMetricsCoordinatorReadGetDefault {
	return &ColumnFamilyMetricsCoordinatorReadGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsCoordinatorReadGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsCoordinatorReadGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics coordinator read get default response
func (o *ColumnFamilyMetricsCoordinatorReadGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsCoordinatorReadGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsCoordinatorReadGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsCoordinatorReadGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
