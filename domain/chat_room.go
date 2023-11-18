package domain

type ChatRoom struct {
	ID            uint   `json:"id"`
	RoomName      string `json:"room_name"`
	IsGroup       bool   `json:"is_group"`
	IdeChatRoomID uint   `json:"ide_chat_room_id"`
}

type ChatRoomFullInfo struct {
	ID            uint   `json:"id"`
	RoomName      string `json:"room_name"`
	IsGroup       bool   `json:"is_group"`
	Messages      []*Message
	Users         []*UserInternal
	IdeChatRoomID uint
}
