package internal

import (
	"context"
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/middleware"
)

type Context struct {
	echo.Context
}

func (c *Context) RequestContext() context.Context {
	return c.Request().Context()
}

func (c *Context) RequestId() string {
	id := c.Get(middleware.RequestIDContextKey)
	if id != nil && reflect.TypeOf(id).Name() == "string" {
		return id.(string)
	}

	xid := c.Request().Header.Get(echo.HeaderXRequestID)
	if len(xid) > 0 {
		return xid
	}

	return ""
}

func (c *Context) UserId() (string, error) {
	id := c.Get(middleware.UserIDContextKey)
	if id != nil && reflect.TypeOf(id).Name() == "string" {
		return id.(string), nil
	}

	return "", fmt.Errorf(`user id not found`)
}
