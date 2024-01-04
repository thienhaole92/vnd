package main

import (
	"event-redis/internal/event"
	"event-redis/internal/route"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/runner"
	"github.com/thienhaole92/vnd/wsub"
)

func main() {
	options := []runner.RunnerOption{
		runner.BuildMonitorServerOption(runner.DefaultMonitorEchoHook),
		runner.BuildRestServerOption(restServiceHook),
		runner.BuildSubscribeHook(initSubscribers),
	}
	r := runner.NewRunner(options...)
	r.Run()
}

func restServiceHook(rn *runner.Runner, e *echo.Echo, eg *echo.Group) error {
	log := logger.GetLogger("restServiceHook")
	defer log.Sync()

	v1 := &route.V1{Group: eg.Group("/v1")}
	if err := v1.Configure(rn); err != nil {
		log.Errorw("failed to register api routes", "error", err)
		return err
	}

	return nil
}

func initSubscribers(rn *runner.Runner) error {
	e := event.NewEvent()

	c, err := wsub.NewConfig()
	if err != nil {
		return err
	}

	mts := wsub.NewMultiSubscriber(c)
	mts.Subscribe(
		event.TOPIC_PRIVATE_USER_INFO,
		e.ConsumeUserInfoEvent,
	)

	rn.AddShutdownHook("close_subscriber", func(*runner.Runner) error {
		mts.Close()
		return nil
	})

	return nil
}
