package domain

import "gorm.io/gorm"

type Repository interface {
	NewTransactionBegin() (*gorm.DB, error)

	// * User
	GetUserByID(userID uint) (*User, error)
	GetUserInternalByID(userID uint) (*UserInternal, error)

	// * Chat Room
	CreateChatRoom(tx *gorm.DB, chatRoomForUpdate *ChatRoomFullInfo) (*ChatRoomFullInfo, error)
	UpdateChatRoom(tx *gorm.DB, chatRoomID uint, chatRoomForUpdate *ChatRoomFullInfo) (*ChatRoom, error)
	GetChatRoomByID(chatRoomID uint) (*ChatRoomFullInfo, error)

	// * IDE Chat Room
	CreateIdeChatRoom(tx *gorm.DB, ideChatRoomForCreate *IdeChatRoom) (*IdeChatRoom, error)
	UpdateIdeChatRoom(tx *gorm.DB, ideChatRoomID uint, ideChatRoomForUpdate *IdeChatRoom) (*IdeChatRoom, error)

	// * Code Content
	CreateCodeContent(tx *gorm.DB, codeContentForCreate *CodeContent) (*CodeContent, error)

	// * User Chat Room
	RefreshLastReadAt(tx *gorm.DB, userID uint, chatRoomID uint) error
}
