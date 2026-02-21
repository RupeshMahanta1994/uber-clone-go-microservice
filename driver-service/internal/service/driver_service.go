package service

import (
	"context"
	"errors"
	"log"

	"github.com/rupeshmahanta/driver-service/internal/model"
	"github.com/rupeshmahanta/driver-service/internal/repository"
	"github.com/rupeshmahanta/driver-service/internal/utils"
)

type DriverService struct {
	repo repository.IDriverRepository
}

func NewDriverProfile(repo repository.IDriverRepository) *DriverService {
	return &DriverService{repo: repo}
}
func (s *DriverService) Onboard(ctx context.Context, driver *model.Driver) (string, error) {
	driver.UserId = utils.GenerateId()
	driver.Status = "offline"
	driver.IsAvailable = false
	driver.Rating = 5.0
	return s.repo.Create(ctx, driver)
}
func (s *DriverService) ToggleAvailability(ctx context.Context, userID string, available bool) (string, error) {

	driver, err := s.repo.GetByUserId(ctx, userID)
	if err != nil {
		log.Println("Error in getting Driver Info from while Changin availability", err)
		return "Error in getting driver Info", nil
	}
	if driver.Status == model.StatusOnTrip && !available {
		log.Println("Driver is on Trip")
		return "Driver is on Trip", nil
	}
	driver.IsAvailable = available
	if available {
		log.Println("Driver is ready for trip")
		driver.Status = model.StatusIdle
	} else {
		log.Println("Driver went offline")
		driver.Status = model.StatusOffline
	}
	return s.repo.Update(ctx, driver)

}
func (s *DriverService) UpdateStatus(ctx context.Context, userID string, status model.Status) (string, error) {
	driver, err := s.repo.GetByUserId(ctx, userID)
	if err != nil {
		log.Println("Error in getting Driver Info from while Updating status", err)
		return "Error in getting driver Info", errors.New("Error in driver info")
	}
	if driver.Status == model.StatusOffline && status == model.StatusOnTrip {
		log.Println("Offline driver cannot start trip")
		return "Offline driver canot start trip", errors.New("Driver is offline")
	}
	driver.Status = status
	return s.repo.Update(ctx, driver)
}
func (s *DriverService) GetDriver(ctx context.Context, userID string) (*model.Driver, error) {
	return s.repo.GetByUserId(ctx, userID)

}
