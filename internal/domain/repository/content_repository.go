package repository

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

type IContentRepository interface {
	Count(int, string) (int64, error)
	Get(int, string) (*entity.Content, error)
}

func (r *ContentRepository) Count(serviceId int, name string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Content{}).Where("service_id = ? AND name = ?", serviceId, name).Count(&count).Error
	return count, err
}

func (r *ContentRepository) Get(serviceId int, name string) (*entity.Content, error) {
	var e entity.Content
	err := r.db.Where("service_id = ? AND name = ?", serviceId, name).First(&e).Error
	return &e, err
}
