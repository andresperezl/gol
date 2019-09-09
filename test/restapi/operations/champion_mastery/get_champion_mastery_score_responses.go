// Code generated by go-swagger; DO NOT EDIT.

package champion_mastery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/andresperezl/gol/lolapi/models"
)

// GetChampionMasteryScoreOKCode is the HTTP code returned for type GetChampionMasteryScoreOK
const GetChampionMasteryScoreOKCode int = 200

/*GetChampionMasteryScoreOK OK

swagger:response getChampionMasteryScoreOK
*/
type GetChampionMasteryScoreOK struct {

	/*
	  In: Body
	*/
	Payload int32 `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreOK creates GetChampionMasteryScoreOK with default headers values
func NewGetChampionMasteryScoreOK() *GetChampionMasteryScoreOK {

	return &GetChampionMasteryScoreOK{}
}

// WithPayload adds the payload to the get champion mastery score o k response
func (o *GetChampionMasteryScoreOK) WithPayload(payload int32) *GetChampionMasteryScoreOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score o k response
func (o *GetChampionMasteryScoreOK) SetPayload(payload int32) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetChampionMasteryScoreBadRequestCode is the HTTP code returned for type GetChampionMasteryScoreBadRequest
const GetChampionMasteryScoreBadRequestCode int = 400

/*GetChampionMasteryScoreBadRequest Bad request

swagger:response getChampionMasteryScoreBadRequest
*/
type GetChampionMasteryScoreBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreBadRequest creates GetChampionMasteryScoreBadRequest with default headers values
func NewGetChampionMasteryScoreBadRequest() *GetChampionMasteryScoreBadRequest {

	return &GetChampionMasteryScoreBadRequest{}
}

// WithPayload adds the payload to the get champion mastery score bad request response
func (o *GetChampionMasteryScoreBadRequest) WithPayload(payload *models.Error) *GetChampionMasteryScoreBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score bad request response
func (o *GetChampionMasteryScoreBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreUnauthorizedCode is the HTTP code returned for type GetChampionMasteryScoreUnauthorized
const GetChampionMasteryScoreUnauthorizedCode int = 401

/*GetChampionMasteryScoreUnauthorized Unauthorized

swagger:response getChampionMasteryScoreUnauthorized
*/
type GetChampionMasteryScoreUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreUnauthorized creates GetChampionMasteryScoreUnauthorized with default headers values
func NewGetChampionMasteryScoreUnauthorized() *GetChampionMasteryScoreUnauthorized {

	return &GetChampionMasteryScoreUnauthorized{}
}

// WithPayload adds the payload to the get champion mastery score unauthorized response
func (o *GetChampionMasteryScoreUnauthorized) WithPayload(payload *models.Error) *GetChampionMasteryScoreUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score unauthorized response
func (o *GetChampionMasteryScoreUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreForbiddenCode is the HTTP code returned for type GetChampionMasteryScoreForbidden
const GetChampionMasteryScoreForbiddenCode int = 403

/*GetChampionMasteryScoreForbidden Forbidden

swagger:response getChampionMasteryScoreForbidden
*/
type GetChampionMasteryScoreForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreForbidden creates GetChampionMasteryScoreForbidden with default headers values
func NewGetChampionMasteryScoreForbidden() *GetChampionMasteryScoreForbidden {

	return &GetChampionMasteryScoreForbidden{}
}

// WithPayload adds the payload to the get champion mastery score forbidden response
func (o *GetChampionMasteryScoreForbidden) WithPayload(payload *models.Error) *GetChampionMasteryScoreForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score forbidden response
func (o *GetChampionMasteryScoreForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreNotFoundCode is the HTTP code returned for type GetChampionMasteryScoreNotFound
const GetChampionMasteryScoreNotFoundCode int = 404

/*GetChampionMasteryScoreNotFound Not found

swagger:response getChampionMasteryScoreNotFound
*/
type GetChampionMasteryScoreNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreNotFound creates GetChampionMasteryScoreNotFound with default headers values
func NewGetChampionMasteryScoreNotFound() *GetChampionMasteryScoreNotFound {

	return &GetChampionMasteryScoreNotFound{}
}

// WithPayload adds the payload to the get champion mastery score not found response
func (o *GetChampionMasteryScoreNotFound) WithPayload(payload *models.Error) *GetChampionMasteryScoreNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score not found response
func (o *GetChampionMasteryScoreNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreMethodNotAllowedCode is the HTTP code returned for type GetChampionMasteryScoreMethodNotAllowed
const GetChampionMasteryScoreMethodNotAllowedCode int = 405

/*GetChampionMasteryScoreMethodNotAllowed Method not allowed

swagger:response getChampionMasteryScoreMethodNotAllowed
*/
type GetChampionMasteryScoreMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreMethodNotAllowed creates GetChampionMasteryScoreMethodNotAllowed with default headers values
func NewGetChampionMasteryScoreMethodNotAllowed() *GetChampionMasteryScoreMethodNotAllowed {

	return &GetChampionMasteryScoreMethodNotAllowed{}
}

// WithPayload adds the payload to the get champion mastery score method not allowed response
func (o *GetChampionMasteryScoreMethodNotAllowed) WithPayload(payload *models.Error) *GetChampionMasteryScoreMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score method not allowed response
func (o *GetChampionMasteryScoreMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreUnsupportedMediaTypeCode is the HTTP code returned for type GetChampionMasteryScoreUnsupportedMediaType
const GetChampionMasteryScoreUnsupportedMediaTypeCode int = 415

/*GetChampionMasteryScoreUnsupportedMediaType Unsopported media type

swagger:response getChampionMasteryScoreUnsupportedMediaType
*/
type GetChampionMasteryScoreUnsupportedMediaType struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreUnsupportedMediaType creates GetChampionMasteryScoreUnsupportedMediaType with default headers values
func NewGetChampionMasteryScoreUnsupportedMediaType() *GetChampionMasteryScoreUnsupportedMediaType {

	return &GetChampionMasteryScoreUnsupportedMediaType{}
}

// WithPayload adds the payload to the get champion mastery score unsupported media type response
func (o *GetChampionMasteryScoreUnsupportedMediaType) WithPayload(payload *models.Error) *GetChampionMasteryScoreUnsupportedMediaType {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score unsupported media type response
func (o *GetChampionMasteryScoreUnsupportedMediaType) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreUnsupportedMediaType) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(415)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreTooManyRequestsCode is the HTTP code returned for type GetChampionMasteryScoreTooManyRequests
const GetChampionMasteryScoreTooManyRequestsCode int = 429

/*GetChampionMasteryScoreTooManyRequests Rate limit Exceeded

swagger:response getChampionMasteryScoreTooManyRequests
*/
type GetChampionMasteryScoreTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreTooManyRequests creates GetChampionMasteryScoreTooManyRequests with default headers values
func NewGetChampionMasteryScoreTooManyRequests() *GetChampionMasteryScoreTooManyRequests {

	return &GetChampionMasteryScoreTooManyRequests{}
}

// WithPayload adds the payload to the get champion mastery score too many requests response
func (o *GetChampionMasteryScoreTooManyRequests) WithPayload(payload *models.Error) *GetChampionMasteryScoreTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score too many requests response
func (o *GetChampionMasteryScoreTooManyRequests) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreInternalServerErrorCode is the HTTP code returned for type GetChampionMasteryScoreInternalServerError
const GetChampionMasteryScoreInternalServerErrorCode int = 500

/*GetChampionMasteryScoreInternalServerError Internal server error

swagger:response getChampionMasteryScoreInternalServerError
*/
type GetChampionMasteryScoreInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreInternalServerError creates GetChampionMasteryScoreInternalServerError with default headers values
func NewGetChampionMasteryScoreInternalServerError() *GetChampionMasteryScoreInternalServerError {

	return &GetChampionMasteryScoreInternalServerError{}
}

// WithPayload adds the payload to the get champion mastery score internal server error response
func (o *GetChampionMasteryScoreInternalServerError) WithPayload(payload *models.Error) *GetChampionMasteryScoreInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score internal server error response
func (o *GetChampionMasteryScoreInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreBadGatewayCode is the HTTP code returned for type GetChampionMasteryScoreBadGateway
const GetChampionMasteryScoreBadGatewayCode int = 502

/*GetChampionMasteryScoreBadGateway Bad gateway

swagger:response getChampionMasteryScoreBadGateway
*/
type GetChampionMasteryScoreBadGateway struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreBadGateway creates GetChampionMasteryScoreBadGateway with default headers values
func NewGetChampionMasteryScoreBadGateway() *GetChampionMasteryScoreBadGateway {

	return &GetChampionMasteryScoreBadGateway{}
}

// WithPayload adds the payload to the get champion mastery score bad gateway response
func (o *GetChampionMasteryScoreBadGateway) WithPayload(payload *models.Error) *GetChampionMasteryScoreBadGateway {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score bad gateway response
func (o *GetChampionMasteryScoreBadGateway) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreServiceUnavailableCode is the HTTP code returned for type GetChampionMasteryScoreServiceUnavailable
const GetChampionMasteryScoreServiceUnavailableCode int = 503

/*GetChampionMasteryScoreServiceUnavailable Service Unavailable

swagger:response getChampionMasteryScoreServiceUnavailable
*/
type GetChampionMasteryScoreServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreServiceUnavailable creates GetChampionMasteryScoreServiceUnavailable with default headers values
func NewGetChampionMasteryScoreServiceUnavailable() *GetChampionMasteryScoreServiceUnavailable {

	return &GetChampionMasteryScoreServiceUnavailable{}
}

// WithPayload adds the payload to the get champion mastery score service unavailable response
func (o *GetChampionMasteryScoreServiceUnavailable) WithPayload(payload *models.Error) *GetChampionMasteryScoreServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score service unavailable response
func (o *GetChampionMasteryScoreServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChampionMasteryScoreGatewayTimeoutCode is the HTTP code returned for type GetChampionMasteryScoreGatewayTimeout
const GetChampionMasteryScoreGatewayTimeoutCode int = 504

/*GetChampionMasteryScoreGatewayTimeout Gateway timeout

swagger:response getChampionMasteryScoreGatewayTimeout
*/
type GetChampionMasteryScoreGatewayTimeout struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChampionMasteryScoreGatewayTimeout creates GetChampionMasteryScoreGatewayTimeout with default headers values
func NewGetChampionMasteryScoreGatewayTimeout() *GetChampionMasteryScoreGatewayTimeout {

	return &GetChampionMasteryScoreGatewayTimeout{}
}

// WithPayload adds the payload to the get champion mastery score gateway timeout response
func (o *GetChampionMasteryScoreGatewayTimeout) WithPayload(payload *models.Error) *GetChampionMasteryScoreGatewayTimeout {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get champion mastery score gateway timeout response
func (o *GetChampionMasteryScoreGatewayTimeout) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChampionMasteryScoreGatewayTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(504)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
