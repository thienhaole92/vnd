package route

import (
	"postgres-db/internal/service"
	"postgres-db/internal/service/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	vndmiddleware "github.com/thienhaole92/vnd/middleware"
	"github.com/thienhaole92/vnd/runner"
)

type V1 struct {
	*echo.Group
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	p, err := rn.GetInfra().Postgres()
	if err != nil {
		return err
	}

	r := repo.NewRepo(p.Pool)

	s := service.NewService(r)
	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.Use(vndmiddleware.RequestID(middleware.DefaultSkipper))

	return nil
}
