package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type CodeContent struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	IdeChatRoomID uint   `gorm:"index"`
	Code          []byte `gorm:"type:bytea;not null"`
	FileName      string `gorm:"type:varchar(150);not null"`
}

func AddCodeContentTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202311182242",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&CodeContent{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("code_contents")
		},
	}
}
