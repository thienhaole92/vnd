package rest

import (
	"net/http"
	"reflect"
	"runtime"
	"time"

	vnderror "github.com/thienhaole92/vnd/error"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/context"
	"github.com/thienhaole92/vnd/internal"
	"github.com/thienhaole92/vnd/logger"
)

const RequestObjectContextKey = "service_request_object"

func Wrapper[TREQ any](wrapped func(context.Context, *TREQ) (*Result, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.GetLogger("Wrapper")
		defer log.Sync()

		vndc := c.(*internal.Context)
		requestId := vndc.RequestId()
		handler := runtime.FuncForPC(reflect.ValueOf(wrapped).Pointer()).Name()
		log.Infow("request begin",
			"request_id", requestId,
			"at", time.Now().Format(time.RFC3339),
			"path", c.Request().RequestURI,
			"handler", handler,
		)

		var req TREQ
		if err := c.Bind(&req); err != nil {
			log.Errorw("fail to bind request", "request_uri", c.Request().RequestURI, "err", err)
			return &vnderror.Error{
				CustomCode: -40001,
				HTTPError: &echo.HTTPError{
					Code:     http.StatusBadRequest,
					Message:  err.Error(),
					Internal: err,
				},
			}
		}

		if err := c.Validate(&req); err != nil {
			log.Errorw("fail to validate request", "request_uri", c.Request().RequestURI, "request_object", req, "err", err)
			return &vnderror.Error{
				CustomCode: -40002,
				HTTPError: &echo.HTTPError{
					Code:     http.StatusBadRequest,
					Message:  err.Error(),
					Internal: err,
				},
			}
		}

		c.Set(RequestObjectContextKey, req)

		res, err := wrapped(vndc, &req)
		if err != nil {
			return err
		}

		status := c.Response().Status
		if status != 0 {
			log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", status)
			return c.JSON(status, res)
		}

		log.Infow("request end", "request_id", requestId, "at", time.Now().Format(time.RFC3339), "status", http.StatusOK)
		return c.JSON(http.StatusOK, res)
	}
}
