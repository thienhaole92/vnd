package wpub

import (
	"fmt"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/publisher"
	"github.com/thienhaole92/vnd/redis"
	"github.com/thienhaole92/vnd/runner"
)

type wpub struct {
	*redisstream.Publisher
	topic string
}

func NewPublisher(config *Config, topic string) (publisher.Publisher, error) {
	log := logger.GetLogger("NewPublisher")
	defer log.Sync()

	c, err := redis.NewConfig()
	if err != nil {
		return nil, err
	}

	log.Infow("loaded redis publisher config")

	r, err := redis.NewRedis(c)
	if err != nil {
		return nil, err
	}
	log.Infow("redis publisher connected")

	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     r,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
		},
		watermill.NewStdLogger(config.LoggerDebug, config.LoggerTrace),
	)

	if err != nil {
		return nil, err
	}

	return &wpub{
		Publisher: publisher,
		topic:     topic,
	}, nil
}

func (w *wpub) PublishMessage(messages ...string) error {
	msgs := make([]*message.Message, 0)

	for _, v := range messages {
		msgs = append(msgs, message.NewMessage(watermill.NewUUID(), []byte(v)))
	}

	return w.Publish(
		w.topic,
		msgs...,
	)
}

func (w *wpub) Topic() string {
	return w.topic
}

func RegisterPublisher(rn *runner.Runner, topic string) (publisher.Publisher, error) {
	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	pub, err := NewPublisher(c, topic)
	if err != nil {
		return nil, err
	}

	hn := fmt.Sprintf("close_%s_publisher", topic)
	hn = strings.ReplaceAll(hn, " ", "_")
	hn = strings.ReplaceAll(hn, ".", "_")

	rn.AddShutdownHook(hn, func(*runner.Runner) error {
		return pub.Close()
	})

	return pub, nil
}
