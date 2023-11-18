package gorm

import (
	"cody/domain"
	model "cody/internal/repo/model"

	"go.uber.org/zap"
)

func (repo *repository) GetUserByID(userID uint) (*domain.User, error) {
	var user model.User

	if err := repo.db.Where("id = ?", userID).First(&user).Error; err != nil {
		zap.S().Warnf("GetUserByID - err: %v", err)
		return nil, err
	}

	dc := model.UserModelToDomain(&user)

	return dc, nil
}

func (repo *repository) GetUserInternalByID(userID uint) (*domain.UserInternal, error) {
	var user model.User

	if err := repo.db.Where("id = ?", userID).First(&user).Error; err != nil {
		zap.S().Warnf("GetUserByID - err: %v", err)
		return nil, err
	}

	dc := model.UserModelToInternalDomain(user)

	return dc, nil
}
