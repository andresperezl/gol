// Code generated by go-swagger; DO NOT EDIT.

package champion

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

// NewGetChampionInfoParams creates a new GetChampionInfoParams object
// with the default values initialized.
func NewGetChampionInfoParams() *GetChampionInfoParams {

	return &GetChampionInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChampionInfoParamsWithTimeout creates a new GetChampionInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChampionInfoParamsWithTimeout(timeout time.Duration) *GetChampionInfoParams {

	return &GetChampionInfoParams{

		timeout: timeout,
	}
}

// NewGetChampionInfoParamsWithContext creates a new GetChampionInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChampionInfoParamsWithContext(ctx context.Context) *GetChampionInfoParams {

	return &GetChampionInfoParams{

		Context: ctx,
	}
}

// NewGetChampionInfoParamsWithHTTPClient creates a new GetChampionInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChampionInfoParamsWithHTTPClient(client *http.Client) *GetChampionInfoParams {

	return &GetChampionInfoParams{
		HTTPClient: client,
	}
}

/*GetChampionInfoParams contains all the parameters to send to the API endpoint
for the get champion info operation typically these are written to a http.Request
*/
type GetChampionInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get champion info params
func (o *GetChampionInfoParams) WithTimeout(timeout time.Duration) *GetChampionInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get champion info params
func (o *GetChampionInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get champion info params
func (o *GetChampionInfoParams) WithContext(ctx context.Context) *GetChampionInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get champion info params
func (o *GetChampionInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get champion info params
func (o *GetChampionInfoParams) WithHTTPClient(client *http.Client) *GetChampionInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get champion info params
func (o *GetChampionInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetChampionInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
