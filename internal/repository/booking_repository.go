package repository

import "tourism-backend/internal/domain"

type BookingRepository interface {
	Create(booking *domain.Booking) error
	GetByID(id int) (*domain.Booking, error)
	GetByUserID(userID int) ([]*domain.Booking, error)
	GetAll() ([]*domain.Booking, error)
	UpdateStatus(id int, status domain.BookingStatus) error
	Delete(id int) error
}
