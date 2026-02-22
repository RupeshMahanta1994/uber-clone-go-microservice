package service

import (
	"context"

	"github.com/rupeshmahanta/driver-service/internal/model"
)

type IDriverService interface {
	Onboard(ctx context.Context, driver *model.Driver) (string, error)
	ToggleAvailability(ctx context.Context, userID string, available bool) (string, error)
	UpdateStatus(ctx context.Context, userID string, status model.Status) (string, error)
	GetDriver(ctx context.Context, userID string) (*model.Driver, error)
}
