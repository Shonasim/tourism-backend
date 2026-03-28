package repository

import "tourism-backend/internal/domain"

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	GetByID(id int) (*domain.Payment, error)
	GetByBookingID(bookingID int) (*domain.Payment, error)
	UpdateStatus(id int, status domain.PaymentStatus) error
}
