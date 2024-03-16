package models

import "time"

type CarLocation struct {
	Name      string
	UUID      string
	IsActive  bool
	Lat       float64
	Lon       float64
	CreatedAt time.Time
}
