package msrv

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
	Host       string `env:"MONITOR_SERVER_HOST,required"`
	Port       int    `env:"MONITOR_SERVER_PORT,required"`
	MetricPath string `env:"MONITOR_SERVER_METRIC_PATH" envDefault:"/metric"`
	StatusPath string `env:"MONITOR_SERVER_STATUS_PATH" envDefault:"/status"`
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
		Name:         "msrv",
		Host:         c.Host,
		Port:         c.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		GracePeriod:  5 * time.Second,
	}
}
