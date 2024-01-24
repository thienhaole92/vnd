package route

import (
	"user-service-v5/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	vndmiddleware "github.com/thienhaole92/vnd/middleware"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/runner"
)

type V1 struct {
	*echo.Group
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	s, err := service.NewService(rn)
	if err != nil {
		return err
	}

	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.Use(vndmiddleware.RequestID(middleware.DefaultSkipper))

	v1.GET("/sync-all-user", rest.Wrapper(s.SyncAllUser))

	return nil
}
