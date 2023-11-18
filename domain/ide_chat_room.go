package domain

type IdeChatRoom struct {
	ID           uint           `json:"id"`
	RoomName     string         `json:"room_Name"`
	RoomUUID     string         `json:"room_uuid"`
	ChatRoom     *ChatRoom      `json:"chat_room"`
	CodeContents []*CodeContent `json:"code_contents"`
}
