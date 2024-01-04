package event

import (
	"context"
	"postgres-db/internal/service/repo"

	"github.com/thienhaole92/vnd/event"
	"github.com/thienhaole92/vnd/logger"
)

func (e *Event) ConsumeUserInfoEvent(ctx context.Context, message event.EventString) error {
	delegate := func(log *logger.Logger, ctx context.Context, schema *event.EventSchema[*repo.UserInfoRecord]) error {
		exec := NewConsumeUserInfo(log, e.repo)
		return exec.Execute(ctx, schema.Data)
	}
	return event.Consume[*repo.UserInfoRecord](ctx, message, "ConsumeUserInfoEvent", delegate)
}

type ConsumeUserInfoRepo interface {
	BatchUpsertUserInfo(ctx context.Context, ul repo.UserList) error
}

type consumeUserInfo struct {
	log  *logger.Logger
	repo ConsumeUserInfoRepo
}

func NewConsumeUserInfo(log *logger.Logger, repo ConsumeUserInfoRepo) *consumeUserInfo {
	return &consumeUserInfo{
		log:  log,
		repo: repo,
	}
}

func (s *consumeUserInfo) Execute(ctx context.Context, user *repo.UserInfoRecord) error {
	return s.repo.BatchUpsertUserInfo(ctx, repo.UserList{
		user,
	})
}
