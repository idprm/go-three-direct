package handlers

import (
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
)

type RenewalHandler struct {
	cfg         *config.Secret
	serviceRepo repository.IServiceRepository
	contentRepo repository.IContentRepository
}

func NewRenewalHandler(
	cfg *config.Secret,
	serviceRepo repository.IServiceRepository,
	contentRepo repository.IContentRepository,
) *RenewalHandler {
	return &RenewalHandler{
		cfg:         cfg,
		serviceRepo: serviceRepo,
		contentRepo: contentRepo,
	}
}

func (h *RenewalHandler) Dailypush() {

}
