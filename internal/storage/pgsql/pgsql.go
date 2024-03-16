package pgsql

import (
	"carshare-api/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(dsn string) (*Storage, error) {
	const op = "storage.pgsql.New"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetCarLocations() ([]models.CarLocation, error) {
	const op = "storage.postgres.GetCarLocations"

	query :=
		`SELECT c.name, c.uuid, c.is_active, l.lat, l.lon, l.created_at
		FROM car c
		LEFT JOIN (
			SELECT car_uuid, lat, lon, created_at,
				ROW_NUMBER() OVER (PARTITION BY car_uuid ORDER BY created_at DESC) AS rn
			FROM location
		) l ON c.uuid = l.car_uuid AND l.rn = 1
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %w", op, err)
	}

	var locations []models.CarLocation
	for rows.Next() {
		var location models.CarLocation
		err := rows.Scan(&location.Name, &location.UUID, &location.IsActive, &location.Lat, &location.Lon, &location.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows: %w", op, err)
	}

	return locations, nil
}
