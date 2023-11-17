package gorm

import (
	"cody/config"
	"cody/domain"

	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	redis  domain.GoRedis
	config *config.Config
}

func NewRepository(db *gorm.DB, redis domain.GoRedis, config *config.Config) domain.Repository {
	return &repository{
		db:     db,
		redis:  redis,
		config: config,
	}
}
