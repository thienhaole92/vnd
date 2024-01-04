package repo

import (
	"github.com/thienhaole92/vnd/postgres"
)

type Repo struct {
	db postgres.PostgreSQL
}

func NewRepo(db postgres.PostgreSQL) *Repo {
	return &Repo{
		db: db,
	}
}
