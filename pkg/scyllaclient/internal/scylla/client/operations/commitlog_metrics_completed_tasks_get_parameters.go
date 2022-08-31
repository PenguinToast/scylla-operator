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
)

// NewCommitlogMetricsCompletedTasksGetParams creates a new CommitlogMetricsCompletedTasksGetParams object
// with the default values initialized.
func NewCommitlogMetricsCompletedTasksGetParams() *CommitlogMetricsCompletedTasksGetParams {

	return &CommitlogMetricsCompletedTasksGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommitlogMetricsCompletedTasksGetParamsWithTimeout creates a new CommitlogMetricsCompletedTasksGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommitlogMetricsCompletedTasksGetParamsWithTimeout(timeout time.Duration) *CommitlogMetricsCompletedTasksGetParams {

	return &CommitlogMetricsCompletedTasksGetParams{

		timeout: timeout,
	}
}

// NewCommitlogMetricsCompletedTasksGetParamsWithContext creates a new CommitlogMetricsCompletedTasksGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommitlogMetricsCompletedTasksGetParamsWithContext(ctx context.Context) *CommitlogMetricsCompletedTasksGetParams {

	return &CommitlogMetricsCompletedTasksGetParams{

		Context: ctx,
	}
}

// NewCommitlogMetricsCompletedTasksGetParamsWithHTTPClient creates a new CommitlogMetricsCompletedTasksGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommitlogMetricsCompletedTasksGetParamsWithHTTPClient(client *http.Client) *CommitlogMetricsCompletedTasksGetParams {

	return &CommitlogMetricsCompletedTasksGetParams{
		HTTPClient: client,
	}
}

/*
CommitlogMetricsCompletedTasksGetParams contains all the parameters to send to the API endpoint
for the commitlog metrics completed tasks get operation typically these are written to a http.Request
*/
type CommitlogMetricsCompletedTasksGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) WithTimeout(timeout time.Duration) *CommitlogMetricsCompletedTasksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) WithContext(ctx context.Context) *CommitlogMetricsCompletedTasksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) WithHTTPClient(client *http.Client) *CommitlogMetricsCompletedTasksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commitlog metrics completed tasks get params
func (o *CommitlogMetricsCompletedTasksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CommitlogMetricsCompletedTasksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
