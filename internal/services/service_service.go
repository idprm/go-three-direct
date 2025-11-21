package services

import (
	"log"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
)

type ServiceService struct {
	serviceRepo repository.IServiceRepository
}

func NewServiceService(
	serviceRepo repository.IServiceRepository,
) *ServiceService {
	return &ServiceService{
		serviceRepo: serviceRepo,
	}
}

type IServiceService interface {
	CountByCode(string) bool
	GetById(int) (*entity.Service, error)
	GetByCode(string) (*entity.Service, error)
}

func (s *ServiceService) CountByCode(code string) bool {
	count, err := s.serviceRepo.CountByCode(code)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *ServiceService) GetById(v int) (*entity.Service, error) {
	return s.serviceRepo.GetById(v)
}

func (s *ServiceService) GetByCode(v string) (*entity.Service, error) {
	return s.serviceRepo.GetByCode(v)
}
