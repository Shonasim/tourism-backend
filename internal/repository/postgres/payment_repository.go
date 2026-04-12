package postgres

import (
	"database/sql"
	"tourism-backend/internal/domain"
	"tourism-backend/internal/repository/queries"
)

type PaymentRepositoryPostgres struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepositoryPostgres {
	return &PaymentRepositoryPostgres{db: db}
}

func (r *PaymentRepositoryPostgres) Create(payment *domain.Payment) error {
	return r.db.QueryRow(queries.CreatePayment,
		payment.BookingID,
		payment.Amount,
		payment.Status,
	).Scan(&payment.ID, &payment.CreatedAt)
}

func (r *PaymentRepositoryPostgres) GetByID(id int) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.QueryRow(queries.GetPaymentByID, id).Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &payment, err
}

func (r *PaymentRepositoryPostgres) GetByBookingID(bookingID int) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.QueryRow(queries.GetPaymentByBookingID, bookingID).Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &payment, err
}

func (r *PaymentRepositoryPostgres) UpdateStatus(id int, status domain.PaymentStatus) error {
	_, err := r.db.Exec(queries.UpdatePaymentStatus, status, id)
	return err
}
