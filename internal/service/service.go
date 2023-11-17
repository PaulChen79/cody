package service

import (
	"cody/config"
	"cody/domain"
)

type service struct {
	repo   domain.Repository
	redis  domain.GoRedis
	config *config.Config
}

func NewService(repo domain.Repository, redis domain.GoRedis, config *config.Config) domain.Service {
	return &service{
		repo:   repo,
		redis:  redis,
		config: config,
	}
}
