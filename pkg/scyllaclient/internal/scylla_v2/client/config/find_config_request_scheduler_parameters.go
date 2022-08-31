// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigRequestSchedulerParams creates a new FindConfigRequestSchedulerParams object
// with the default values initialized.
func NewFindConfigRequestSchedulerParams() *FindConfigRequestSchedulerParams {

	return &FindConfigRequestSchedulerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigRequestSchedulerParamsWithTimeout creates a new FindConfigRequestSchedulerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigRequestSchedulerParamsWithTimeout(timeout time.Duration) *FindConfigRequestSchedulerParams {

	return &FindConfigRequestSchedulerParams{

		timeout: timeout,
	}
}

// NewFindConfigRequestSchedulerParamsWithContext creates a new FindConfigRequestSchedulerParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigRequestSchedulerParamsWithContext(ctx context.Context) *FindConfigRequestSchedulerParams {

	return &FindConfigRequestSchedulerParams{

		Context: ctx,
	}
}

// NewFindConfigRequestSchedulerParamsWithHTTPClient creates a new FindConfigRequestSchedulerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigRequestSchedulerParamsWithHTTPClient(client *http.Client) *FindConfigRequestSchedulerParams {

	return &FindConfigRequestSchedulerParams{
		HTTPClient: client,
	}
}

/*
FindConfigRequestSchedulerParams contains all the parameters to send to the API endpoint
for the find config request scheduler operation typically these are written to a http.Request
*/
type FindConfigRequestSchedulerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) WithTimeout(timeout time.Duration) *FindConfigRequestSchedulerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) WithContext(ctx context.Context) *FindConfigRequestSchedulerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) WithHTTPClient(client *http.Client) *FindConfigRequestSchedulerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config request scheduler params
func (o *FindConfigRequestSchedulerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigRequestSchedulerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
