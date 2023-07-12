package services

import (
	"waki.mobi/go-yatta-h3i/src/domain/entity"
	"waki.mobi/go-yatta-h3i/src/domain/repository"
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
	GetContent(int, string) (*entity.Content, error)
}

func (s *ContentService) GetContent(serviceId int, name string) (*entity.Content, error) {
	result, err := s.contentRepo.Get(serviceId, name)
	if err != nil {
		return nil, err
	}

	var content entity.Content

	if result != nil {
		content = entity.Content{
			Value: result.Value,
		}
	}
	return &content, nil
}
