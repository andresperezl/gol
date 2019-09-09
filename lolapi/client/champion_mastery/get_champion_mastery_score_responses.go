// Code generated by go-swagger; DO NOT EDIT.

package champion_mastery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/andresperezl/gol/lolapi/models"
)

// GetChampionMasteryScoreReader is a Reader for the GetChampionMasteryScore structure.
type GetChampionMasteryScoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChampionMasteryScoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChampionMasteryScoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetChampionMasteryScoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetChampionMasteryScoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetChampionMasteryScoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetChampionMasteryScoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetChampionMasteryScoreMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetChampionMasteryScoreUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetChampionMasteryScoreTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetChampionMasteryScoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetChampionMasteryScoreBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetChampionMasteryScoreServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetChampionMasteryScoreGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChampionMasteryScoreOK creates a GetChampionMasteryScoreOK with default headers values
func NewGetChampionMasteryScoreOK() *GetChampionMasteryScoreOK {
	return &GetChampionMasteryScoreOK{}
}

/*GetChampionMasteryScoreOK handles this case with default header values.

OK
*/
type GetChampionMasteryScoreOK struct {
	Payload int32
}

func (o *GetChampionMasteryScoreOK) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreOK  %+v", 200, o.Payload)
}

func (o *GetChampionMasteryScoreOK) GetPayload() int32 {
	return o.Payload
}

func (o *GetChampionMasteryScoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreBadRequest creates a GetChampionMasteryScoreBadRequest with default headers values
func NewGetChampionMasteryScoreBadRequest() *GetChampionMasteryScoreBadRequest {
	return &GetChampionMasteryScoreBadRequest{}
}

/*GetChampionMasteryScoreBadRequest handles this case with default header values.

Bad request
*/
type GetChampionMasteryScoreBadRequest struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreBadRequest) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreBadRequest  %+v", 400, o.Payload)
}

func (o *GetChampionMasteryScoreBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreUnauthorized creates a GetChampionMasteryScoreUnauthorized with default headers values
func NewGetChampionMasteryScoreUnauthorized() *GetChampionMasteryScoreUnauthorized {
	return &GetChampionMasteryScoreUnauthorized{}
}

/*GetChampionMasteryScoreUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChampionMasteryScoreUnauthorized struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreUnauthorized  %+v", 401, o.Payload)
}

func (o *GetChampionMasteryScoreUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreForbidden creates a GetChampionMasteryScoreForbidden with default headers values
func NewGetChampionMasteryScoreForbidden() *GetChampionMasteryScoreForbidden {
	return &GetChampionMasteryScoreForbidden{}
}

/*GetChampionMasteryScoreForbidden handles this case with default header values.

Forbidden
*/
type GetChampionMasteryScoreForbidden struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreForbidden) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreForbidden  %+v", 403, o.Payload)
}

func (o *GetChampionMasteryScoreForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreNotFound creates a GetChampionMasteryScoreNotFound with default headers values
func NewGetChampionMasteryScoreNotFound() *GetChampionMasteryScoreNotFound {
	return &GetChampionMasteryScoreNotFound{}
}

/*GetChampionMasteryScoreNotFound handles this case with default header values.

Not found
*/
type GetChampionMasteryScoreNotFound struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreNotFound) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreNotFound  %+v", 404, o.Payload)
}

func (o *GetChampionMasteryScoreNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreMethodNotAllowed creates a GetChampionMasteryScoreMethodNotAllowed with default headers values
func NewGetChampionMasteryScoreMethodNotAllowed() *GetChampionMasteryScoreMethodNotAllowed {
	return &GetChampionMasteryScoreMethodNotAllowed{}
}

/*GetChampionMasteryScoreMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetChampionMasteryScoreMethodNotAllowed struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetChampionMasteryScoreMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreUnsupportedMediaType creates a GetChampionMasteryScoreUnsupportedMediaType with default headers values
func NewGetChampionMasteryScoreUnsupportedMediaType() *GetChampionMasteryScoreUnsupportedMediaType {
	return &GetChampionMasteryScoreUnsupportedMediaType{}
}

/*GetChampionMasteryScoreUnsupportedMediaType handles this case with default header values.

Unsopported media type
*/
type GetChampionMasteryScoreUnsupportedMediaType struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetChampionMasteryScoreUnsupportedMediaType) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreTooManyRequests creates a GetChampionMasteryScoreTooManyRequests with default headers values
func NewGetChampionMasteryScoreTooManyRequests() *GetChampionMasteryScoreTooManyRequests {
	return &GetChampionMasteryScoreTooManyRequests{}
}

/*GetChampionMasteryScoreTooManyRequests handles this case with default header values.

Rate limit Exceeded
*/
type GetChampionMasteryScoreTooManyRequests struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetChampionMasteryScoreTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreInternalServerError creates a GetChampionMasteryScoreInternalServerError with default headers values
func NewGetChampionMasteryScoreInternalServerError() *GetChampionMasteryScoreInternalServerError {
	return &GetChampionMasteryScoreInternalServerError{}
}

/*GetChampionMasteryScoreInternalServerError handles this case with default header values.

Internal server error
*/
type GetChampionMasteryScoreInternalServerError struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreInternalServerError  %+v", 500, o.Payload)
}

func (o *GetChampionMasteryScoreInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreBadGateway creates a GetChampionMasteryScoreBadGateway with default headers values
func NewGetChampionMasteryScoreBadGateway() *GetChampionMasteryScoreBadGateway {
	return &GetChampionMasteryScoreBadGateway{}
}

/*GetChampionMasteryScoreBadGateway handles this case with default header values.

Bad gateway
*/
type GetChampionMasteryScoreBadGateway struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreBadGateway) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreBadGateway  %+v", 502, o.Payload)
}

func (o *GetChampionMasteryScoreBadGateway) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreServiceUnavailable creates a GetChampionMasteryScoreServiceUnavailable with default headers values
func NewGetChampionMasteryScoreServiceUnavailable() *GetChampionMasteryScoreServiceUnavailable {
	return &GetChampionMasteryScoreServiceUnavailable{}
}

/*GetChampionMasteryScoreServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetChampionMasteryScoreServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetChampionMasteryScoreServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChampionMasteryScoreGatewayTimeout creates a GetChampionMasteryScoreGatewayTimeout with default headers values
func NewGetChampionMasteryScoreGatewayTimeout() *GetChampionMasteryScoreGatewayTimeout {
	return &GetChampionMasteryScoreGatewayTimeout{}
}

/*GetChampionMasteryScoreGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetChampionMasteryScoreGatewayTimeout struct {
	Payload *models.Error
}

func (o *GetChampionMasteryScoreGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}][%d] getChampionMasteryScoreGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetChampionMasteryScoreGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetChampionMasteryScoreGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}