package middleware

import (
	"context"
	"errors"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	vnderror "github.com/thienhaole92/vnd/error"
)

type AuthProvider interface {
	VerifyIDTokenAndCheckRevoked(ctx context.Context, idToken string) (*auth.Token, error)
}

func FirebaseAuth(skipper middleware.Skipper, auth AuthProvider) echo.MiddlewareFunc {
	fbauth := func(next echo.HandlerFunc) echo.HandlerFunc {
		handler := func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			req := c.Request()
			authHeader := strings.TrimSpace(req.Header.Get(echo.HeaderAuthorization))

			if len(authHeader) == 0 {
				return vnderror.Unauthorized(errors.New("an authorization header is required"))
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				return vnderror.Unauthorized(errors.New("invalid authorization token"))
			}

			token, err := auth.VerifyIDTokenAndCheckRevoked(req.Context(), bearerToken[1])
			if err != nil {
				return vnderror.Unauthorized(err)
			}

			c.Set(UserIDContextKey, token.UID)

			return next(c)
		}

		return handler
	}
	return fbauth
}
