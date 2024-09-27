package gorm

import (
	"gorm.io/gorm"
)

func (repo *repository) NewTransactionBegin() (*gorm.DB, error) {
	tx := repo.db.Begin()

	if err := tx.Error; err != nil {
		return nil, err
	}
	return tx, nil
}
