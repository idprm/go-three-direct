package services

import "waki.mobi/go-yatta-h3i/src/domain/repository"

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(transactionRepo repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
}
