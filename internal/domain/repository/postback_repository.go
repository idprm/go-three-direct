package repository

import "gorm.io/gorm"

type PostbackRepository struct {
	db *gorm.DB
}

func NewPostbackRepository(db *gorm.DB) *PostbackRepository {
	return &PostbackRepository{
		db: db,
	}
}

type IPostbackRepository interface {
}
