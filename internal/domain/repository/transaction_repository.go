package repository

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

type ITransactionRepository interface {
	Save(*entity.Transaction) error
	Update(*entity.Transaction) error
	Delete(*entity.Transaction) error
}

func (r *TransactionRepository) Save(t *entity.Transaction) error {
	return r.db.Create(t).Error
}

func (r *TransactionRepository) Update(t *entity.Transaction) error {
	return r.db.Save(t).Error
}

func (r *TransactionRepository) Delete(t *entity.Transaction) error {
	return r.db.Delete(t).Error
}
