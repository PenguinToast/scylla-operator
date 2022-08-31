// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/mermaidclient/internal/models"
)

// GetClusterClusterIDTaskTaskTypeTaskIDReader is a Reader for the GetClusterClusterIDTaskTaskTypeTaskID structure.
type GetClusterClusterIDTaskTaskTypeTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTaskTaskTypeTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDOK creates a GetClusterClusterIDTaskTaskTypeTaskIDOK with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDOK() *GetClusterClusterIDTaskTaskTypeTaskIDOK {
	return &GetClusterClusterIDTaskTaskTypeTaskIDOK{}
}

/*
GetClusterClusterIDTaskTaskTypeTaskIDOK handles this case with default header values.

Task info
*/
type GetClusterClusterIDTaskTaskTypeTaskIDOK struct {
	Payload *models.Task
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] getClusterClusterIdTaskTaskTypeTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDDefault creates a GetClusterClusterIDTaskTaskTypeTaskIDDefault with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDDefault(code int) *GetClusterClusterIDTaskTaskTypeTaskIDDefault {
	return &GetClusterClusterIDTaskTaskTypeTaskIDDefault{
		_statusCode: code,
	}
}

/*
GetClusterClusterIDTaskTaskTypeTaskIDDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDTaskTaskTypeTaskIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID task task type task ID default response
func (o *GetClusterClusterIDTaskTaskTypeTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] GetClusterClusterIDTaskTaskTypeTaskID default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
