package gorm

import "cody/domain"

type ChatRoom struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	RoomName      string    `gorm:"type:varchar(150);not null"`
	IsGroup       bool      `gorm:"type:boolean;default:false;not null"`
	Messages      []Message `gorm:"foreignKey:ChatRoomID"`
	Users         []User    `gorm:"many2many:user_chat_rooms;"`
	IdeChatRoomID uint      `gorm:"index"`
}

func ChatRoomModelToDomain(chatRoom *ChatRoom) *domain.ChatRoom {
	return &domain.ChatRoom{
		ID:            chatRoom.ID,
		RoomName:      chatRoom.RoomName,
		IsGroup:       chatRoom.IsGroup,
		IdeChatRoomID: chatRoom.IdeChatRoomID,
	}
}

func ChatRoomModelToFullInfoDomain(chatRoom *ChatRoom) *domain.ChatRoomFullInfo {
	domainMessages := []*domain.Message{}

	for i := range chatRoom.Messages {
		dm := MessageModelToDomain(&chatRoom.Messages[i])
		domainMessages = append(domainMessages, dm)
	}

	return &domain.ChatRoomFullInfo{
		ID:            chatRoom.ID,
		RoomName:      chatRoom.RoomName,
		IsGroup:       chatRoom.IsGroup,
		IdeChatRoomID: chatRoom.IdeChatRoomID,
		Messages:      domainMessages,
	}
}

func ChatRoomDomainToModel(chatRoom *domain.ChatRoom) *ChatRoom {
	return &ChatRoom{
		ID:            chatRoom.ID,
		RoomName:      chatRoom.RoomName,
		IsGroup:       chatRoom.IsGroup,
		IdeChatRoomID: chatRoom.IdeChatRoomID,
	}
}

func ChatRoomFullInfoDomainToModel(chatRoom *domain.ChatRoomFullInfo) *ChatRoom {
	messages := []Message{}

	for i := range chatRoom.Messages {
		messages = append(messages, *MessageDomainToModel(chatRoom.Messages[i]))
	}

	users := []User{}

	for i := range chatRoom.Users {
		users = append(users, *UserInternalDomainToModel(chatRoom.Users[i]))
	}

	return &ChatRoom{
		ID:            chatRoom.ID,
		RoomName:      chatRoom.RoomName,
		IsGroup:       chatRoom.IsGroup,
		IdeChatRoomID: chatRoom.IdeChatRoomID,
		Messages:      messages,
		Users:         users,
	}
}
