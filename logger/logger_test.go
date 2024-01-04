package logger_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"

	"github.com/thienhaole92/vnd/logger"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Logger", func() {
	It("load from environment", func() {
		config, err := logger.NewConfig()
		if err != nil {
			Fail(fmt.Sprintf("fail to load logger config: %+v", err))
		}

		log := logger.GetLoggerWithConfig("from_environment_1", config.ToZapConfig())
		defer log.Sync()

		log.Info("info logging enabled")
		log.Error("error logging enabled")
		// Output:
		// no output as error only
		// 2022-03-25T14:45:41.509+0800    error   from_environment_1        logger/logger_test.go:49        error logging enabled
		// <<< STACK TRACE HERE >>>
	})

	It("default is development debug console", func() {
		log := logger.GetLogger("default")
		defer log.Sync()

		log.Info("info logging enabled")
		log.Debug("debug logging enabled")

		// Output:
		// 2022-03-25T14:43:02.964+0800    info    default     logger/logger_test.go:75        info logging enabled
		// 2022-03-25T14:43:02.964+0800    debug   default     logger/logger_test.go:76        debug logging enabled
	})

	It("development debug console", func() {
		config := logger.Config{
			Mode:     "development",
			Level:    "debug",
			Encoding: "console",
		}

		log := logger.GetLoggerWithConfig("development_debug_console", config.ToZapConfig())
		defer log.Sync()

		log.Info("info logging enabled")
		log.Debug("debug logging enabled")

		// Output:
		// 2022-03-25T14:43:02.964+0800    info    development_debug_console     logger/logger_test.go:92        info logging enabled
		// 2022-03-25T14:43:02.964+0800    debug   development_debug_console     logger/logger_test.go:93        debug logging enabled
	})

	It("production info json", func() {
		config := logger.Config{
			Mode:     "production",
			Level:    "info",
			Encoding: "json",
		}

		log := logger.GetLoggerWithConfig("production_info_json", config.ToZapConfig())
		defer log.Sync()

		log.Info("info logging enabled")
		log.Debug("debug logging disabled")
		// Output:
		// production is fixed to json
		// {"level":"info","datetime":"2022-03-25T14:43:02.965+0800","logger":"production_info_json","msg":"info logging enabled"}
		// debug is not printed
	})

	It("change log level", func() {
		config := logger.Config{
			Mode:     "development",
			Level:    "info",
			Encoding: "console",
		}

		log := logger.GetLoggerWithConfig("change_log_level", config.ToZapConfig())
		defer log.Sync()

		log.Info("info logging enabled")
		log.Debug("debug logging enabled")
		// Output:
		// 2022-03-25T14:43:02.964+0800    info    change_log_level     logger/logger_test.go:127        info logging enabled
		// debug is not printed

		log.LogLevel(zapcore.DebugLevel)
		log.Debug("debug logging enabled now")
		// 2022-03-25T14:43:02.964+0800    info    change_log_level     logger/logger_test.go:134        debug logging enabled now
	})

	It("canonical", func() {
		// boilerplate to log canonical, if you extract it to logger, the logged function will be incorrect
		log := logger.GetLogger("canonical")
		defer log.Sync()
		defer func(start time.Time) {
			log.Infow("completed", "canonical", true, "elapsed", time.Since(start))
		}(time.Now())

		log.Infow("start")
		log.With("key", "value")
		time.Sleep(8888 * time.Nanosecond)
		// 2022-05-16T14:41:41.502+0800   info    canonical       logger/logger_test.go:139       start
		// 2022-05-16T14:41:41.502+0800   info    canonical       logger/logger_test.go:136       completed       {"key": "value", "canonical": true, "elapsed": 0.000077416}
	})

})
