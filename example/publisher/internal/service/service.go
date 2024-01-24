package service

import (
	topic "user-service-v5/internal/event"

	"github.com/thienhaole92/vnd/runner"
	"github.com/thienhaole92/vnd/wpub"
)

type Service struct {
	publisher *Publisher
}

func NewService(rn *runner.Runner) (*Service, error) {
	pub, err := wpub.RegisterPublisher(rn, topic.TOPIC_PRIVATE_USER_INFO)
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		UserInfo: pub,
	}

	s := &Service{
		publisher: publisher,
	}
	return s, nil
}
