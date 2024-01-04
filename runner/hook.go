package runner

import (
	"context"
	"time"

	"github.com/thienhaole92/vnd/firebase"
	"github.com/thienhaole92/vnd/logger"
	"github.com/thienhaole92/vnd/mongo"
	"github.com/thienhaole92/vnd/postgres"
	"github.com/thienhaole92/vnd/redis"
)

func RedisHook(rn *Runner) error {
	log := logger.GetLogger("RedisHook")
	defer log.Sync()

	config, err := redis.NewConfig()
	if err != nil {
		return err
	}

	log.Infow("loaded redis config")

	r, err := redis.NewRedis(config)
	if err != nil {
		return err
	}
	log.Infow("redis connected")

	rn.GetInfra().SetRedis(r)

	shutdown := func(rn *Runner) error {
		log := logger.GetLogger("AddShutdownHook")
		defer log.Sync()

		rd, err := rn.GetInfra().Redis()
		if err != nil {
			return err
		}

		if err = rd.Close(); err != nil {
			return err
		}

		log.Infow("shutdown server", "name", "redis")
		return nil
	}
	rn.AddShutdownHook("redis_shutdown", shutdown)

	return nil
}

func MongoHook(rn *Runner) error {
	log := logger.GetLogger("MongoHook")
	defer log.Sync()

	config, err := mongo.NewConfig()
	if err != nil {
		return err
	}

	log.Infow("loaded mongo db config")

	db, err := mongo.NewMongo(config)
	if err != nil {
		return err
	}
	log.Infow("mongo connected")

	rn.GetInfra().SetMongo(db)

	shutdown := func(rn *Runner) error {
		log := logger.GetLogger("AddShutdownHook")
		defer log.Sync()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		mg, err := rn.GetInfra().Mongo()
		if err != nil {
			return err
		}
		if err = mg.Disconnect(ctx); err != nil {
			return err
		}

		log.Infow("shutdown server", "name", "mongodb")
		return nil
	}
	rn.AddShutdownHook("mongo_shutdown", shutdown)

	return nil
}

func FirebaseHook(rn *Runner) error {
	log := logger.GetLogger("FirebaseHook")
	defer log.Sync()

	config, err := firebase.NewConfig()
	if err != nil {
		return err
	}
	log.Infow("loaded firebase config")

	fb, err := firebase.NewFirebase(config)
	if err != nil {
		return err
	}

	rn.GetInfra().SetFirebase(fb)

	log.Infow("firebase connected")

	return nil
}

func PostgresHook(rn *Runner) error {
	log := logger.GetLogger("PostgreHook")
	defer log.Sync()

	config, err := postgres.NewConfig()
	if err != nil {
		return err
	}

	log.Infow("loaded postgres db config")

	db, err := postgres.NewPostgres(config)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*20))
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		return err
	}
	log.Infow("postgres connected")

	rn.GetInfra().SetPostgres(db)

	shutdown := func(rn *Runner) error {
		log := logger.GetLogger("AddShutdownHook")
		defer log.Sync()

		mg, err := rn.GetInfra().Postgres()
		if err != nil {
			return err
		}
		mg.Close()
		log.Infow("shutdown server", "name", "postgres")
		return nil
	}
	rn.AddShutdownHook("postgres_shutdown", shutdown)

	return nil
}
