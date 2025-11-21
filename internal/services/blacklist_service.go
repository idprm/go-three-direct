package services

import "github.com/idprm/go-three-direct/internal/domain/repository"

type BlacklistService struct {
	blacklistRepo repository.IBlacklistRepository
}

func NewBlacklistService(blacklistRepo repository.IBlacklistRepository) *BlacklistService {
	return &BlacklistService{
		blacklistRepo: blacklistRepo,
	}
}

type IBlacklistService interface {
	CountByMsisdn(msisdn string) bool
}

func (s *BlacklistService) CountByMsisdn(msisdn string) bool {
	count, _ := s.blacklistRepo.CountByMsisdn(msisdn)
	return count > 0
}
