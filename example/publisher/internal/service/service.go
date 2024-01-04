package service

import (
	"github.com/thienhaole92/vnd/publisher"
)

type Service struct {
	publisher publisher.Publisher
}

func NewService(publisher publisher.Publisher) *Service {
	s := &Service{
		publisher: publisher,
	}
	return s
}
