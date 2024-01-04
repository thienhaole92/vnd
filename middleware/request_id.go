package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestID(skipper middleware.Skipper) echo.MiddlewareFunc {
	reqid := func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			req := c.Request()
			rid := strings.TrimSpace(req.Header.Get(echo.HeaderXRequestID))
			if rid == "" {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("missing header '%s'", echo.HeaderXRequestID))
			}

			if _, err := uuid.Parse(rid); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("'%s' must be uuid", echo.HeaderXRequestID))
			}

			c.Set(RequestIDContextKey, rid)

			// next service can override the response header
			return next(c)
		}

		return handler
	}

	return reqid
}
