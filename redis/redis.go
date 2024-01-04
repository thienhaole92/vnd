package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*redis.Client
}

func NewRedis(config *Config) (*Redis, error) {
	addr := fmt.Sprintf(`%s:%s`, config.Host, config.Port)

	opts := redis.Options{
		Addr:         addr,
		Password:     config.Password,
		DB:           config.DB,
		DialTimeout:  config.DialTimeout,
		MaxIdleConns: config.MaxIdleConns,
		MinIdleConns: config.MinIdleConns,
	}

	if config.UseTls {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	r := &Redis{
		Client: redis.NewClient(&opts),
	}

	if _, err := r.Client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return r, nil
}
