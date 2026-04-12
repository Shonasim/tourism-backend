package queries

// TOURS
const (
	CreateTour = `
INSERT INTO tours (destination_id,name,description,price,start_date,end_date,capacity)
VALUES ($1,$2,$3,$4,$5,$6,$7)
returning id,created_at`
	GetTourByID = `
SELECT * from tours where id = $1`
	GetAllTours = `
SELECT * from tours`
	GetByDestinationID = `ю
SELECT * from tours where destination_id = $1`
	UpdateTour = `
UPDATE tours SET destination_id = $1, name = $2, description = $3, price = $4,start_date = $5,end_date = $6,capacity = $7
WHERE id = $8`
	DeleteTour = `
	DELETE FROM tours WHERE id = $1`
	GetToursByDestinationID = `
SELECT id, destination_id, name, description,
       price, start_date, end_date, capacity, created_at
FROM tours WHERE destination_id = $1`
)
