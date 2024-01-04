package mongo

import "github.com/caarlos0/env/v10"

type Config struct {
	Uri         string `env:"MONGO_URI,required"`
	MaxPoolSize uint64 `env:"MONGO_MAX_POOL_SIZE,required"`
	MinPoolSize uint64 `env:"MONGO_MIN_POOL_SIZE,required"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
