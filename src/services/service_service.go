package services

import (
	"waki.mobi/go-yatta-h3i/src/domain/entity"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
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
	// GetService(string) bool
	GetServiceId(int) (*entity.Service, error)
	GetServiceByCode(string) (*entity.Service, error)
}

// func (s *ServiceService) GetService(name string) bool {
// 	count, _ := s.serviceRepo.GetServiceByName(name)
// 	return count > 0
// }

func (s *ServiceService) GetServiceId(id int) (*entity.Service, error) {
	result, err := s.serviceRepo.GetServiceById(id)
	if err != nil {
		return nil, err
	}

	var srv entity.Service

	if result != nil {
		srv = entity.Service{
			ID:              result.ID,
			Code:            result.Code,
			Name:            result.Name,
			Day:             result.Day,
			UrlNotifSub:     result.UrlNotifSub,
			UrlNotifUnsub:   result.UrlNotifUnsub,
			UrlNotifRenewal: result.UrlNotifRenewal,
			UrlPostback:     result.UrlPostback,
		}
	}
	return &srv, nil
}

func (s *ServiceService) GetServiceByCode(code string) (*entity.Service, error) {
	result, err := s.serviceRepo.GetServiceByCode(code)
	if err != nil {
		return nil, err
	}

	var srv entity.Service

	if result != nil {
		srv = entity.Service{
			ID:              result.ID,
			Code:            result.Code,
			Name:            result.Name,
			Day:             result.Day,
			UrlNotifSub:     result.UrlNotifSub,
			UrlNotifUnsub:   result.UrlNotifUnsub,
			UrlNotifRenewal: result.UrlNotifRenewal,
			UrlPostback:     result.UrlPostback,
		}
	}
	return &srv, nil
}
