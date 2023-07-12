package services

import "waki.mobi/go-yatta-h3i/src/domain/repository"

type AdnetService struct {
	adnetRepo repository.IAdnetRepository
}

func NewAdnetService(adnetRepo repository.IAdnetRepository) *AdnetService {
	return &AdnetService{
		adnetRepo: adnetRepo,
	}
}

type IAdnetService interface {
}
