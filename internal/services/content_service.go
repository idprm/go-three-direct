package services

import (
	"log"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
)

type ContentService struct {
	contentRepo repository.IContentRepository
}

func NewContentService(contentRepo repository.IContentRepository) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
	}
}

type IContentService interface {
	IsContent(int, string) bool
	Get(int, string) (*entity.Content, error)
}

func (s *ContentService) IsContent(serviceId int, name string) bool {
	count, err := s.contentRepo.Count(serviceId, name)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (s *ContentService) Get(serviceId int, name string) (*entity.Content, error) {
	return s.contentRepo.Get(serviceId, name)
}
