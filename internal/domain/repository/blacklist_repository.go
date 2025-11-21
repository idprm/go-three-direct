package repository

import (
	"gorm.io/gorm"
)

type BlacklistRepository struct {
	db *gorm.DB
}

func NewBlacklistRepository(db *gorm.DB) *BlacklistRepository {
	return &BlacklistRepository{
		db: db,
	}
}

type IBlacklistRepository interface {
	CountByMsisdn(string) (int64, error)
}

func (r *BlacklistRepository) CountByMsisdn(msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&BlacklistRepository{}).Where("msisdn = ?", msisdn).Count(&count).Error
	return count, err
}
