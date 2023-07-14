package handlers

import (
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
)

type DRHandler struct {
	cfg              *config.Secret
	subscriptionRepo repository.ISubscriptionRepository
	transactionRepo  repository.ITransactionRepository
}

func NewDRHandler(cfg *config.Secret,
	subscriptionRepo repository.ISubscriptionRepository,
	transactionRepo repository.ITransactionRepository,
) *DRHandler {
	return &DRHandler{
		cfg:              cfg,
		subscriptionRepo: subscriptionRepo,
		transactionRepo:  transactionRepo,
	}
}

func (h *DRHandler) Sync() {

}
