package gorm

import "cody/domain"

type User struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	Email        string        `gorm:"type:varchar(100);unique;not null"`
	Password     string        `gorm:"type:varchar(100);unique;not null"`
	Username     string        `gorm:"type:varchar(255);unique;not null"`
	ChatRooms    []ChatRoom    `gorm:"many2many:user_chat_rooms;"`
	IdeChatRooms []IdeChatRoom `gorm:"many2many:user_ide_chat_rooms;"`
}

func UserModelToDomain(user User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func UserInternalDomainToModel(user *domain.UserInternal) *User {
	return &User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}
}
