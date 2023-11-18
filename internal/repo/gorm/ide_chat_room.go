package gorm

import (
	"cody/domain"
	model "cody/internal/repo/model"

	"gorm.io/gorm"
)

func (repo *repository) CreateIdeChatRoom(tx *gorm.DB, ideChatRoomForCreate *domain.IdeChatRoom) (*domain.IdeChatRoom, error) {
	if tx == nil {
		tx = repo.db
	}

	m := model.IdeChatRoomDomainToModel(ideChatRoomForCreate)

	if err := tx.Create(m).Error; err != nil {
		return nil, err
	}

	dic := model.IdeChatRoomModelToDomain(m)

	return dic, nil
}

func (repo *repository) UpdateIdeChatRoom(tx *gorm.DB, ideChatRoomID uint, ideChatRoomForUpdate *domain.IdeChatRoom) (*domain.IdeChatRoom, error) {
	if tx == nil {
		tx = repo.db
	}

	m := model.IdeChatRoomDomainToModel(ideChatRoomForUpdate)

	if err := tx.Where("id = ?", ideChatRoomID).Updates(m).Error; err != nil {
		return nil, err
	}

	dic := model.IdeChatRoomModelToDomain(m)

	return dic, nil
}
