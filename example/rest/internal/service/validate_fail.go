package service

import (
	vndcontext "github.com/thienhaole92/vnd/context"
	_ "github.com/thienhaole92/vnd/error"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/rest"
)

type ValidateFailReq struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0,lte=130"`
}

func (s *Service) ValidateFail(e vndcontext.Context, req *ValidateFailReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx vndcontext.Context, req *ValidateFailReq) (*rest.Result, error) {
		exec := NewValidateFail(log)
		return exec.Execute(ctx, req)
	}
	return rest.Call[ValidateFailReq](e, req, "ValidateFail", delegate)
}

type validateFail struct {
	log *logger.Logger
}

func NewValidateFail(log *logger.Logger) *validateFail {
	return &validateFail{
		log: log,
	}
}

func (s *validateFail) Execute(ctx vndcontext.Context, req *ValidateFailReq) (*rest.Result, error) {
	return &rest.Result{
		Data: nil,
	}, nil
}
