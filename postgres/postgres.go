package postgres

import (
	"context"

	zap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/thienhaole92/vnd/logger"
)

type Postgres struct {
	*pgxpool.Pool
}

func NewPostgres(config *Config) (*Postgres, error) {
	pgConfig, err := pgxpool.ParseConfig(config.Url)
	if err != nil {
		return nil, err
	}

	tracer := &tracelog.TraceLog{
		Logger:   zap.NewLogger(logger.GetLogger("db").Desugar()),
		LogLevel: tracelog.LogLevel(config.LogLevel),
	}

	pgConfig.MaxConns = config.MaxConnection
	pgConfig.MinConns = config.MinConnection
	pgConfig.MaxConnIdleTime = config.MaxConnectionIdleTime
	pgConfig.ConnConfig.Tracer = tracer

	pool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Pool: pool,
	}, nil
}
