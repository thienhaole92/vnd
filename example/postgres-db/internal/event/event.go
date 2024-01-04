package event

import (
	"postgres-db/internal/service/repo"
)

type Event struct {
	repo *repo.Repo
}

func NewEvent(repo *repo.Repo) *Event {
	s := &Event{
		repo: repo,
	}
	return s
}
