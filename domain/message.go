package domain

import "time"

type Message struct {
	ID         uint      `json:"id"`
	SendFrom   uint      `json:"send_From"`
	Content    []byte    `json:"content"`
	ChatRoomID uint      `json:"chat_room_id"`
	SendAt     time.Time `json:"send_at"`
	Sequence   uint      `json:"sequence"`
}
