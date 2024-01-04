package event

import (
	"context"

	"github.com/thienhaole92/vnd/event"
	"github.com/thienhaole92/vnd/logger"
)

type UserInfoRecord struct {
	Id string `json:"id"`
}

func (e *Event) ConsumeUserInfoEvent(ctx context.Context, message event.EventString) error {
	delegate := func(log *logger.Logger, ctx context.Context, schema *event.EventSchema[*UserInfoRecord]) error {
		exec := NewConsumeUserInfo(log)
		return exec.Execute(ctx, schema.Data)
	}
	return event.Consume[*UserInfoRecord](ctx, message, "ConsumeUserInfoEvent", delegate)
}

type consumeUserInfo struct {
	log *logger.Logger
}

func NewConsumeUserInfo(log *logger.Logger) *consumeUserInfo {
	return &consumeUserInfo{
		log: log,
	}
}

func (s *consumeUserInfo) Execute(ctx context.Context, user *UserInfoRecord) error {
	s.log.Infow("message", "data", user)
	return nil
}
