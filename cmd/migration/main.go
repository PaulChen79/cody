package main

import (
	v1 "cody/cmd/migration/v1"
	internalLogget "cody/internal/logger"
	"cody/internal/provider"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	zapLogger := internalLogget.NewLogger()
	zap.ReplaceGlobals(zapLogger)
	defer zapLogger.Sync()

	config := provider.NewConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s %s",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Database,
		config.DB.Port,
		config.DB.Flag,
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
		zap.S().Warnf("DatabaseConnection - err: %v", err)
		return
	}
	if db == nil {
		zap.S().Warn("DatabaseConnection - db is nil")
		return
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		v1.AddUserTable(),
		v1.AddMessageTable(),
		v1.AddCodeContentTable(),
		v1.AddIdeChatRoomTable(),
		v1.AddChatRoomTable(),
		v1.AddUserChatRoomTable(),
		v1.AddUserIdeChatRoomTable(),
	})

	if err = m.Migrate(); err != nil {
		zap.S().Warnf("Migration failed: %v", err)
		return
	}
	zap.S().Info("Migration did run successfully")
}
