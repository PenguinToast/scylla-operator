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

// FindConfigVolatileSystemKeyspaceForTestingReader is a Reader for the FindConfigVolatileSystemKeyspaceForTesting structure.
type FindConfigVolatileSystemKeyspaceForTestingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigVolatileSystemKeyspaceForTestingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigVolatileSystemKeyspaceForTestingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigVolatileSystemKeyspaceForTestingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigVolatileSystemKeyspaceForTestingOK creates a FindConfigVolatileSystemKeyspaceForTestingOK with default headers values
func NewFindConfigVolatileSystemKeyspaceForTestingOK() *FindConfigVolatileSystemKeyspaceForTestingOK {
	return &FindConfigVolatileSystemKeyspaceForTestingOK{}
}

/*
FindConfigVolatileSystemKeyspaceForTestingOK handles this case with default header values.

Config value
*/
type FindConfigVolatileSystemKeyspaceForTestingOK struct {
	Payload bool
}

func (o *FindConfigVolatileSystemKeyspaceForTestingOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigVolatileSystemKeyspaceForTestingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigVolatileSystemKeyspaceForTestingDefault creates a FindConfigVolatileSystemKeyspaceForTestingDefault with default headers values
func NewFindConfigVolatileSystemKeyspaceForTestingDefault(code int) *FindConfigVolatileSystemKeyspaceForTestingDefault {
	return &FindConfigVolatileSystemKeyspaceForTestingDefault{
		_statusCode: code,
	}
}

/*
FindConfigVolatileSystemKeyspaceForTestingDefault handles this case with default header values.

unexpected error
*/
type FindConfigVolatileSystemKeyspaceForTestingDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config volatile system keyspace for testing default response
func (o *FindConfigVolatileSystemKeyspaceForTestingDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigVolatileSystemKeyspaceForTestingDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigVolatileSystemKeyspaceForTestingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigVolatileSystemKeyspaceForTestingDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
