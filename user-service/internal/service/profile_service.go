package service

import (
	"context"
	"errors"
	"github/rupeshmahanta/user-service/internal/models"
	"github/rupeshmahanta/user-service/internal/repository"
	"github/rupeshmahanta/user-service/internal/utils"
)

type IProfileService interface {
	CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetProfileById(ctx context.Context, userId string) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) error
}

type ProfileService struct {
	repo repository.IProfileRepository
}

func NewProfileService(repo repository.IProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	profile.UserID = utils.GenerateId()
	if profile.Rating == 0 {
		profile.Rating = 5.0
	}
	userData, err := s.repo.CreateProfile(ctx, profile)
	if err != nil {
		return nil, errors.New("Error in getting User Data")
	}
	return userData, err

}

func (s *ProfileService) GetProfileById(ctx context.Context, userId string) (*models.Profile, error) {
	profile, err := s.repo.GetProfileByUserId(ctx, userId)
	if err != nil {
		return profile, errors.New("Error in getting user data" + err.Error())
	}
	return profile, nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	err := s.repo.UpdateProfile(ctx, profile)
	if err != nil {
		return errors.New("Error in updating Profiel" + err.Error())
	}
	return nil
}
