package errors

import (
	"github.com/andresperezl/gol/lolapi/models"
	"testing"
)

func TestAPIError(t *testing.T) {
	msg := "I'm an error"
	var code int32 = 123
	ae := &APIError{}
	ae.Status = &models.Status{
		StatusCode: code,
		Message:    msg,
	}
	if got := ae.Error(); got != msg {
		t.Errorf("Expected \"%s\" message, got \"%s\"", msg, got)
	}
	if got := ae.Code(); got != code {
		t.Errorf("Expected %d code, got %d", code, got)
	}
}
