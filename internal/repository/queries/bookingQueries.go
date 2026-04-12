package queries

const (
	CreateBooking = `
INSERT INTO bookings (user_id, tour_id, status)
VALUES ($1, $2, $3)
RETURNING id, created_at`

	GetBookingByID = `
SELECT id, user_id, tour_id, status, created_at
FROM bookings WHERE id = $1`

	GetBookingsByUserID = `
SELECT id, user_id, tour_id, status, created_at
FROM bookings WHERE user_id = $1`

	GetAllBookings = `
SELECT id, user_id, tour_id, status, created_at
FROM bookings`

	UpdateBookingStatus = `
UPDATE bookings SET status = $1 WHERE id = $2`

	DeleteBooking = `
DELETE FROM bookings WHERE id = $1`
)
