// Code generated by go-swagger; DO NOT EDIT.

package champion_mastery

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

// NewGetAllChampionMasteriesParams creates a new GetAllChampionMasteriesParams object
// with the default values initialized.
func NewGetAllChampionMasteriesParams() *GetAllChampionMasteriesParams {
	var ()
	return &GetAllChampionMasteriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllChampionMasteriesParamsWithTimeout creates a new GetAllChampionMasteriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllChampionMasteriesParamsWithTimeout(timeout time.Duration) *GetAllChampionMasteriesParams {
	var ()
	return &GetAllChampionMasteriesParams{

		timeout: timeout,
	}
}

// NewGetAllChampionMasteriesParamsWithContext creates a new GetAllChampionMasteriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllChampionMasteriesParamsWithContext(ctx context.Context) *GetAllChampionMasteriesParams {
	var ()
	return &GetAllChampionMasteriesParams{

		Context: ctx,
	}
}

// NewGetAllChampionMasteriesParamsWithHTTPClient creates a new GetAllChampionMasteriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllChampionMasteriesParamsWithHTTPClient(client *http.Client) *GetAllChampionMasteriesParams {
	var ()
	return &GetAllChampionMasteriesParams{
		HTTPClient: client,
	}
}

/*GetAllChampionMasteriesParams contains all the parameters to send to the API endpoint
for the get all champion masteries operation typically these are written to a http.Request
*/
type GetAllChampionMasteriesParams struct {

	/*EncryptedSummonerID
	  Summoner ID associated with the player

	*/
	EncryptedSummonerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) WithTimeout(timeout time.Duration) *GetAllChampionMasteriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) WithContext(ctx context.Context) *GetAllChampionMasteriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) WithHTTPClient(client *http.Client) *GetAllChampionMasteriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncryptedSummonerID adds the encryptedSummonerID to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) WithEncryptedSummonerID(encryptedSummonerID string) *GetAllChampionMasteriesParams {
	o.SetEncryptedSummonerID(encryptedSummonerID)
	return o
}

// SetEncryptedSummonerID adds the encryptedSummonerId to the get all champion masteries params
func (o *GetAllChampionMasteriesParams) SetEncryptedSummonerID(encryptedSummonerID string) {
	o.EncryptedSummonerID = encryptedSummonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllChampionMasteriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param encryptedSummonerId
	if err := r.SetPathParam("encryptedSummonerId", o.EncryptedSummonerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
