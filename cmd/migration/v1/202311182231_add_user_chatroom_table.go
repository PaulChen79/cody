package v1

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type UserChatRoom struct {
	UserID     uint      `gorm:"primaryKey;autoIncrement:false"`
	ChatRoomID uint      `gorm:"primaryKey;autoIncrement:false"`
	LastReadAt time.Time `gorm:"not null"`
}

func AddUserChatRoomTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182231",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&UserChatRoom{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("user_chat_rooms")
		},
	}
}
