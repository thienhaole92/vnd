package logger

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Mode     string `env:"LOG_MODE" envDefault:"development"`
	Level    string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	Encoding string `env:"LOG_ENCODING" envDefault:"console"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) ToZapConfig() zap.Config {
	var logConfig zap.Config

	if c.Mode == "production" {
		logConfig = zap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logConfig.EncoderConfig.TimeKey = "datetime"
		logConfig.DisableCaller = true
		logConfig.Encoding = "json"
	} else {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig = zap.NewProductionEncoderConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logConfig.EncoderConfig.TimeKey = "datetime"
		logConfig.Encoding = c.Encoding
	}

	atom := zap.NewAtomicLevel()
	if c.Level != "" {
		newLevel := zap.NewAtomicLevel()
		err := newLevel.UnmarshalText([]byte(c.Level))
		if err != nil {
			fmt.Printf("ignored to invalid log level '%s'\n", c.Level)
		} else {
			atom.SetLevel(newLevel.Level())
		}
	}
	logConfig.Level = atom

	return logConfig
}
