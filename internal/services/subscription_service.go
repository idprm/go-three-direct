package services

import (
	"log"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

func NewSubscriptionService(
	subscriptionRepo repository.ISubscriptionRepository,
) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

type ISubscriptionService interface {
	IsActive(int, string) bool
	IsSub(int, string) bool
	IsRenewal(int, string) bool
	IsRetry(int, string) bool
	Get(int, string) (*entity.Subscription, error)
	Save(*entity.Subscription) error
	Update(*entity.Subscription) error
}

func (s *SubscriptionService) IsActive(serviceId int, msisdn string) bool {
	count, err := s.subscriptionRepo.CountActive(serviceId, msisdn)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *SubscriptionService) IsSub(serviceId int, msisdn string) bool {
	count, err := s.subscriptionRepo.Count(serviceId, msisdn)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *SubscriptionService) IsRenewal(serviceId int, msisdn string) bool {
	count, err := s.subscriptionRepo.CountRenewal(serviceId, msisdn)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *SubscriptionService) IsRetry(serviceId int, msisdn string) bool {
	count, err := s.subscriptionRepo.CountRetry(serviceId, msisdn)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *SubscriptionService) Get(serviceId int, msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.Get(serviceId, msisdn)
}

func (s *SubscriptionService) Save(e *entity.Subscription) error {
	return s.subscriptionRepo.Save(e)
}

func (s *SubscriptionService) Update(e *entity.Subscription) error {
	return s.subscriptionRepo.Update(e)
}
