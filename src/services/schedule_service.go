package services

import (
	"waki.mobi/go-yatta-h3i/src/domain/entity"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
)

type ScheduleService struct {
	scheduleRepo repository.IScheduleRepository
}

func NewScheduleService(scheduleRepo repository.IScheduleRepository) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

type IScheduleService interface {
	GetLocked(string, string) bool
	GetUnlocked(string, string) bool
	UpdateSchedule(bool, string)
}

func (s *ScheduleService) GetLocked(name string, hour string) bool {
	count, _ := s.scheduleRepo.CountLocked(name, hour)
	return count > 0
}

func (s *ScheduleService) GetUnlocked(name string, hour string) bool {
	count, _ := s.scheduleRepo.CountUnlocked(name, hour)
	return count > 0
}

func (s *ScheduleService) UpdateSchedule(unlocked bool, name string) {
	s.scheduleRepo.ScheduleUpdate(
		&entity.Schedule{
			Status: unlocked,
			Name:   name,
		},
	)
}
