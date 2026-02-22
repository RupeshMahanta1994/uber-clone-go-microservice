package repository

import (
	"context"

	"github.com/rupeshmahanta/driver-service/internal/model"
)

type IDriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) (string, error)
	GetByUserId(ctx context.Context, userId string) (*model.Driver, error)
	Update(ctx context.Context, driver *model.Driver) (string, error)
}
