package runner

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/infra"
	"github.com/thienhaole92/vnd/logger"
)

type RunnerOption = func(*Runner)

// Hook to add stuff to RunnerRuntime
type Hook = func(*Runner) error

type RestServerHook = func(*Runner, *echo.Echo, *echo.Group) error

type MonitorServerHook = func(*Runner, *echo.Echo) error

type MigrationHook = func() (string, string, error)

type SubscriberHook = func(*Runner) error

type Runner struct {
	log              *logger.Logger
	infra            *infra.Infra
	infraHooks       []Hook
	dbMigrationHooks []Hook
	startupHooks     []Hook
	runHooks         []Hook
	subscriberHook   []Hook
	shutdownHooks    []Hook
}

func NewRunner(options ...RunnerOption) *Runner {
	log := logger.GetLogger("Runner")
	defer log.Sync()

	r := &Runner{
		log:              log,
		infra:            &infra.Infra{},
		infraHooks:       []Hook{},
		dbMigrationHooks: []Hook{},
		startupHooks:     []Hook{},
		runHooks:         []Hook{},
		subscriberHook:   []Hook{},
		shutdownHooks:    []Hook{},
	}

	for _, opt := range options {
		opt(r)
	}

	return r
}

func (r *Runner) AddInfraHook(name string, h Hook) {
	r.infraHooks = append(r.infraHooks, h)
	r.log.Debugw("registered infra hook", "hook_name", name)
}

func (r *Runner) AddDatabaseMigrationHook(name string, h Hook) {
	r.dbMigrationHooks = append(r.dbMigrationHooks, h)
	r.log.Debugw("registered database migration hook", "hook_name", name)
}

func (r *Runner) AddStartupHook(name string, h Hook) {
	r.startupHooks = append(r.startupHooks, h)
	r.log.Debugw("registered startup hook", "hook_name", name)
}

func (r *Runner) AddRunHook(name string, h Hook) {
	r.runHooks = append(r.runHooks, h)
	r.log.Debugw("registered run hook", "hook_name", name)
}

func (r *Runner) AddSubscriberHook(name string, h Hook) {
	r.subscriberHook = append(r.subscriberHook, h)
	r.log.Debugw("registered subscriber hook", "hook_name", name)
}

func (r *Runner) AddShutdownHook(name string, h Hook) {
	r.shutdownHooks = append(r.shutdownHooks, h)
	r.log.Debugw("registered shutdown hook", "hook_name", name)
}

func (r *Runner) GetInfra() *infra.Infra {
	return r.infra
}

func (r *Runner) Run() {
	log := logger.GetLogger("Run")
	defer log.Sync()

	log.Info("starting up runners")

	for _, hook := range r.infraHooks {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	for _, hook := range r.dbMigrationHooks {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	for _, hook := range r.startupHooks {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	// listen to signal to graceful shutdown before we start the server
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	for _, hook := range r.runHooks {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	for _, hook := range r.subscriberHook {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	// signal to thread dump
	go func() {
		threadDump := make(chan os.Signal, 1)
		signal.Notify(threadDump, syscall.SIGUSR2)
		buf := make([]byte, 1<<20)
		for {
			<-threadDump
			stacklen := runtime.Stack(buf, true)
			log.Infow("stack trace", "dump", string(buf[:stacklen]))
		}
	}()

	// gracefully shutdown
	sig := <-shutdown
	log.Info(fmt.Sprintf("received signal: %s, shutting down now...", sig.String()))

	for _, hook := range r.shutdownHooks {
		if err := hook(r); err != nil {
			panic(err)
		}
	}

	log.Infow("stopped everything")
}

func NewInfraHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddInfraHook(name, h)
	}
}

func NewDatabaseMigrationHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddDatabaseMigrationHook(name, h)
	}
}

func NewStartupHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddStartupHook(name, h)
	}
}

func NewRunHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddRunHook(name, h)
	}
}

func NewSubscriberHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddSubscriberHook(name, h)
	}
}

func NewShutdownHookOption(name string, h Hook) RunnerOption {
	return func(r *Runner) {
		r.AddShutdownHook(name, h)
	}
}
