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

// StorageProxyMetricsWriteUnavailablesGetReader is a Reader for the StorageProxyMetricsWriteUnavailablesGet structure.
type StorageProxyMetricsWriteUnavailablesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsWriteUnavailablesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsWriteUnavailablesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsWriteUnavailablesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsWriteUnavailablesGetOK creates a StorageProxyMetricsWriteUnavailablesGetOK with default headers values
func NewStorageProxyMetricsWriteUnavailablesGetOK() *StorageProxyMetricsWriteUnavailablesGetOK {
	return &StorageProxyMetricsWriteUnavailablesGetOK{}
}

/*
StorageProxyMetricsWriteUnavailablesGetOK handles this case with default header values.

StorageProxyMetricsWriteUnavailablesGetOK storage proxy metrics write unavailables get o k
*/
type StorageProxyMetricsWriteUnavailablesGetOK struct {
	Payload int32
}

func (o *StorageProxyMetricsWriteUnavailablesGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageProxyMetricsWriteUnavailablesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsWriteUnavailablesGetDefault creates a StorageProxyMetricsWriteUnavailablesGetDefault with default headers values
func NewStorageProxyMetricsWriteUnavailablesGetDefault(code int) *StorageProxyMetricsWriteUnavailablesGetDefault {
	return &StorageProxyMetricsWriteUnavailablesGetDefault{
		_statusCode: code,
	}
}

/*
StorageProxyMetricsWriteUnavailablesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsWriteUnavailablesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics write unavailables get default response
func (o *StorageProxyMetricsWriteUnavailablesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsWriteUnavailablesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsWriteUnavailablesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsWriteUnavailablesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
