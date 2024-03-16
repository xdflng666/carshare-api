package models

import "time"

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

//"carName": item[0],
//"uuid": item[1],
//"isActive": item[2],
//"location": {"lat": item[3], "lon": item[4]},
//"last_updated": item[5].isoformat()
