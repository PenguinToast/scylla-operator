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

// FindConfigCounterCacheKeysToSaveReader is a Reader for the FindConfigCounterCacheKeysToSave structure.
type FindConfigCounterCacheKeysToSaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCounterCacheKeysToSaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCounterCacheKeysToSaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCounterCacheKeysToSaveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCounterCacheKeysToSaveOK creates a FindConfigCounterCacheKeysToSaveOK with default headers values
func NewFindConfigCounterCacheKeysToSaveOK() *FindConfigCounterCacheKeysToSaveOK {
	return &FindConfigCounterCacheKeysToSaveOK{}
}

/*
FindConfigCounterCacheKeysToSaveOK handles this case with default header values.

Config value
*/
type FindConfigCounterCacheKeysToSaveOK struct {
	Payload int64
}

func (o *FindConfigCounterCacheKeysToSaveOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCounterCacheKeysToSaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCounterCacheKeysToSaveDefault creates a FindConfigCounterCacheKeysToSaveDefault with default headers values
func NewFindConfigCounterCacheKeysToSaveDefault(code int) *FindConfigCounterCacheKeysToSaveDefault {
	return &FindConfigCounterCacheKeysToSaveDefault{
		_statusCode: code,
	}
}

/*
FindConfigCounterCacheKeysToSaveDefault handles this case with default header values.

unexpected error
*/
type FindConfigCounterCacheKeysToSaveDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config counter cache keys to save default response
func (o *FindConfigCounterCacheKeysToSaveDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCounterCacheKeysToSaveDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCounterCacheKeysToSaveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCounterCacheKeysToSaveDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
