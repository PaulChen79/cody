package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type IdeChatRoom struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	RoomName     string        `gorm:"type:varchar(150);not null"`
	RoomUUID     string        `gorm:"type:varchar(150);not null"`
	ChatRoom     ChatRoom      `gorm:"foreignKey:IdeChatRoomID"`
	Users        []User        `gorm:"many2many:user_ide_chat_rooms;"`
	CodeContents []CodeContent `gorm:"foreignKey:IdeChatRoomID"`
}

func AddIdeChatRoomTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182206",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&IdeChatRoom{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("chat_rooms")
		},
	}
}
