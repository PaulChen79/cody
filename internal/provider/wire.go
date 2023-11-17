//go:build wireinject
// +build wireinject

package provider

import (
	"log"
	"sync"

	"cody/config"
	"cody/domain"
	"cody/internal/database"
	"cody/internal/redis"
	repo "cody/internal/repo/gorm"
	svc "cody/internal/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once

func NewDB() (*gorm.DB, error) {
	var err error
	if db == nil {
		dbOnce.Do(func() {
			log.Println("connect db")
			db, err = database.DatabaseConnection(NewConfig().DB)
			if err != nil {
				return
			}
			log.Println("connect db success")
		})
	}
	return db, err
}

var cg *config.Config
var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		log.Println("read config")
		cg = config.NewConfig()
		log.Println("read config success")
	})
	return cg
}

func NewRedis() (domain.GoRedis, error) {
	panic(wire.Build(redis.NewGoRedis, NewConfig))
}

func NewRepo() (domain.Repository, error) {
	panic(wire.Build(repo.NewRepository, NewDB, NewRedis, NewConfig))
}

func NewService() (domain.Service, error) {
	panic(wire.Build(svc.NewService, repo.NewRepository, NewDB, NewRedis, NewConfig))
}
