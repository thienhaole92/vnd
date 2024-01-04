package esrv

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v10"
)

type serverConfig struct {
	Name         string
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	GracePeriod  time.Duration
}

func (sc *serverConfig) Addr() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

type Config struct {
	Host         string        `env:"REST_SERVER_HOST,required"`
	Port         int           `env:"REST_SERVER_PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"REST_SERVER_READ_TIMEOUT,required"`
	WriteTimeout time.Duration `env:"REST_SERVER_WRITE_TIMEOUT,required"`
	EnableCors   bool          `env:"REST_ENABLE_CORS" envDefault:"false"`
	BodyLimit    string        `env:"REST_BODY_LIMIT" envDefault:"8K"`
}

func NewConfig() (*Config, error) {
	c := new(Config)
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) ToServerConfig() *serverConfig {
	return &serverConfig{
		Name:         "esrv",
		Host:         c.Host,
		Port:         c.Port,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		GracePeriod:  5 * time.Second,
	}
}
