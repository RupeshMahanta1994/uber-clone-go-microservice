package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rupeshmahanta/auth-service/internal/model"
	"github.com/rupeshmahanta/auth-service/internal/repository"
	"github.com/rupeshmahanta/auth-service/internal/utils"
)

type IAuthService interface {
	Register(ctx context.Context, email, password, role string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
}
type AuthService struct {
	repo repository.IUserRepository
}

func NewAuthService(repo repository.IUserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, email, password, role string) (string, error) {
	hashed, error := utils.HashPassword(password)
	if error != nil {
		return "", error
	}
	user := &model.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashed,
		Role:     role,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return "", err
	}
	return utils.GenerateToken(user.ID, user.Role)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, error := s.repo.GetByEmailId(ctx, email)
	if error != nil {
		return "Error in Loing Service", nil
	}
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "Passowrd Didn't match", err
	}
	return utils.GenerateToken(user.ID, user.Role)

}
