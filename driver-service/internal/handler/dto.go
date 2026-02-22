package handler

import "github.com/rupeshmahanta/driver-service/internal/model"

type OnboardRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType   string `json:"vehicle_type"`
}

type ToggleAvailabilityRequest struct {
	Availability bool `json:"availability"`
}

type StatusRequest struct {
	Status model.Status `json:"status"`
}
