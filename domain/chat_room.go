package domain

type ChatRoom struct {
	ID            uint   `json:"id"`
	RoomName      string `json:"room_name"`
	IsGroup       bool   `json:"isGroup"`
	IdeChatRoomID uint   `json:"ide_chat_room_id"`
}

type ChatRoomFullInfo struct {
	ID            uint       `json:"id"`
	RoomName      string     `json:"room_name"`
	IsGroup       bool       `json:"isGroup"`
	Messages      []*Message `json:"messages"`
	IdeChatRoomID uint       `json:"ide_chat_room_id"`
}
