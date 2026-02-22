package model

import "time"

type Status string

const (
	StatusOffline Status = "offline"
	StatusIdle    Status = "idle"
	StatusOnTrip  Status = "on_trip"
)

type Driver struct {
	UserID        string    `json:"user_id"`
	VehicleNumber string    `json:"vehicle_number"`
	VehicleType   string    `json:"vehicle_type"`
	Status        Status    `json:"status"`
	IsAvailable   bool      `json:"is_available"`
	Rating        float64   `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
