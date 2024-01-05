package wsub

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/thienhaole92/vnd/event"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/redis"
	"github.com/thienhaole92/vnd/subscriber"
)

type wsub struct {
	*redisstream.Subscriber
	topic       string
	groupID     string
	done        chan bool // to end infinite loop
	consumeFunc subscriber.ConsumeFunc
}

func NewSubscriber(config *Config, topic string, consumeFunc subscriber.ConsumeFunc) (subscriber.Subscriber, error) {
	log := logger.GetLogger("NewWatermillSubscriber")
	defer log.Sync()

	c, err := redis.NewConfig()
	if err != nil {
		return nil, err
	}

	log.Infow("loaded redis subscribe config")

	r, err := redis.NewRedis(c)
	if err != nil {
		return nil, err
	}
	log.Infow("redis subscribe connected")

	s, err := redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        r,
			Unmarshaller:  redisstream.DefaultMarshallerUnmarshaller{},
			ConsumerGroup: config.ConsumerGroup,
		},
		watermill.NewStdLogger(config.LoggerDebug, config.LoggerTrace),
	)
	if err != nil {
		return nil, err
	}

	return &wsub{
		Subscriber:  s,
		consumeFunc: consumeFunc,
		done:        make(chan bool, 1),
		topic:       topic,
		groupID:     config.ConsumerGroup,
	}, nil
}

func (r *wsub) Close() error {
	r.done <- true
	defer close(r.done)

	return r.Subscriber.Close()
}

func (r *wsub) GroupID() string {
	return r.groupID
}

func (r *wsub) Start() {
	log := logger.GetLogger("Redis Subcriber")
	defer log.Sync()

	log.With("topic", r.topic)

	log.Infow("subscription start...")
	messages, err := r.Subscriber.Subscribe(context.Background(), r.topic)
	if err != nil {
		log.Panicw("subscribe error, panic now", "error", err)
		r.Close()
	}

	for {
		select {
		case <-r.done:
			log.Info("subscription ended")
			return
		default:
			// continue below to fetch message, etc...
		}

		message := <-messages

		if message == nil || message.UUID == "" {
			log.Debugw("empty message id", "topic", r.topic)
			continue
		}

		if err := r.consumeMessage(context.Background(), message); err != nil {
			log.Errorw("subscription error, fail to consume the message", "error", err, "topic", r.topic)
			continue
		}
	}
}

func (r *wsub) Topic() string {
	return r.topic
}

func (r *wsub) consumeMessage(ctx context.Context, msg *message.Message) error {
	log := logger.GetLogger("")
	defer log.Sync()

	defer func(start time.Time) {
		log.Infow("consume message completed", "elapsed", time.Since(start))
	}(time.Now())

	log.Debugw("consume message start")
	if err := r.consumeFunc(ctx, event.EventString(string(msg.Payload))); err != nil {
		msg.Ack()
		return err
	}

	msg.Ack()
	return nil
}
