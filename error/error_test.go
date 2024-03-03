package error

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestError(t *testing.T) {
	testcases := []struct {
		name       string
		code       int
		customCode int
		err        error
		expect     string
	}{
		{
			"with custom code",
			http.StatusInternalServerError,
			0,
			errors.New("with custom code error"),
			"custom_code=0, code=500, message=with custom code error, internal=with custom code error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			err := Error{
				CustomCode: tt.customCode,
				HTTPError: &echo.HTTPError{
					Code:     tt.code,
					Message:  tt.err.Error(),
					Internal: tt.err,
				},
			}
			out := err.Error()
			if out != tt.expect {
				t.Errorf("want %s, got %s", tt.expect, out)
			}
		})
	}
}
