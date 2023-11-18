package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type ChatRoom struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	RoomName      string    `gorm:"type:varchar(150);not null"`
	IsGroup       bool      `gorm:"type:boolean;default:false;not null"`
	Messages      []Message `gorm:"foreignKey:ChatRoomID"`
	Users         []User    `gorm:"many2many:user_chat_rooms;"`
	IdeChatRoomID uint      `gorm:"index"`
}

func AddChatRoomTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182208",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&ChatRoom{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("chat_rooms")
		},
	}
}
