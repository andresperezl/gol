package errors

import (
	"github.com/andresperezl/gol/lolapi/models"
)

type APIError struct {
	models.APIError
}

func (ae APIError) Error() string {
	return ae.Status.Message
}

func (ae APIError) Code() int32 {
	return ae.Status.StatusCode
}

type ErrorResponse interface {
	GetPayload() *models.APIError
}
