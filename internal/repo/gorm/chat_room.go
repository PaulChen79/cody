package gorm

import (
	"cody/domain"
	model "cody/internal/repo/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (repo *repository) CreateChatRoom(tx *gorm.DB, chatRoomForUpdate *domain.ChatRoomFullInfo) (*domain.ChatRoomFullInfo, error) {
	if tx == nil {
		tx = repo.db
	}

	newChatRoom := model.ChatRoomFullInfoDomainToModel(chatRoomForUpdate)

	if err := tx.Create(newChatRoom).Error; err != nil {
		zap.S().Warnf("CreateChatRoom - err: %v", err)
		return nil, err
	}

	dc := model.ChatRoomModelToFullInfoDomain(newChatRoom)

	return dc, nil
}

func (repo *repository) UpdateChatRoom(tx *gorm.DB, chatRoomID uint, chatRoomForUpdate *domain.ChatRoomFullInfo) (*domain.ChatRoom, error) {
	if tx == nil {
		tx = repo.db
	}

	m := model.ChatRoomFullInfoDomainToModel(chatRoomForUpdate)

	if err := tx.Where("id = ?", chatRoomID).Updates(&m).Error; err != nil {
		zap.S().Warnf("CreateChatRoom - err: %v", err)
		return nil, err
	}

	dc := model.ChatRoomModelToDomain(m)

	return dc, nil
}

func (repo *repository) GetChatRoomByID(chatRoomID uint) (*domain.ChatRoomFullInfo, error) {
	var chatRoom model.ChatRoom

	if err := repo.db.Where("id = ?", chatRoomID).First(&chatRoom).Error; err != nil {
		zap.S().Warnf("GetChatRoomByID - err: %v", err)
		return nil, err
	}

	dc := model.ChatRoomModelToFullInfoDomain(&chatRoom)

	return dc, nil
}
