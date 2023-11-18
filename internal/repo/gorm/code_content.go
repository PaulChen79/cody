package gorm

import (
	"cody/domain"
	model "cody/internal/repo/model"

	"gorm.io/gorm"
)

func (repo *repository) CreateCodeContent(tx *gorm.DB, codeContentForCreate *domain.CodeContent) (*domain.CodeContent, error) {
	if tx == nil {
		tx = repo.db
	}

	m := model.CodeContentDomainToModel(codeContentForCreate)

	if err := tx.Create(m).Error; err != nil {
		return nil, err
	}

	dc := model.CodeContentModelToDomain(m)

	return dc, nil
}
