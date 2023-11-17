package database

import (
	"fmt"
	"time"

	"cody/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func DatabaseConnection(config *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		config.Host,
		config.User,
		config.Password,
		config.Database,
		config.Port,
		config.Flag,
	)

	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("dialect not implement")
	}

	// setup conn
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Duration(120) * time.Second)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
