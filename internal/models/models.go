package models

import "time"

type Car struct {
	Name     string `json:"carName"`
	UUID     string `json:"uuid"`
	IsActive bool   `json:"isActive"`
}

type CarLocation struct {
	Name        string    `json:"carName"`
	UUID        string    `json:"uuid"`
	IsActive    bool      `json:"isActive"`
	LastUpdated time.Time `json:"last_updated"`
	Location    struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
