package repository

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

type IServiceRepository interface {
	CountByCode(string) (int64, error)
	GetById(int) (*entity.Service, error)
	GetByCode(string) (*entity.Service, error)
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) CountByCode(code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Service{}).Where("code = ?", code).Count(&count).Error
	return count, err
}

func (r *ServiceRepository) GetById(v int) (*entity.Service, error) {
	var e entity.Service
	err := r.db.Where("id = ?", v).First(&e).Error
	return &e, err
}

func (r *ServiceRepository) GetByCode(v string) (*entity.Service, error) {
	var e entity.Service
	err := r.db.Where("code = ?", v).First(&e).Error
	return &e, err
}
