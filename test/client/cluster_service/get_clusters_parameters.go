// Code generated by go-swagger; DO NOT EDIT.

package cluster_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClustersParams creates a new GetClustersParams object
// with the default values initialized.
func NewGetClustersParams() *GetClustersParams {
	var ()
	return &GetClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClustersParamsWithTimeout creates a new GetClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClustersParamsWithTimeout(timeout time.Duration) *GetClustersParams {
	var ()
	return &GetClustersParams{

		timeout: timeout,
	}
}

// NewGetClustersParamsWithContext creates a new GetClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClustersParamsWithContext(ctx context.Context) *GetClustersParams {
	var ()
	return &GetClustersParams{

		Context: ctx,
	}
}

// NewGetClustersParamsWithHTTPClient creates a new GetClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClustersParamsWithHTTPClient(client *http.Client) *GetClustersParams {
	var ()
	return &GetClustersParams{
		HTTPClient: client,
	}
}

/*GetClustersParams contains all the parameters to send to the API endpoint
for the get clusters operation typically these are written to a http.Request
*/
type GetClustersParams struct {

	/*Ids*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) WithTimeout(timeout time.Duration) *GetClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clusters params
func (o *GetClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clusters params
func (o *GetClustersParams) WithContext(ctx context.Context) *GetClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clusters params
func (o *GetClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) WithHTTPClient(client *http.Client) *GetClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clusters params
func (o *GetClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get clusters params
func (o *GetClustersParams) WithIds(ids string) *GetClustersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get clusters params
func (o *GetClustersParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ids
	if err := r.SetPathParam("ids", o.Ids); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
