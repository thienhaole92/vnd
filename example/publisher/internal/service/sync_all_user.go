package service

import (
	"context"

	"github.com/google/uuid"
	vndcontext "github.com/thienhaole92/vnd/context"
	"github.com/thienhaole92/vnd/rest"

	"github.com/thienhaole92/vnd/logger"
)

type SyncAllUserReq struct {
}

func (s *Service) SyncAllUser(e vndcontext.Context, req *SyncAllUserReq) (*rest.Result, error) {
	delegate := func(log *logger.Logger, ctx vndcontext.Context, req *SyncAllUserReq) (*rest.Result, error) {
		exec := NewSyncAllUsers(
			log,
			NewPublishUserInfo(s.publisher),
		)
		return exec.Execute(ctx, req)
	}
	return rest.Call[SyncAllUserReq](e, req, "SyncAllUser", delegate)
}

type syncAllUsers struct {
	log       *logger.Logger
	publisher PublishUser
}

func NewSyncAllUsers(log *logger.Logger, publisher PublishUser) *syncAllUsers {
	return &syncAllUsers{
		log:       log,
		publisher: publisher,
	}
}

func (s *syncAllUsers) Execute(ctx vndcontext.Context, req *SyncAllUserReq) (*rest.Result, error) {
	defer func() {
		s.log.Sync()
	}()

	if err := s.publisher.PublishUserInfo(context.Background(), &UserInfo{
		Id: uuid.NewString(),
	}); err != nil {
		s.log.Errorw("fail to publish user info events", "error", err)
	}

	return &rest.Result{
		Data: nil,
	}, nil
}
