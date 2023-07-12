package services

import (
	"waki.mobi/go-yatta-h3i/src/domain/entity"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.ISubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

type ISubscriptionService interface {
	RenewalSubscription() *[]entity.Subscription
	RetrySubscription() *[]entity.Subscription
}

func (s *SubscriptionService) RenewalSubscription() *[]entity.Subscription {
	subs, _ := s.subscriptionRepo.GetRenewal()
	return subs
}

func (s *SubscriptionService) RetrySubscription() *[]entity.Subscription {
	subs, _ := s.subscriptionRepo.GetRetry()
	return subs
}
