// Code generated by go-swagger; DO NOT EDIT.

package league

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

// NewGetApexLeagueParams creates a new GetApexLeagueParams object
// with the default values initialized.
func NewGetApexLeagueParams() *GetApexLeagueParams {
	var ()
	return &GetApexLeagueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApexLeagueParamsWithTimeout creates a new GetApexLeagueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApexLeagueParamsWithTimeout(timeout time.Duration) *GetApexLeagueParams {
	var ()
	return &GetApexLeagueParams{

		timeout: timeout,
	}
}

// NewGetApexLeagueParamsWithContext creates a new GetApexLeagueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApexLeagueParamsWithContext(ctx context.Context) *GetApexLeagueParams {
	var ()
	return &GetApexLeagueParams{

		Context: ctx,
	}
}

// NewGetApexLeagueParamsWithHTTPClient creates a new GetApexLeagueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApexLeagueParamsWithHTTPClient(client *http.Client) *GetApexLeagueParams {
	var ()
	return &GetApexLeagueParams{
		HTTPClient: client,
	}
}

/*GetApexLeagueParams contains all the parameters to send to the API endpoint
for the get apex league operation typically these are written to a http.Request
*/
type GetApexLeagueParams struct {

	/*ApexLeague*/
	ApexLeague string
	/*Queue
	  Note that the queue value must be a valid ranked queue.

	*/
	Queue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apex league params
func (o *GetApexLeagueParams) WithTimeout(timeout time.Duration) *GetApexLeagueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apex league params
func (o *GetApexLeagueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apex league params
func (o *GetApexLeagueParams) WithContext(ctx context.Context) *GetApexLeagueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apex league params
func (o *GetApexLeagueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apex league params
func (o *GetApexLeagueParams) WithHTTPClient(client *http.Client) *GetApexLeagueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apex league params
func (o *GetApexLeagueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApexLeague adds the apexLeague to the get apex league params
func (o *GetApexLeagueParams) WithApexLeague(apexLeague string) *GetApexLeagueParams {
	o.SetApexLeague(apexLeague)
	return o
}

// SetApexLeague adds the apexLeague to the get apex league params
func (o *GetApexLeagueParams) SetApexLeague(apexLeague string) {
	o.ApexLeague = apexLeague
}

// WithQueue adds the queue to the get apex league params
func (o *GetApexLeagueParams) WithQueue(queue string) *GetApexLeagueParams {
	o.SetQueue(queue)
	return o
}

// SetQueue adds the queue to the get apex league params
func (o *GetApexLeagueParams) SetQueue(queue string) {
	o.Queue = queue
}

// WriteToRequest writes these params to a swagger request
func (o *GetApexLeagueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apexLeague
	if err := r.SetPathParam("apexLeague", o.ApexLeague); err != nil {
		return err
	}

	// path param queue
	if err := r.SetPathParam("queue", o.Queue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
