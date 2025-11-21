package services

import "github.com/idprm/go-three-direct/internal/domain/repository"

type PostbackService struct {
	postbackRepo repository.IPostbackRepository
}

func NewPostbackService(postbackRepo repository.IPostbackRepository) *PostbackService {
	return &PostbackService{
		postbackRepo: postbackRepo,
	}
}

type IPostbackService interface {
}

func (s *PostbackService) Handle() error {
	return nil
}
