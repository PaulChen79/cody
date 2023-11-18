package v1

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Message struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	SendFrom   uint      `gorm:"type:int;not null"`
	Content    []byte    `gorm:"type:bytea;not null"`
	ChatRoomID uint      `gorm:"index"`
	SendAt     time.Time `gorm:"not null"`
	Sequence   uint      `gorm:"type:int;autoIncrement"`
}

func AddMessageTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182204",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&Message{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("messages")
		},
	}
}
