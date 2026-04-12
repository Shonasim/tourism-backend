package queries

const (
	CreateDestination = `
INSERT INTO destinations (name, description)
VALUES ($1, $2)
RETURNING id, created_at`
	GetDestinationByID = `
SELECT id, name, description, created_at
FROM destinations WHERE id = $1`
	GetAllDestinations = `
SELECT id, name, description, created_at
FROM destinations`
	DeleteDestination = `
DELETE FROM destinations WHERE id = $1`
)
