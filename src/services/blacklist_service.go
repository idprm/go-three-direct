package services

import "waki.mobi/go-yatta-h3i/src/domain/repository"

type BlacklistService struct {
	blacklistRepo repository.IBlacklistRepository
}

func NewBlacklistService(blacklistRepo repository.IBlacklistRepository) *BlacklistService {
	return &BlacklistService{
		blacklistRepo: blacklistRepo,
	}
}

type IBlacklistService interface {
	GetBlacklist(msisdn string) bool
}

func (s *BlacklistService) GetBlacklist(msisdn string) bool {
	count, _ := s.blacklistRepo.CountByMsisdn(msisdn)
	return count > 0
}
