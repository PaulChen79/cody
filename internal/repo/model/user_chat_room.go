package gorm

import "time"

type UserChatRoom struct {
	UserID     uint `gorm:"primaryKey;autoIncrement:false"`
	ChatRoomID uint `gorm:"primaryKey;autoIncrement:false"`
	LastReadAt *time.Time
}
