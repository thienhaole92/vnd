package infra

import (
	"errors"

	"github.com/thienhaole92/vnd/firebase"
	"github.com/thienhaole92/vnd/mongo"
	"github.com/thienhaole92/vnd/postgres"
	"github.com/thienhaole92/vnd/redis"
)

type Infra struct {
	redis    *redis.Redis
	mongo    *mongo.Mongo
	postgres *postgres.Postgres
	firebase *firebase.Firebase
}

func (i *Infra) Redis() (*redis.Redis, error) {
	if i.redis == nil {
		return nil, errors.New("redis client is not set")
	}

	return i.redis, nil
}

func (i *Infra) SetRedis(r *redis.Redis) {
	i.redis = r
}

func (i *Infra) Mongo() (*mongo.Mongo, error) {
	if i.mongo == nil {
		return nil, errors.New("mongo client is not set")
	}

	return i.mongo, nil
}

func (i *Infra) SetMongo(m *mongo.Mongo) {
	i.mongo = m
}

func (i *Infra) Firebase() (*firebase.Firebase, error) {
	if i.firebase == nil {
		return nil, errors.New("firebase client is not set")
	}

	return i.firebase, nil
}

func (i *Infra) SetFirebase(f *firebase.Firebase) {
	i.firebase = f
}

func (i *Infra) Postgres() (*postgres.Postgres, error) {
	if i.postgres == nil {
		return nil, errors.New("postgres client is not set")
	}

	return i.postgres, nil
}

func (i *Infra) SetPostgres(p *postgres.Postgres) {
	i.postgres = p
}
