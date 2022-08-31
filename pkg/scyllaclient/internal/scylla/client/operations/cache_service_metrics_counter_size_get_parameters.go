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

// NewCacheServiceMetricsCounterSizeGetParams creates a new CacheServiceMetricsCounterSizeGetParams object
// with the default values initialized.
func NewCacheServiceMetricsCounterSizeGetParams() *CacheServiceMetricsCounterSizeGetParams {

	return &CacheServiceMetricsCounterSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsCounterSizeGetParamsWithTimeout creates a new CacheServiceMetricsCounterSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsCounterSizeGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsCounterSizeGetParams {

	return &CacheServiceMetricsCounterSizeGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsCounterSizeGetParamsWithContext creates a new CacheServiceMetricsCounterSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsCounterSizeGetParamsWithContext(ctx context.Context) *CacheServiceMetricsCounterSizeGetParams {

	return &CacheServiceMetricsCounterSizeGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsCounterSizeGetParamsWithHTTPClient creates a new CacheServiceMetricsCounterSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsCounterSizeGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsCounterSizeGetParams {

	return &CacheServiceMetricsCounterSizeGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsCounterSizeGetParams contains all the parameters to send to the API endpoint
for the cache service metrics counter size get operation typically these are written to a http.Request
*/
type CacheServiceMetricsCounterSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsCounterSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) WithContext(ctx context.Context) *CacheServiceMetricsCounterSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsCounterSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics counter size get params
func (o *CacheServiceMetricsCounterSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsCounterSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
