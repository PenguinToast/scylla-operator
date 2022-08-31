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

// HintedHandoffPausePostReader is a Reader for the HintedHandoffPausePost structure.
type HintedHandoffPausePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HintedHandoffPausePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHintedHandoffPausePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHintedHandoffPausePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHintedHandoffPausePostOK creates a HintedHandoffPausePostOK with default headers values
func NewHintedHandoffPausePostOK() *HintedHandoffPausePostOK {
	return &HintedHandoffPausePostOK{}
}

/*
HintedHandoffPausePostOK handles this case with default header values.

HintedHandoffPausePostOK hinted handoff pause post o k
*/
type HintedHandoffPausePostOK struct {
}

func (o *HintedHandoffPausePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHintedHandoffPausePostDefault creates a HintedHandoffPausePostDefault with default headers values
func NewHintedHandoffPausePostDefault(code int) *HintedHandoffPausePostDefault {
	return &HintedHandoffPausePostDefault{
		_statusCode: code,
	}
}

/*
HintedHandoffPausePostDefault handles this case with default header values.

internal server error
*/
type HintedHandoffPausePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the hinted handoff pause post default response
func (o *HintedHandoffPausePostDefault) Code() int {
	return o._statusCode
}

func (o *HintedHandoffPausePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *HintedHandoffPausePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *HintedHandoffPausePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
