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

// NewStorageProxyReadRepairRepairedBackgroundGetParams creates a new StorageProxyReadRepairRepairedBackgroundGetParams object
// with the default values initialized.
func NewStorageProxyReadRepairRepairedBackgroundGetParams() *StorageProxyReadRepairRepairedBackgroundGetParams {

	return &StorageProxyReadRepairRepairedBackgroundGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyReadRepairRepairedBackgroundGetParamsWithTimeout creates a new StorageProxyReadRepairRepairedBackgroundGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyReadRepairRepairedBackgroundGetParamsWithTimeout(timeout time.Duration) *StorageProxyReadRepairRepairedBackgroundGetParams {

	return &StorageProxyReadRepairRepairedBackgroundGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyReadRepairRepairedBackgroundGetParamsWithContext creates a new StorageProxyReadRepairRepairedBackgroundGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyReadRepairRepairedBackgroundGetParamsWithContext(ctx context.Context) *StorageProxyReadRepairRepairedBackgroundGetParams {

	return &StorageProxyReadRepairRepairedBackgroundGetParams{

		Context: ctx,
	}
}

// NewStorageProxyReadRepairRepairedBackgroundGetParamsWithHTTPClient creates a new StorageProxyReadRepairRepairedBackgroundGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyReadRepairRepairedBackgroundGetParamsWithHTTPClient(client *http.Client) *StorageProxyReadRepairRepairedBackgroundGetParams {

	return &StorageProxyReadRepairRepairedBackgroundGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyReadRepairRepairedBackgroundGetParams contains all the parameters to send to the API endpoint
for the storage proxy read repair repaired background get operation typically these are written to a http.Request
*/
type StorageProxyReadRepairRepairedBackgroundGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) WithTimeout(timeout time.Duration) *StorageProxyReadRepairRepairedBackgroundGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) WithContext(ctx context.Context) *StorageProxyReadRepairRepairedBackgroundGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) WithHTTPClient(client *http.Client) *StorageProxyReadRepairRepairedBackgroundGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy read repair repaired background get params
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyReadRepairRepairedBackgroundGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
