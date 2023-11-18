package gorm

import "cody/domain"

type IdeChatRoom struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	RoomName     string        `gorm:"type:varchar(150);not null"`
	RoomUUID     string        `gorm:"type:varchar(150);not null"`
	ChatRoom     ChatRoom      `gorm:"foreignKey:IdeChatRoomID;references:ID"`
	Users        []User        `gorm:"many2many:user_ide_chat_rooms;"`
	CodeContents []CodeContent `gorm:"foreignKey:IdeChatRoomID"`
}

func IdeChatRoomModelToDomain(ideChatRoom IdeChatRoom) *domain.IdeChatRoom {
	domainCodeContents := []*domain.CodeContent{}

	for i := range ideChatRoom.CodeContents {
		dcc := CodeContentModelToDomain(ideChatRoom.CodeContents[i])
		domainCodeContents = append(domainCodeContents, dcc)
	}

	return &domain.IdeChatRoom{
		ID:           ideChatRoom.ID,
		RoomName:     ideChatRoom.RoomName,
		RoomUUID:     ideChatRoom.RoomUUID,
		CodeContents: domainCodeContents,
		ChatRoom:     ChatRoomModelToDomain(ideChatRoom.ChatRoom),
	}
}

func IdeChatRoomDomainToModel(ideChatRoom *domain.IdeChatRoom) *IdeChatRoom {
	return &IdeChatRoom{
		ID:       ideChatRoom.ID,
		RoomName: ideChatRoom.RoomName,
		RoomUUID: ideChatRoom.RoomUUID,
	}
}
