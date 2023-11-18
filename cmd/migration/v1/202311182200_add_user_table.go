package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	Email        string        `gorm:"type:varchar(100);unique;not null"`
	Password     string        `gorm:"type:varchar(100);unique;not null"`
	Username     string        `gorm:"type:varchar(255);unique;not null"`
	ChatRooms    []ChatRoom    `gorm:"many2many:user_chat_rooms;"`
	IdeChatRooms []IdeChatRoom `gorm:"many2many:user_ide_chat_rooms;"`
}

func AddUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182200",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
