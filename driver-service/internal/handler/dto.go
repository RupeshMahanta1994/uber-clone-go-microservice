package handler

type OnboardRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType   string `json:"vehicle_type"`
}
