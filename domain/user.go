package domain

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserWithChatRooms struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	ChatRooms []ChatRoom `json:"chat_rooms"`
}

type UserWithIdeChatRooms struct {
	ID           uint          `json:"id"`
	Email        string        `json:"email"`
	Username     string        `json:"username"`
	IdeChatRooms []IdeChatRoom `json:"Ide_chat_rooms"`
}

type UserWithAllTypeChatRooms struct {
	ID           uint          `json:"id"`
	Email        string        `json:"email"`
	Username     string        `json:"username"`
	ChatRooms    []ChatRoom    `json:"chat_rooms"`
	IdeChatRooms []IdeChatRoom `json:"Ide_chat_rooms"`
}

type UserInternal struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
