// internal/service/booking_service.go
package service

import (
	"errors"
	"tourism-backend/internal/domain"
	"tourism-backend/internal/repository"
)

type BookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) Create(booking *domain.Booking) (*domain.Booking, error) {
	booking.Status = domain.BookingStatusPending
	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *BookingService) GetByID(id int) (*domain.Booking, error) {
	booking, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if booking == nil {
		return nil, errors.New("booking not found")
	}
	return booking, nil
}

func (s *BookingService) GetByUserID(userID int) ([]*domain.Booking, error) {
	return s.repo.GetByUserID(userID)
}

func (s *BookingService) GetAll() ([]*domain.Booking, error) {
	return s.repo.GetAll()
}

func (s *BookingService) UpdateStatus(id int, status domain.BookingStatus) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *BookingService) Delete(id int) error {
	return s.repo.Delete(id)
}
