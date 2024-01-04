package service

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	topic "user-service-v5/internal/event"

	"github.com/google/uuid"
	"github.com/thienhaole92/vnd/event"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/publisher"
)

var source = "user-service-v5"

type UserInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func init() {
	if debugInfo, ok := debug.ReadBuildInfo(); ok {
		if debugInfo.Main.Version != "" {
			source = fmt.Sprintf("%s:%s", source, debugInfo.Main.Version)
		}
	}
}

type PublishUser interface {
	Close() error
	PublishUserInfo(ctx context.Context, ui *UserInfo) error
}

type batchPublishUserInfo struct {
	publisher.Publisher
	log *logger.Logger
}

func NewPublishUserInfo(ep publisher.Publisher) PublishUser {
	return &batchPublishUserInfo{
		Publisher: ep,
		log:       logger.GetLogger("BatchPublishUserInfo"),
	}
}

func (s *batchPublishUserInfo) PublishUserInfo(ctx context.Context, ui *UserInfo) (err error) {
	count := 0
	defer func() {
		s.log.Infow("published", "error", err)
		s.log.Sync()
	}()

	messageList := make([]string, 0)
	e := event.EventSchema[*UserInfo]{
		Specversion: "1.0",
		Type:        "User.UserInfo",
		Source:      source,
		Id:          uuid.NewString(),
		Time:        time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
		Subject:     topic.TOPIC_PRIVATE_USER_INFO,
		Data:        ui,
	}

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	messageList = append(messageList, string(b))

	if err := s.Publisher.PublishMessage(messageList...); err != nil {
		return err
	}
	s.log.Debugw("batchPublish", "count", count)
	return nil
}
