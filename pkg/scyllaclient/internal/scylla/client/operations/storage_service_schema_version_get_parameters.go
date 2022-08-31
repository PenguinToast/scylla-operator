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

// NewStorageServiceSchemaVersionGetParams creates a new StorageServiceSchemaVersionGetParams object
// with the default values initialized.
func NewStorageServiceSchemaVersionGetParams() *StorageServiceSchemaVersionGetParams {

	return &StorageServiceSchemaVersionGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceSchemaVersionGetParamsWithTimeout creates a new StorageServiceSchemaVersionGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceSchemaVersionGetParamsWithTimeout(timeout time.Duration) *StorageServiceSchemaVersionGetParams {

	return &StorageServiceSchemaVersionGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceSchemaVersionGetParamsWithContext creates a new StorageServiceSchemaVersionGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceSchemaVersionGetParamsWithContext(ctx context.Context) *StorageServiceSchemaVersionGetParams {

	return &StorageServiceSchemaVersionGetParams{

		Context: ctx,
	}
}

// NewStorageServiceSchemaVersionGetParamsWithHTTPClient creates a new StorageServiceSchemaVersionGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceSchemaVersionGetParamsWithHTTPClient(client *http.Client) *StorageServiceSchemaVersionGetParams {

	return &StorageServiceSchemaVersionGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceSchemaVersionGetParams contains all the parameters to send to the API endpoint
for the storage service schema version get operation typically these are written to a http.Request
*/
type StorageServiceSchemaVersionGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) WithTimeout(timeout time.Duration) *StorageServiceSchemaVersionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) WithContext(ctx context.Context) *StorageServiceSchemaVersionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) WithHTTPClient(client *http.Client) *StorageServiceSchemaVersionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service schema version get params
func (o *StorageServiceSchemaVersionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceSchemaVersionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
