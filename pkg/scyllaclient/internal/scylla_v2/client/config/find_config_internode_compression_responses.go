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

// FindConfigInternodeCompressionReader is a Reader for the FindConfigInternodeCompression structure.
type FindConfigInternodeCompressionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigInternodeCompressionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigInternodeCompressionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigInternodeCompressionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigInternodeCompressionOK creates a FindConfigInternodeCompressionOK with default headers values
func NewFindConfigInternodeCompressionOK() *FindConfigInternodeCompressionOK {
	return &FindConfigInternodeCompressionOK{}
}

/*
FindConfigInternodeCompressionOK handles this case with default header values.

Config value
*/
type FindConfigInternodeCompressionOK struct {
	Payload string
}

func (o *FindConfigInternodeCompressionOK) GetPayload() string {
	return o.Payload
}

func (o *FindConfigInternodeCompressionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigInternodeCompressionDefault creates a FindConfigInternodeCompressionDefault with default headers values
func NewFindConfigInternodeCompressionDefault(code int) *FindConfigInternodeCompressionDefault {
	return &FindConfigInternodeCompressionDefault{
		_statusCode: code,
	}
}

/*
FindConfigInternodeCompressionDefault handles this case with default header values.

unexpected error
*/
type FindConfigInternodeCompressionDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config internode compression default response
func (o *FindConfigInternodeCompressionDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigInternodeCompressionDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigInternodeCompressionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigInternodeCompressionDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
