package services

import (
	"log"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
)

type PopulateService struct {
	populateRepo repository.IPopulateRepository
}

func NewPopulateService(
	populateRepo repository.IPopulateRepository,
) *PopulateService {
	return &PopulateService{
		populateRepo: populateRepo,
	}
}

type IPopulateService interface {
	Renewal() *[]entity.Subscription
	Retry() *[]entity.Subscription
}

func (s *PopulateService) Renewal() *[]entity.Subscription {
	subs, err := s.populateRepo.SelectRenewal()
	if err != nil {
		log.Println(err)
		return nil
	}
	return subs
}

func (s *PopulateService) Retry() *[]entity.Subscription {
	subs, err := s.populateRepo.SelectRetry()
	if err != nil {
		log.Println(err)
		return nil
	}
	return subs
}
