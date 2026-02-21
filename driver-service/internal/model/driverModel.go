package model

import "time"

type Status string

const (
	StatusOffline Status = "offline"
	StatusIdle    Status = "idle"
	StatusOnTrip  Status = "on_trip"
)

type Driver struct {
	UserId        string    `json:"userId"`
	VechileNubmer string    `json:"vechileNumber"`
	VechileType   string    `json:"vechileType"`
	Status        Status    `json:"status"`
	IsAvailable   bool      `json:"isAvailable"`
	Rating        float64   `json:"rating"`
	CreatedAt     time.Time `json:"createAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
