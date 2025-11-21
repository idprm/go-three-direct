package services

import (
	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
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
	Update(bool, string)
}

func (s *ScheduleService) GetLocked(name string, hour string) bool {
	count, _ := s.scheduleRepo.CountLocked(name, hour)
	return count > 0
}

func (s *ScheduleService) GetUnlocked(name string, hour string) bool {
	count, _ := s.scheduleRepo.CountUnlocked(name, hour)
	return count > 0
}

func (s *ScheduleService) Update(unlocked bool, name string) {
	s.scheduleRepo.ScheduleUpdate(
		&entity.Schedule{
			Status: unlocked,
			Name:   name,
		},
	)
}
