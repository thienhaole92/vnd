package service

import (
	"postgres-db/internal/service/repo"
)

type Service struct {
	repo *repo.Repo
}

func NewService(repo *repo.Repo) *Service {
	s := &Service{
		repo: repo,
	}
	return s
}
