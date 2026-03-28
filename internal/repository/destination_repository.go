package repository

import "tourism-backend/internal/domain"

type DestinationRepository interface {
	Create(destination *domain.Destination) error
	GetByID(id int) (*domain.Destination, error)
	GetAll() ([]*domain.Destination, error)
	Delete(id int) error
}
