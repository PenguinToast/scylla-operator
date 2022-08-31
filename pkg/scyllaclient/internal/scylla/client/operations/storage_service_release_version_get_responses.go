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

// StorageServiceReleaseVersionGetReader is a Reader for the StorageServiceReleaseVersionGet structure.
type StorageServiceReleaseVersionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceReleaseVersionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceReleaseVersionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceReleaseVersionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceReleaseVersionGetOK creates a StorageServiceReleaseVersionGetOK with default headers values
func NewStorageServiceReleaseVersionGetOK() *StorageServiceReleaseVersionGetOK {
	return &StorageServiceReleaseVersionGetOK{}
}

/*
StorageServiceReleaseVersionGetOK handles this case with default header values.

StorageServiceReleaseVersionGetOK storage service release version get o k
*/
type StorageServiceReleaseVersionGetOK struct {
	Payload string
}

func (o *StorageServiceReleaseVersionGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceReleaseVersionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceReleaseVersionGetDefault creates a StorageServiceReleaseVersionGetDefault with default headers values
func NewStorageServiceReleaseVersionGetDefault(code int) *StorageServiceReleaseVersionGetDefault {
	return &StorageServiceReleaseVersionGetDefault{
		_statusCode: code,
	}
}

/*
StorageServiceReleaseVersionGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceReleaseVersionGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service release version get default response
func (o *StorageServiceReleaseVersionGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceReleaseVersionGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceReleaseVersionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceReleaseVersionGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
