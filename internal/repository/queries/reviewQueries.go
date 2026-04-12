package queries

const (
	CreateReview = `
INSERT INTO reviews (user_id, tour_id, rating, comment)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at`

	GetReviewsByTourID = `
SELECT id, user_id, tour_id, rating, comment, created_at
FROM reviews WHERE tour_id = $1`

	GetReviewsByUserID = `
SELECT id, user_id, tour_id, rating, comment, created_at
FROM reviews WHERE user_id = $1`

	DeleteReview = `
DELETE FROM reviews WHERE id = $1`
)
