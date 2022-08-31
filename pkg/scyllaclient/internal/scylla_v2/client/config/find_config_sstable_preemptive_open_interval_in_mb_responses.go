// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigSstablePreemptiveOpenIntervalInMbReader is a Reader for the FindConfigSstablePreemptiveOpenIntervalInMb structure.
type FindConfigSstablePreemptiveOpenIntervalInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigSstablePreemptiveOpenIntervalInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigSstablePreemptiveOpenIntervalInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigSstablePreemptiveOpenIntervalInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigSstablePreemptiveOpenIntervalInMbOK creates a FindConfigSstablePreemptiveOpenIntervalInMbOK with default headers values
func NewFindConfigSstablePreemptiveOpenIntervalInMbOK() *FindConfigSstablePreemptiveOpenIntervalInMbOK {
	return &FindConfigSstablePreemptiveOpenIntervalInMbOK{}
}

/*
FindConfigSstablePreemptiveOpenIntervalInMbOK handles this case with default header values.

Config value
*/
type FindConfigSstablePreemptiveOpenIntervalInMbOK struct {
	Payload int64
}

func (o *FindConfigSstablePreemptiveOpenIntervalInMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigSstablePreemptiveOpenIntervalInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigSstablePreemptiveOpenIntervalInMbDefault creates a FindConfigSstablePreemptiveOpenIntervalInMbDefault with default headers values
func NewFindConfigSstablePreemptiveOpenIntervalInMbDefault(code int) *FindConfigSstablePreemptiveOpenIntervalInMbDefault {
	return &FindConfigSstablePreemptiveOpenIntervalInMbDefault{
		_statusCode: code,
	}
}

/*
FindConfigSstablePreemptiveOpenIntervalInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigSstablePreemptiveOpenIntervalInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config sstable preemptive open interval in mb default response
func (o *FindConfigSstablePreemptiveOpenIntervalInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigSstablePreemptiveOpenIntervalInMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigSstablePreemptiveOpenIntervalInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigSstablePreemptiveOpenIntervalInMbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
