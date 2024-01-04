package main

import (
	"postgres-db/internal/event"
	"postgres-db/internal/route"
	"postgres-db/internal/service/repo"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/postgres"
	"github.com/thienhaole92/vnd/runner"
	"github.com/thienhaole92/vnd/wsub"
)

func main() {
	options := []runner.RunnerOption{
		runner.NewInfraHookOption("postgres_db", runner.PostgresHook),
		runner.BuildPostgresDatabaseMigrationHook(migrateDatabaseHook),
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

func migrateDatabaseHook() (string, string, error) {
	c, err := postgres.NewConfig()
	if err != nil {
		return "", "", err
	}

	source := "file://db/migrations"
	databaseURL := c.Url
	return source, databaseURL, nil
}

func initSubscribers(rn *runner.Runner) error {
	p, err := rn.GetInfra().Postgres()
	if err != nil {
		return err
	}

	r := repo.NewRepo(p.Pool)
	e := event.NewEvent(r)

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
