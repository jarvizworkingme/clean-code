package business_exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBusinessError(error, message string) *gin.Error {
	if message == "" {
		message = errorMessages[error]
	}

	errorHttpStatus := errorHttpStatus[error]
	if errorHttpStatus == 0 {
		errorHttpStatus = http.StatusUnprocessableEntity
	}

	return &gin.Error{
		Err:  fmt.Errorf("http status %d error: %s message: %s", errorHttpStatus, error, message),
		Type: gin.ErrorTypePublic,
		Meta: Error{
			HttpStatus: errorHttpStatus,
			Code:       error,
			Message:    message,
		},
	}
}

type Error struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return e.Code
}
