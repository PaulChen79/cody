package service

import (
	"cody/domain"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (svc *service) CreateIdeChatRoom(creatorID uint, chatRoomID *uint) (*domain.IdeChatRoom, error) {
	creator, err := svc.repo.GetUserInternalByID(creatorID)
	if err != nil {
		zap.S().Warnf("GetUserInternalByID - err: %v", err)
		return nil, err
	}

	tx, err := svc.repo.NewTransactionBegin()
	if err != nil {
		zap.S().Warnf("NewTransactionBegin - err: %v", err)
		return nil, err
	}
	defer tx.Rollback()

	newUUID := uuid.New()

	IdeChatRoom, err := svc.repo.CreateIdeChatRoom(tx, &domain.IdeChatRoom{
		RoomName: "Untitled",
		RoomUUID: newUUID.String(),
	})
	if err != nil {
		zap.S().Warnf("CreateIdeChatRoom - err: %v", err)
		return nil, err
	}

	var chatRoom *domain.ChatRoom

	if chatRoomID == nil {
		chatRoom, err = svc.repo.CreateChatRoom(tx, &domain.ChatRoomFullInfo{
			RoomName:      creator.Username,
			IsGroup:       true,
			Users:         []*domain.UserInternal{creator},
			IdeChatRoomID: IdeChatRoom.ID,
		})
		if err != nil {
			zap.S().Warnf("CreateChatRoom - err: %v", err)
			return nil, err
		}
	} else {
		chatRoom, err = svc.repo.GetChatRoomByID(*chatRoomID)
		if err != nil {
			zap.S().Warnf("GetChatRoomByID - err: %v", err)
			return nil, err
		}
	}

	newIdeChatRoom, err := svc.repo.UpdateIdeChatRoom(tx, IdeChatRoom.ID, &domain.IdeChatRoom{
		ChatRoom: chatRoom,
	})
	if err != nil {
		zap.S().Warnf("UpdateIdeChatRoom - err: %v", err)
		return nil, err
	}

	tx.Commit()

	return newIdeChatRoom, nil
}
