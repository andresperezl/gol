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

// GetAllChampionMasteriesReader is a Reader for the GetAllChampionMasteries structure.
type GetAllChampionMasteriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllChampionMasteriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllChampionMasteriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllChampionMasteriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllChampionMasteriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllChampionMasteriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllChampionMasteriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAllChampionMasteriesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAllChampionMasteriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAllChampionMasteriesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllChampionMasteriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetAllChampionMasteriesBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAllChampionMasteriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAllChampionMasteriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllChampionMasteriesOK creates a GetAllChampionMasteriesOK with default headers values
func NewGetAllChampionMasteriesOK() *GetAllChampionMasteriesOK {
	return &GetAllChampionMasteriesOK{}
}

/*GetAllChampionMasteriesOK handles this case with default header values.

OK
*/
type GetAllChampionMasteriesOK struct {
	Payload []*models.ChampionMastery
}

func (o *GetAllChampionMasteriesOK) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesOK  %+v", 200, o.Payload)
}

func (o *GetAllChampionMasteriesOK) GetPayload() []*models.ChampionMastery {
	return o.Payload
}

func (o *GetAllChampionMasteriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesBadRequest creates a GetAllChampionMasteriesBadRequest with default headers values
func NewGetAllChampionMasteriesBadRequest() *GetAllChampionMasteriesBadRequest {
	return &GetAllChampionMasteriesBadRequest{}
}

/*GetAllChampionMasteriesBadRequest handles this case with default header values.

Bad request
*/
type GetAllChampionMasteriesBadRequest struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllChampionMasteriesBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesUnauthorized creates a GetAllChampionMasteriesUnauthorized with default headers values
func NewGetAllChampionMasteriesUnauthorized() *GetAllChampionMasteriesUnauthorized {
	return &GetAllChampionMasteriesUnauthorized{}
}

/*GetAllChampionMasteriesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllChampionMasteriesUnauthorized struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllChampionMasteriesUnauthorized) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesForbidden creates a GetAllChampionMasteriesForbidden with default headers values
func NewGetAllChampionMasteriesForbidden() *GetAllChampionMasteriesForbidden {
	return &GetAllChampionMasteriesForbidden{}
}

/*GetAllChampionMasteriesForbidden handles this case with default header values.

Forbidden
*/
type GetAllChampionMasteriesForbidden struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesForbidden) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesForbidden  %+v", 403, o.Payload)
}

func (o *GetAllChampionMasteriesForbidden) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesNotFound creates a GetAllChampionMasteriesNotFound with default headers values
func NewGetAllChampionMasteriesNotFound() *GetAllChampionMasteriesNotFound {
	return &GetAllChampionMasteriesNotFound{}
}

/*GetAllChampionMasteriesNotFound handles this case with default header values.

Not found
*/
type GetAllChampionMasteriesNotFound struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesNotFound) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesNotFound  %+v", 404, o.Payload)
}

func (o *GetAllChampionMasteriesNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesMethodNotAllowed creates a GetAllChampionMasteriesMethodNotAllowed with default headers values
func NewGetAllChampionMasteriesMethodNotAllowed() *GetAllChampionMasteriesMethodNotAllowed {
	return &GetAllChampionMasteriesMethodNotAllowed{}
}

/*GetAllChampionMasteriesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GetAllChampionMasteriesMethodNotAllowed struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAllChampionMasteriesMethodNotAllowed) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesUnsupportedMediaType creates a GetAllChampionMasteriesUnsupportedMediaType with default headers values
func NewGetAllChampionMasteriesUnsupportedMediaType() *GetAllChampionMasteriesUnsupportedMediaType {
	return &GetAllChampionMasteriesUnsupportedMediaType{}
}

/*GetAllChampionMasteriesUnsupportedMediaType handles this case with default header values.

Unsopported media type
*/
type GetAllChampionMasteriesUnsupportedMediaType struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAllChampionMasteriesUnsupportedMediaType) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesTooManyRequests creates a GetAllChampionMasteriesTooManyRequests with default headers values
func NewGetAllChampionMasteriesTooManyRequests() *GetAllChampionMasteriesTooManyRequests {
	return &GetAllChampionMasteriesTooManyRequests{}
}

/*GetAllChampionMasteriesTooManyRequests handles this case with default header values.

Rate limit Exceeded
*/
type GetAllChampionMasteriesTooManyRequests struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAllChampionMasteriesTooManyRequests) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesInternalServerError creates a GetAllChampionMasteriesInternalServerError with default headers values
func NewGetAllChampionMasteriesInternalServerError() *GetAllChampionMasteriesInternalServerError {
	return &GetAllChampionMasteriesInternalServerError{}
}

/*GetAllChampionMasteriesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAllChampionMasteriesInternalServerError struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllChampionMasteriesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesBadGateway creates a GetAllChampionMasteriesBadGateway with default headers values
func NewGetAllChampionMasteriesBadGateway() *GetAllChampionMasteriesBadGateway {
	return &GetAllChampionMasteriesBadGateway{}
}

/*GetAllChampionMasteriesBadGateway handles this case with default header values.

Bad gateway
*/
type GetAllChampionMasteriesBadGateway struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesBadGateway) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesBadGateway  %+v", 502, o.Payload)
}

func (o *GetAllChampionMasteriesBadGateway) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesServiceUnavailable creates a GetAllChampionMasteriesServiceUnavailable with default headers values
func NewGetAllChampionMasteriesServiceUnavailable() *GetAllChampionMasteriesServiceUnavailable {
	return &GetAllChampionMasteriesServiceUnavailable{}
}

/*GetAllChampionMasteriesServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type GetAllChampionMasteriesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAllChampionMasteriesServiceUnavailable) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllChampionMasteriesGatewayTimeout creates a GetAllChampionMasteriesGatewayTimeout with default headers values
func NewGetAllChampionMasteriesGatewayTimeout() *GetAllChampionMasteriesGatewayTimeout {
	return &GetAllChampionMasteriesGatewayTimeout{}
}

/*GetAllChampionMasteriesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAllChampionMasteriesGatewayTimeout struct {
	Payload *models.APIError
}

func (o *GetAllChampionMasteriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}][%d] getAllChampionMasteriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAllChampionMasteriesGatewayTimeout) GetPayload() *models.APIError {
	return o.Payload
}

func (o *GetAllChampionMasteriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
