package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type UserIdeChatRoom struct {
	UserID        uint `gorm:"primaryKey;autoIncrement:false"`
	IdeChatRoomID uint `gorm:"primaryKey;autoIncrement:false"`
}

func AddUserIdeChatRoomTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182238",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&UserIdeChatRoom{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("user_ide_chat_rooms")
		},
	}
}
