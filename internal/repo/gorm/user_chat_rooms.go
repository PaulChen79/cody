package gorm

import (
	model "cody/internal/repo/model"
	"time"

	"gorm.io/gorm"
)

func (repo *repository) RefreshLastReadAt(tx *gorm.DB, userID uint, chatRoomID uint) error {
	if tx == nil {
		tx = repo.db
	}

	now := time.Now()

	if err := tx.Model(&model.UserChatRoom{}).Where("user_id = ? AND chat_room_id = ?", userID, chatRoomID).Update("last_read_at", now).Error; err != nil {
		return err
	}

	return nil
}
