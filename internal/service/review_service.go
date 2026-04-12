// internal/service/review_service.go
package service

import (
	"tourism-backend/internal/domain"
	"tourism-backend/internal/repository"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) Create(review *domain.Review) (*domain.Review, error) {
	if err := s.repo.Create(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) GetByTourID(tourID int) ([]*domain.Review, error) {
	return s.repo.GetByTourID(tourID)
}

func (s *ReviewService) GetByUserID(userID int) ([]*domain.Review, error) {
	return s.repo.GetByUserID(userID)
}

func (s *ReviewService) Delete(id int) error {
	return s.repo.Delete(id)
}
