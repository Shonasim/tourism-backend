// internal/service/destination_service.go
package service

import (
	"errors"
	"tourism-backend/internal/domain"
	"tourism-backend/internal/repository"
)

type DestinationService struct {
	repo repository.DestinationRepository
}

func NewDestinationService(repo repository.DestinationRepository) *DestinationService {
	return &DestinationService{repo: repo}
}

func (s *DestinationService) Create(destination *domain.Destination) (*domain.Destination, error) {
	if err := s.repo.Create(destination); err != nil {
		return nil, err
	}
	return destination, nil
}

func (s *DestinationService) GetByID(id int) (*domain.Destination, error) {
	destination, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if destination == nil {
		return nil, errors.New("destination not found")
	}
	return destination, nil
}

func (s *DestinationService) GetAll() ([]*domain.Destination, error) {
	return s.repo.GetAll()
}

func (s *DestinationService) Delete(id int) error {
	return s.repo.Delete(id)
}
