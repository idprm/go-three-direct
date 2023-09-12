package handler

import (
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
)

type MOHandler struct {
	cfg         *config.Secret
	serviceRepo repository.IServiceRepository
	contentRepo repository.IContentRepository
}

func NewMOHandler(
	cfg *config.Secret,
	serviceRepo repository.IServiceRepository,
	contentRepo repository.IContentRepository,
) *MOHandler {
	return &MOHandler{
		cfg:         cfg,
		serviceRepo: serviceRepo,
		contentRepo: contentRepo,
	}
}

func (h *MOHandler) Firstpush() {

}
