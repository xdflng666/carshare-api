package pgsql

import (
	"carshare-api/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

func (s *Storage) PostCarLocation(lat, lon float64, carUUID string) error {
	const op = "storage.pgsql.PostCarLocation"

	rows, err := s.db.Query(`INSERT INTO location (lat, lon, car_uuid) VALUES ($1, $2, $3)`, lat, lon, carUUID)

	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	fmt.Println(rows)

	return nil
}

func (s *Storage) GetCars() ([]models.Car, error) {
	const op = "storage.postgres.GetCars"

	query := `SELECT name, uuid, is_active FROM car`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: query: %w", op, err)
	}

	var cars []models.Car
	for rows.Next() {
		car := models.Car{}
		err = rows.Scan(&car.Name, &car.UUID, &car.IsActive)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows: %w", op, err)
	}

	return cars, nil
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
		err := rows.Scan(
			&location.Name,
			&location.UUID,
			&location.IsActive,
			&location.Location.Lat,
			&location.Location.Lon,
			&location.LastUpdated,
		)
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

func (s *Storage) GetAuthCredentials() map[string]string {
	const op = "storage.postgres.GetAuthCredentials"

	query := "SELECT username, password FROM users"

	rows, err := s.db.Query(query)

	if err != nil {
		log.Printf("%s:%s\n", op, err)
		log.Fatal("Couldn't load auth credentials")
	}

	credMap := make(map[string]string)

	for rows.Next() {
		var username, password string
		err = rows.Scan(&username, &password)
		if err != nil {
			log.Fatal("Error while reading auth credentials")
		}
		credMap[username] = password
	}
	if rows.Err() != nil {
		log.Fatal("Error while reading auth credentials")
	}

	return credMap
}

func (s *Storage) GetStory(carUUID string) ([]models.Point, error) {
	const op = "storage.postgres.GetStoryOf"

	query := "SELECT lat, lon FROM location WHERE car_uuid = $1 LIMIT 5"

	rows, err := s.db.Query(query, carUUID)

	if err != nil {
		return nil, fmt.Errorf("%s: query: %w", op, err)
	}

	var coordinates []models.Point
	for rows.Next() {
		var p models.Point
		err := rows.Scan(&p.Lat, &p.Lon)
		if err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		coordinates = append(coordinates, p)
	}

	return coordinates, nil
}
