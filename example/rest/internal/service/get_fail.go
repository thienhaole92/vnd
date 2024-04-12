package service

import (
	vndcontext "github.com/thienhaole92/vnd/context"
	vnderror "github.com/thienhaole92/vnd/error"

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
	return rest.Call(e, req, "GetFail", delegate)
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
	uid, err := ctx.UserId()
	if err != nil {
		return nil, vnderror.InternalServerError(err)
	}

	return &rest.Result{Data: uid}, nil
}
