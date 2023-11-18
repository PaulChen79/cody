package domain

type CodeContent struct {
	ID            uint   `json:"id"`
	IdeChatRoomID uint   `json:"ide_chat_room_id"`
	Code          []byte `json:"code"`
	FileName      string `json:"file_name"`
}
