package wsub

import (
	"fmt"
	"strings"

	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/subscriber"
)

// Subscriber to manage multiple subscribers
type MultiSubscriber struct {
	config      *Config
	subscribers []subscriber.Subscriber
}

func NewMultiSubscriber(config *Config) *MultiSubscriber {
	return &MultiSubscriber{
		config:      config,
		subscribers: make([]subscriber.Subscriber, 0),
	}
}

func (s *MultiSubscriber) Subscribe(topic string, consumeFunc subscriber.ConsumeFunc) error {
	log := logger.GetLogger("Subscribe")
	defer log.Sync()

	newSubscriber, err := NewSubscriber(s.config, topic, consumeFunc)
	if err != nil {
		return err
	}
	s.subscribers = append(s.subscribers, newSubscriber)

	go newSubscriber.Start()

	return nil
}

func (s *MultiSubscriber) Close() error {
	errStrings := []string{}
	for _, sub := range s.subscribers {
		if err := sub.Close(); err != nil {
			errStrings = append(errStrings, fmt.Sprintf("[%s/%s:%s]", sub.GroupID(), sub.Topic(), err.Error()))
		}
	}
	if len(errStrings) == 0 {
		return nil
	}

	return fmt.Errorf("errors when close subscribers: %s", strings.Join(errStrings, ","))
}
