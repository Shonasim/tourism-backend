package queries

const (
	CreatePayment = `
INSERT INTO payments (booking_id, amount, status)
VALUES ($1, $2, $3)
RETURNING id, created_at`

	GetPaymentByID = `
SELECT id, booking_id, amount, status, created_at
FROM payments WHERE id = $1`

	GetPaymentByBookingID = `
SELECT id, booking_id, amount, status, created_at
FROM payments WHERE booking_id = $1`

	UpdatePaymentStatus = `
UPDATE payments SET status = $1 WHERE id = $2`
)
