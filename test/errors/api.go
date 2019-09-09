package errors

import (
	"fmt"
	"github.com/andresperezl/gol/lolapi/models"
	"github.com/go-openapi/errors"
	"net/http"
	"reflect"
)

type Error interface {
	error
	Code() int32
}

type riotError struct {
	status riotStatus
}

type riotStatus struct {
	statusCode int32
	message    string
}

func (r *riotError) Error() string {
	return r.status.message
}

func (r *riotError) Code() int32 {
	return r.status.statusCode
}

func New(code int32, message string, args ...interface{}) errors.Error {
	if len(args) > 0 {
		return &riotError{
			status: riotStatus{
				statusCode: code,
				message:    fmt.Sprintf(message, args...),
			},
		}
	}
	return &riotError{
		status: riotStatus{
			statusCode: code,
			message:    message,
		},
	}
}

// ServeError the error handler interface implementation
func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case Error:
		value := reflect.ValueOf(e)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			errors.ServeError(rw, r, err)
			return
		}
		rw.WriteHeader(asHTTPCode(int(e.Code())))
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(e))
		}
	default:
		errors.ServeError(rw, r, err)
	}
}

func errorAsJSON(err Error) []byte {
	mError := &models.Error{
		Status: &models.Status{
			StatusCode: err.Code(),
			Message:    err.Error(),
		},
	}
	b, _ := mError.MarshalBinary()
	return b
}

func asHTTPCode(input int) int {
	if input >= 600 {
		return errors.DefaultHTTPCode
	}
	return input
}
