package error

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	CustomCode int `json:"code,omitempty"`

	*echo.HTTPError
}

func (e *Error) Error() string {
	return fmt.Sprintf("CustomCode=%d, %s", e.CustomCode, e.HTTPError.Error())
}

func InternalServerError(err error) error {
	return &Error{
		CustomCode: -50011,
		HTTPError: &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		},
	}
}

func BadRequest(err error) error {
	return &Error{
		CustomCode: -40011,
		HTTPError: &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  err.Error(),
			Internal: err,
		},
	}
}
