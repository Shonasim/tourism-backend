package domain

import "time"

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TourID    int       `json:"tour_id"` // ← исправлено
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
