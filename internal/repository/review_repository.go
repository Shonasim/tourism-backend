package repository

import "tourism-backend/internal/domain"

type ReviewRepository interface {
	Create(review *domain.Review) error
	GetByTourID(tourID int) ([]*domain.Review, error)
	GetByUserID(userID int) ([]*domain.Review, error)
	Delete(id int) error
}
