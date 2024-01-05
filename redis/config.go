package redis

import (
	"time"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Host         string        `env:"REDIS_HOST,required"`
	Port         string        `env:"REDIS_PORT,required"`
	Password     string        `env:"REDIS_PASSWORD" envdefault:""`
	DB           int           `env:"REDIS_DB,required"`
	TTL          time.Duration `env:"REDIS_TTL,required"`
	DialTimeout  time.Duration `env:"REDIS_DIAL_TIMEOUT" envdefault:"10s"`
	UseTls       bool          `env:"REDIS_USE_TLS" envdefault:"false"`
	MaxIdleConns int           `env:"REDIS_MAX_IDLE_CONNS,required"`
	MinIdleConns int           `env:"REDIS_MIN_IDLE_CONNS,required"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
