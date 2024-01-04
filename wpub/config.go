package wpub

import "github.com/caarlos0/env/v10"

type Config struct {
	LoggerDebug bool `env:"REDIS_PUB_LOGGER_DEBUG,notEmpty" envdefault:"false"`
	LoggerTrace bool `env:"REDIS_PUB_LOGGER_TRACE,notEmpty" envdefault:"false"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
