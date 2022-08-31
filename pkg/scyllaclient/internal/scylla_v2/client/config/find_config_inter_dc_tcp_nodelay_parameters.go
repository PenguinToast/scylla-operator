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

// NewFindConfigInterDcTCPNodelayParams creates a new FindConfigInterDcTCPNodelayParams object
// with the default values initialized.
func NewFindConfigInterDcTCPNodelayParams() *FindConfigInterDcTCPNodelayParams {

	return &FindConfigInterDcTCPNodelayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigInterDcTCPNodelayParamsWithTimeout creates a new FindConfigInterDcTCPNodelayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigInterDcTCPNodelayParamsWithTimeout(timeout time.Duration) *FindConfigInterDcTCPNodelayParams {

	return &FindConfigInterDcTCPNodelayParams{

		timeout: timeout,
	}
}

// NewFindConfigInterDcTCPNodelayParamsWithContext creates a new FindConfigInterDcTCPNodelayParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigInterDcTCPNodelayParamsWithContext(ctx context.Context) *FindConfigInterDcTCPNodelayParams {

	return &FindConfigInterDcTCPNodelayParams{

		Context: ctx,
	}
}

// NewFindConfigInterDcTCPNodelayParamsWithHTTPClient creates a new FindConfigInterDcTCPNodelayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigInterDcTCPNodelayParamsWithHTTPClient(client *http.Client) *FindConfigInterDcTCPNodelayParams {

	return &FindConfigInterDcTCPNodelayParams{
		HTTPClient: client,
	}
}

/*
FindConfigInterDcTCPNodelayParams contains all the parameters to send to the API endpoint
for the find config inter dc tcp nodelay operation typically these are written to a http.Request
*/
type FindConfigInterDcTCPNodelayParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) WithTimeout(timeout time.Duration) *FindConfigInterDcTCPNodelayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) WithContext(ctx context.Context) *FindConfigInterDcTCPNodelayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) WithHTTPClient(client *http.Client) *FindConfigInterDcTCPNodelayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config inter dc tcp nodelay params
func (o *FindConfigInterDcTCPNodelayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigInterDcTCPNodelayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
