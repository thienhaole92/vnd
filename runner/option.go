package runner

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echolog "github.com/labstack/gommon/log"
	"github.com/thienhaole92/vnd/esrv"
	"github.com/thienhaole92/vnd/internal"
	"github.com/thienhaole92/vnd/logger"
	vndmiddleware "github.com/thienhaole92/vnd/middleware"
	"github.com/thienhaole92/vnd/msrv"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/validator"
	"go.uber.org/zap/zapcore"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func buildRestServer(c *esrv.Config, rn *Runner, rsh RestServerHook) (*echo.Echo, error) {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &internal.Context{Context: c}
			return next(cc)
		}
	})

	e.HideBanner = true
	e.Validator = validator.DefaultRestValidator()
	e.HTTPErrorHandler = vndmiddleware.ErrorHandler(e.DefaultHTTPErrorHandler)

	e.Pre(echoprometheus.NewMiddleware(""))
	e.Use(vndmiddleware.RequestLogger(logger.GetLogger("request_info").Desugar(), rest.RestLogFieldExtractor))
	e.Use(middleware.BodyLimit(c.BodyLimit))

	if c.EnableCors {
		e.Use(middleware.CORS())
	}

	root := e.Group("")
	if err := rsh(rn, e, root); err != nil {
		return nil, err
	}

	return e, nil
}

func BuildRestServerOption(rsh RestServerHook) RunnerOption {
	hook := func(rn *Runner) error {
		c, err := esrv.NewConfig()
		if err != nil {
			return err
		}
		rn.log.Infow("loaded service server config", "config", c)

		e, err := buildRestServer(c, rn, rsh)
		if err != nil {
			return err
		}

		es := esrv.NewEsrv(e, c.ToServerConfig())
		es.ListenAndServe()

		rn.AddShutdownHook("shutdown_rest_server", func(*Runner) error {
			es.Shutdown()
			return nil
		})

		return nil
	}
	return NewStartupHookOption("rest_server", hook)
}

func buildMonitorService(rn *Runner, config *msrv.Config, msh MonitorServerHook) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true

	// middlewares
	log := logger.GetLogger("monitor_echo")
	log.LogLevel(zapcore.ErrorLevel) // only show error for /metrics,/status,etc...

	e.Use(vndmiddleware.RequestLogger(log.Desugar()))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  echolog.ERROR,
	}))

	// services
	e.GET(config.StatusPath, func(c echo.Context) error {
		return c.String(http.StatusOK, `{"status":"ok"}`)
	})

	if err := msh(rn, e); err != nil {
		return nil, err
	}

	return e, nil
}

func BuildMonitorServerOption(msh MonitorServerHook) RunnerOption {
	hook := func(rn *Runner) error {
		c, err := msrv.NewConfig()
		if err != nil {
			return err
		}
		rn.log.Infow("loaded monitor server config", "config", c)

		e, err := buildMonitorService(rn, c, msh)
		if err != nil {
			return err
		}

		es := msrv.NewEsrv(e, c.ToServerConfig())
		es.ListenAndServe()

		rn.AddShutdownHook("shutdown_monitor_server", func(*Runner) error {
			es.Shutdown()
			return nil
		})

		return nil
	}
	return NewStartupHookOption("monitor_server", hook)
}

func DefaultMonitorEchoHook(rn *Runner, e *echo.Echo) error {
	c, err := msrv.NewConfig()
	if err != nil {
		return err
	}

	e.GET(c.MetricPath, echoprometheus.NewHandler())
	return nil
}

func BuildPostgresDatabaseMigrationHook(mh MigrationHook) RunnerOption {
	hook := func(rn *Runner) error {
		log := logger.GetLogger("BuildPostgresDatabaseMigrationHook")
		defer log.Sync()

		source, dbUri, err := mh()
		if err != nil {
			return err
		}

		db, err := sql.Open("postgres", dbUri)
		if err != nil {
			return err
		}
		driver, err := pgx.WithInstance(db, &pgx.Config{})
		if err != nil {
			return err
		}

		m, err := migrate.NewWithDatabaseInstance(
			source,
			"postgres",
			driver,
		)
		if err != nil {
			log.Errorw("failed to init migrate source", "error", err)
			return err
		}

		err = m.Up()
		if err == migrate.ErrNoChange {
			log.Info("migrate database done no change")
			m.Close()
			return nil
		}

		if err != nil && err != migrate.ErrNoChange {
			log.Errorw("migrate database failed", "error", err)
			m.Close()
			return err
		}

		m.Close()
		log.Info("migrate database done")
		return nil
	}

	return NewDatabaseMigrationHookOption("database_migration", hook)
}

func BuildSubscribeHook(sh SubscriberHook) RunnerOption {
	hook := func(rn *Runner) error {
		log := logger.GetLogger("BuildSubscribeHook")
		defer log.Sync()

		err := sh(rn)
		if err != nil {
			return err
		}
		log.Info("init event subscribers done")
		return nil
	}

	return NewDatabaseMigrationHookOption("subscribers_init", hook)
}
