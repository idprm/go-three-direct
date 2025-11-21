package services

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(
	transactionRepo repository.ITransactionRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
	Save(*entity.Transaction) error
	Update(*entity.Transaction) error
	Delete(*entity.Transaction) error
}

func (s *TransactionService) Save(t *entity.Transaction) error {
	return s.transactionRepo.Save(t)
}

func (s *TransactionService) Update(t *entity.Transaction) error {
	return s.transactionRepo.Update(t)
}

func (s *TransactionService) Delete(t *entity.Transaction) error {
	return s.transactionRepo.Delete(t)
}
