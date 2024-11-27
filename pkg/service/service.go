package service

import "EffMob/pkg/repositroy"

type Song interface {
}

type Service struct {
	Song
}

func NewService(repo *repositroy.Repository) *Service {
	return &Service{}
}
