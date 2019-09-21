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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLeagueEntriesExpParams creates a new GetLeagueEntriesExpParams object
// with the default values initialized.
func NewGetLeagueEntriesExpParams() *GetLeagueEntriesExpParams {
	var ()
	return &GetLeagueEntriesExpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLeagueEntriesExpParamsWithTimeout creates a new GetLeagueEntriesExpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLeagueEntriesExpParamsWithTimeout(timeout time.Duration) *GetLeagueEntriesExpParams {
	var ()
	return &GetLeagueEntriesExpParams{

		timeout: timeout,
	}
}

// NewGetLeagueEntriesExpParamsWithContext creates a new GetLeagueEntriesExpParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLeagueEntriesExpParamsWithContext(ctx context.Context) *GetLeagueEntriesExpParams {
	var ()
	return &GetLeagueEntriesExpParams{

		Context: ctx,
	}
}

// NewGetLeagueEntriesExpParamsWithHTTPClient creates a new GetLeagueEntriesExpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLeagueEntriesExpParamsWithHTTPClient(client *http.Client) *GetLeagueEntriesExpParams {
	var ()
	return &GetLeagueEntriesExpParams{
		HTTPClient: client,
	}
}

/*GetLeagueEntriesExpParams contains all the parameters to send to the API endpoint
for the get league entries exp operation typically these are written to a http.Request
*/
type GetLeagueEntriesExpParams struct {

	/*Division*/
	Division string
	/*Page*/
	Page *int32
	/*Queue
	  Note that the queue value must be a valid ranked queue.

	*/
	Queue string
	/*Tier*/
	Tier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithTimeout(timeout time.Duration) *GetLeagueEntriesExpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithContext(ctx context.Context) *GetLeagueEntriesExpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithHTTPClient(client *http.Client) *GetLeagueEntriesExpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDivision adds the division to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithDivision(division string) *GetLeagueEntriesExpParams {
	o.SetDivision(division)
	return o
}

// SetDivision adds the division to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetDivision(division string) {
	o.Division = division
}

// WithPage adds the page to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithPage(page *int32) *GetLeagueEntriesExpParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetPage(page *int32) {
	o.Page = page
}

// WithQueue adds the queue to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithQueue(queue string) *GetLeagueEntriesExpParams {
	o.SetQueue(queue)
	return o
}

// SetQueue adds the queue to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetQueue(queue string) {
	o.Queue = queue
}

// WithTier adds the tier to the get league entries exp params
func (o *GetLeagueEntriesExpParams) WithTier(tier string) *GetLeagueEntriesExpParams {
	o.SetTier(tier)
	return o
}

// SetTier adds the tier to the get league entries exp params
func (o *GetLeagueEntriesExpParams) SetTier(tier string) {
	o.Tier = tier
}

// WriteToRequest writes these params to a swagger request
func (o *GetLeagueEntriesExpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param division
	if err := r.SetPathParam("division", o.Division); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	// path param queue
	if err := r.SetPathParam("queue", o.Queue); err != nil {
		return err
	}

	// path param tier
	if err := r.SetPathParam("tier", o.Tier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}