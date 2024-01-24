package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	vndcontext "github.com/thienhaole92/vnd/context"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/rest"
)

type GetFailReq struct {
}

func (s *Service) GetFail(e vndcontext.Context, req *GetFailReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx vndcontext.Context, req *GetFailReq) (*rest.Result, error) {
		exec := NewGetFail(log)
		return exec.Execute(ctx, req)
	}
	return rest.Call[GetFailReq](e, req, "GetFail", delegate)
}

type getFail struct {
	log *logger.Logger
}

func NewGetFail(log *logger.Logger) *getFail {
	return &getFail{
		log: log,
	}
}

func (s *getFail) Execute(ctx vndcontext.Context, req *GetFailReq) (res *rest.Result, err error) {
	defer func() {
		if err != nil {
			s.log.With("err", err)
		}

		s.log.Infow("completed")
		s.log.Sync()
	}()

	s.log.With("func1", "ok")

	uid, err := ctx.UserId()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	return &rest.Result{
		Data: uid,
	}, nil
}
