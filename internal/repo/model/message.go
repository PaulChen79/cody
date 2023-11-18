package gorm

import (
	"cody/domain"
	"time"
)

type Message struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	SendFrom   uint      `gorm:"type:int;not null"`
	Content    []byte    `gorm:"type:bytea;not null"`
	ChatRoomID uint      `gorm:"index"`
	SendAt     time.Time `gorm:"not null"`
	Sequence   uint      `gorm:"type:int;autoIncrement"`
}

func MessageModelToDomain(message *Message) *domain.Message {
	return &domain.Message{
		ID:         message.ID,
		SendFrom:   message.SendFrom,
		Content:    message.Content,
		ChatRoomID: message.ChatRoomID,
		SendAt:     message.SendAt,
		Sequence:   message.Sequence,
	}
}

func MessageDomainToModel(message *domain.Message) *Message {
	return &Message{
		ID:         message.ID,
		SendFrom:   message.SendFrom,
		Content:    message.Content,
		ChatRoomID: message.ChatRoomID,
		SendAt:     message.SendAt,
		Sequence:   message.Sequence,
	}
}
