package context

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Context interface {
	echo.Context
	RequestContext() context.Context
	RequestId() string
	UserId() (string, error)
}
