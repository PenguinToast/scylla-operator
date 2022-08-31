// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/mermaidclient/internal/models"
)

// NewGetClusterClusterIDTasksRepairTargetParams creates a new GetClusterClusterIDTasksRepairTargetParams object
// with the default values initialized.
func NewGetClusterClusterIDTasksRepairTargetParams() *GetClusterClusterIDTasksRepairTargetParams {
	var ()
	return &GetClusterClusterIDTasksRepairTargetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDTasksRepairTargetParamsWithTimeout creates a new GetClusterClusterIDTasksRepairTargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDTasksRepairTargetParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDTasksRepairTargetParams {
	var ()
	return &GetClusterClusterIDTasksRepairTargetParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDTasksRepairTargetParamsWithContext creates a new GetClusterClusterIDTasksRepairTargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDTasksRepairTargetParamsWithContext(ctx context.Context) *GetClusterClusterIDTasksRepairTargetParams {
	var ()
	return &GetClusterClusterIDTasksRepairTargetParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDTasksRepairTargetParamsWithHTTPClient creates a new GetClusterClusterIDTasksRepairTargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDTasksRepairTargetParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDTasksRepairTargetParams {
	var ()
	return &GetClusterClusterIDTasksRepairTargetParams{
		HTTPClient: client,
	}
}

/*
GetClusterClusterIDTasksRepairTargetParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID tasks repair target operation typically these are written to a http.Request
*/
type GetClusterClusterIDTasksRepairTargetParams struct {

	/*ClusterID*/
	ClusterID string
	/*TaskFields*/
	TaskFields *models.TaskUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDTasksRepairTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) WithContext(ctx context.Context) *GetClusterClusterIDTasksRepairTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDTasksRepairTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) WithClusterID(clusterID string) *GetClusterClusterIDTasksRepairTargetParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithTaskFields adds the taskFields to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) WithTaskFields(taskFields *models.TaskUpdate) *GetClusterClusterIDTasksRepairTargetParams {
	o.SetTaskFields(taskFields)
	return o
}

// SetTaskFields adds the taskFields to the get cluster cluster ID tasks repair target params
func (o *GetClusterClusterIDTasksRepairTargetParams) SetTaskFields(taskFields *models.TaskUpdate) {
	o.TaskFields = taskFields
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDTasksRepairTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.TaskFields != nil {
		if err := r.SetBodyParam(o.TaskFields); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
