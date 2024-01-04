package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger

	// Use this to change log level dynamically
	level zap.AtomicLevel
}

func (log *Logger) LogLevel(newLevel zapcore.Level) {
	log.level.SetLevel(newLevel)
}

func (log *Logger) With(args ...interface{}) {
	log.SugaredLogger = log.SugaredLogger.With(args...)
}

func GetDefaultZapConfig() zap.Config {
	if lc, err := NewConfig(); err != nil {
		panic(fmt.Errorf("fail to load logger config: %+v", err))
	} else {
		return lc.ToZapConfig()
	}
}

func GetLogger(name string, options ...zap.Option) *Logger {
	return GetLoggerWithConfig(name, GetDefaultZapConfig(), options...)
}

func GetLoggerWithConfig(name string, logConfig zap.Config, options ...zap.Option) *Logger {
	log, err := logConfig.Build(options...)
	if err != nil {
		panic("cannot get logger is dead!")
	}

	return &Logger{log.Sugar().Named(name), logConfig.Level}
}
